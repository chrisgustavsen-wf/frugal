/*
 * Copyright 2017 Workiva
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package compiler_test

import (
	"bufio"
	"flag"
	"io"
	"os"
	"testing"

	"github.com/Workiva/frugal/compiler"
)

const (
	outputDir = "testdata/out"
	delim     = "."
)

var (
	frugalGenFile       = idl("variety.frugal")
	includeVendor       = idl("include_vendor.frugal")
	includeVendorNoPath = idl("include_vendor_no_path.frugal")
	vendorNamespace     = idl("vendor_namespace.frugal")

	// Flag that can be passed to `go test` to update the gold files
	copyFilesPtr = flag.Bool("copy-files", false, "")
)

func idl(name string) string { return "testdata/idl/" + name }
func exp(name string) string { return "testdata/expected/" + name }
func gen(name string) string { return outputDir + "/" + name }

type Comparison struct {
	ExpectedPath  string
	GeneratedPath string
}

type ComparisonList []Comparison

func (pairs ComparisonList) Run(t *testing.T, options compiler.Options) {
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}
	for i, pair := range pairs {
		if pair.GeneratedPath[0] != '/' { // don't modify absolute paths
			pairs[i].GeneratedPath = gen(pair.GeneratedPath)
		}
		pairs[i].ExpectedPath = exp(pair.ExpectedPath)
	}
	copyAllFiles(t, pairs)
	compareAllFiles(t, pairs)
}

func compareFiles(t *testing.T, expectedPath, generatedPath string) {
	expected, err := os.Open(expectedPath)
	if err != nil {
		t.Fatalf("Failed to open file %s", expectedPath)
	}
	defer expected.Close()
	generated, err := os.Open(generatedPath)
	if err != nil {
		t.Fatalf("Failed to open file %s", generatedPath)
	}
	defer generated.Close()

	expectedScanner := bufio.NewScanner(expected)
	generatedScanner := bufio.NewScanner(generated)
	line := 1
	for expectedScanner.Scan() {
		generatedScanner.Scan()
		expectedLine := expectedScanner.Text()
		generatedLine := generatedScanner.Text()
		if expectedLine != generatedLine {
			t.Fatalf("\nExpected line \n<%s> (%s)\n generated line\n <%s> (%s) at line %d",
				expectedLine, expectedPath, generatedLine, generatedPath, line)
		}
		line++
	}

	if generatedScanner.Scan() {
		t.Fatalf("Generated has more lines than expected: exp %s gen %s", expectedPath, generatedPath)
	}
}

func compareAllFiles(t *testing.T, pairs []Comparison) {
	for _, pair := range pairs {
		compareFiles(t, pair.ExpectedPath, pair.GeneratedPath)
	}
}

func copyAllFiles(t *testing.T, pairs []Comparison) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if !*copyFilesPtr {
		return
	}

	for _, pair := range pairs {
		if err := copyFilePair(pair); err != nil {
			t.Fatal(err)
		}
	}
}

func copyFilePair(pair Comparison) error {
	// TODO automatically create a missing expected file?
	generatedFile, err := os.Open(pair.GeneratedPath)
	if err != nil {
		return err
	}
	defer generatedFile.Close()

	expectedFile, err := os.Create(pair.ExpectedPath)
	if err != nil {
		return err
	}
	defer expectedFile.Close()

	_, err = io.Copy(expectedFile, generatedFile)
	return err
}

func assertFilesNotExist(t *testing.T, filePaths []string) {
	for _, fileThatShouldNotExist := range filePaths {
		if _, err := os.Stat(fileThatShouldNotExist); !os.IsNotExist(err) {
			if err != nil {
				t.Errorf("Unexpected error checking for existence on %q: %s", fileThatShouldNotExist, err)
			} else {
				t.Errorf("Expected %q not to exist, but it did", fileThatShouldNotExist)
			}
		}
	}
}
