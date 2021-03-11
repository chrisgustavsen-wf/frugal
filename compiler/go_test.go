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
	"testing"

	"github.com/Workiva/frugal/compiler"
)

func TestValidGoWithAsync(t *testing.T) {
	options := compiler.Options{
		File:  frugalGenFile,
		Gen:   "go:package_prefix=github.com/Workiva/frugal/test/_out/async/,async",
		Out:   outputDir + "/async",
		Delim: delim,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"go/variety_async/f_foo_service.txt", "async/variety/f_foo_service.go"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

func TestValidGoFrugalCompiler(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "go:package_prefix=github.com/Workiva/frugal/test/_out/",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"go/actual_base/f_types.txt", "actual_base/golang/f_types.go"},
		{"go/actual_base/f_basefoo_service.txt", "actual_base/golang/f_basefoo_service.go"},

		{"go/variety/f_types.txt", "variety/f_types.go"},
		{"go/variety/f_foo_service.txt", "variety/f_foo_service.go"},
		{"go/variety/f_events_scope.txt", "variety/f_events_scope.go"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

// Ensures correct import references are used when -use-vendor is set and the
// IDL has a vendored include.
func TestValidGoVendor(t *testing.T) {
	options := compiler.Options{
		File:  includeVendor,
		Gen:   "go:package_prefix=github.com/Workiva/frugal/test/_out/,use_vendor",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"go/vendor/f_myscope_scope.txt", "include_vendor/f_myscope_scope.go"},
		{"go/vendor/f_myservice_service.txt", "include_vendor/f_myservice_service.go"},
		{"go/vendor/f_types.txt", "include_vendor/f_types.go"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

// Ensures an error is returned when -use-vendor is set and the vendored
// include does not specify a path.
func TestValidGoVendorPathNotSpecified(t *testing.T) {
	options := compiler.Options{
		File:  includeVendorNoPath,
		Gen:   "go:package_prefix=github.com/Workiva/frugal/test/_out/,use_vendor",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err == nil {
		t.Fatal("Expected error")
	}
}

// Ensures the target IDL is generated when -use-vendor is set and it has a
// vendored namespace.
func TestValidGoVendorNamespaceTargetGenerate(t *testing.T) {
	options := compiler.Options{
		File:  vendorNamespace,
		Gen:   "go:package_prefix=github.com/Workiva/frugal/test/_out/,use_vendor",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"go/vendor_namespace/f_types.txt", "vendor_namespace/f_types.go"},
		{"go/vendor_namespace/f_vendoredbase_service.txt", "vendor_namespace/f_vendoredbase_service.go"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

// Ensures includes are generated in the same order
func TestIncludeOrdering(t *testing.T) {
	options := compiler.Options{
		File:    idl("ordering/main.frugal"),
		Gen:     "go:package_prefix=github.com/Workiva/frugal/test/_out/ordering",
		Out:     gen("ordering"),
		Delim:   delim,
		Recurse: true,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"go/ordering/one/f_types.go", "ordering/one/f_types.go"},
		{"go/ordering/two/f_types.go", "ordering/two/f_types.go"},
		{"go/ordering/three/f_types.go", "ordering/three/f_types.go"},
		{"go/ordering/four/f_types.go", "ordering/four/f_types.go"},
		{"go/ordering/five/f_types.go", "ordering/five/f_types.go"},
	}

	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

// Ensures slim generated files are correct.
func TestSlim(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "go:package_prefix=github.com/Workiva/frugal/test/_out/,slim",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"go/slim/f_types.go", "variety/f_types.go"},
		{"go/slim/f_foo_service.go", "variety/f_foo_service.go"},
		{"go/slim/f_events_scope.go", "variety/f_events_scope.go"},
	}

	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

// Ensures deprecated logging can be suppressed
func TestSuppressedDeprecatedLogging(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "go:package_prefix=github.com/Workiva/frugal/test/_out/,suppress_deprecated_logging",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"go/deprecated_logging/f_foo_service.go", "variety/f_foo_service.go"},
	}

	copyAllFiles(t, files)
	compareAllFiles(t, files)
}
