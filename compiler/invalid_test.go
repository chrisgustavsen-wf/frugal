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

func TestInvalid(t *testing.T) {
	options := compiler.Options{
		File:  idl("invalid.frugalz"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if compiler.Compile(options) == nil {
		t.Fatal("Expected error")
	}
}

func TestDuplicateServices(t *testing.T) {
	options := compiler.Options{
		File:  idl("duplicate_services.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if compiler.Compile(options) == nil {
		t.Fatal("Expected error")
	}
}

func TestDuplicateScopes(t *testing.T) {
	options := compiler.Options{
		File:  idl("duplicate_scopes.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if compiler.Compile(options) == nil {
		t.Fatal("Expected error")
	}
}

func TestDuplicateMethods(t *testing.T) {
	options := compiler.Options{
		File:  idl("duplicate_methods.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if compiler.Compile(options) == nil {
		t.Fatal("Expected error")
	}
}

func TestDuplicateOperations(t *testing.T) {
	options := compiler.Options{
		File:  idl("duplicate_operations.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if compiler.Compile(options) == nil {
		t.Fatal("Expected error")
	}
}

func TestDuplicateMethodArgIds(t *testing.T) {
	options := compiler.Options{
		File:  idl("duplicate_arg_ids.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if compiler.Compile(options) == nil {
		t.Fatal("Expected error")
	}
}

func TestDuplicateStructFieldIds(t *testing.T) {
	options := compiler.Options{
		File:  idl("duplicate_field_ids.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if compiler.Compile(options) == nil {
		t.Fatal("Expected error")
	}
}

// Ensures an error is returned when a "*" namespace has a vendor annotation.
func TestWildcardNamespaceWithVendorAnnotation(t *testing.T) {
	options := compiler.Options{
		File:  idl("bad_namespace.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err == nil {
		t.Fatal("Expected error")
	}
}

// Ensures an error is returned when a scope operation has an invalid type.
func TestInvalidScopeOperationType(t *testing.T) {
	options := compiler.Options{
		File:  idl("ad_op_type.frugal"),
		Gen:   "go",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err == nil {
		t.Fatal("Expected error")
	}
}
