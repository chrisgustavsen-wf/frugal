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

func TestValidDartFrugalCompiler(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "dart",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("unexpected error", err)
	}

	files := []FileComparisonPair{
		{"dart/variety/f_awesome_exception.dart", "variety/lib/src/f_awesome_exception.dart"},
		{"dart/variety/f_event.dart", "variety/lib/src/f_event.dart"},
		{"dart/variety/f_event_wrapper.dart", "variety/lib/src/f_event_wrapper.dart"},
		{"dart/variety/f_its_an_enum.dart", "variety/lib/src/f_its_an_enum.dart"},
		{"dart/variety/f_test_base.dart", "variety/lib/src/f_test_base.dart"},
		{"dart/variety/f_testing_defaults.dart", "variety/lib/src/f_testing_defaults.dart"},
		{"dart/variety/f_testing_unions.dart", "variety/lib/src/f_testing_unions.dart"},
		{"dart/variety/f_health_condition.dart", "variety/lib/src/f_health_condition.dart"},
		{"dart/variety/f_test_lowercase.dart", "variety/lib/src/f_test_lowercase.dart"},
		{"dart/variety/f_foo_args.dart", "variety/lib/src/f_foo_args.dart"},
		{"dart/variety/f_variety_constants.dart", "variety/lib/src/f_variety_constants.dart"},
		{"dart/variety/f_events_scope.dart", "variety/lib/src/f_events_scope.dart"},
		{"dart/variety/f_foo_service.dart", "variety/lib/src/f_foo_service.dart"},
		{"dart/variety/variety.dart", "variety/lib/variety.dart"},

		{"dart/actual_base/f_actual_base_dart_constants.dart", "actual_base_dart/lib/src/f_actual_base_dart_constants.dart"},
		{"dart/actual_base/f_api_exception.dart", "actual_base_dart/lib/src/f_api_exception.dart"},
		{"dart/actual_base/f_thing.dart", "actual_base_dart/lib/src/f_thing.dart"},
		{"dart/actual_base/f_base_health_condition.dart", "actual_base_dart/lib/src/f_base_health_condition.dart"},
		{"dart/actual_base/f_base_foo_service.dart", "actual_base_dart/lib/src/f_base_foo_service.dart"},
		{"dart/actual_base/f_nested_thing.dart", "actual_base_dart/lib/src/f_nested_thing.dart"},
		{"dart/actual_base/actual_base_dart.dart", "actual_base_dart/lib/actual_base_dart.dart"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

func TestValidDartUseNullForUnset(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "dart:use_null_for_unset",
		Out:     outputDir + "/nullunset",
		Delim:   delim,
		Recurse: true,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("unexpected error", err)
	}

	files := []FileComparisonPair{
		{"dart.nullunset/variety/f_awesome_exception.dart", "nullunset/variety/lib/src/f_awesome_exception.dart"},
		{"dart.nullunset/variety/f_event.dart", "nullunset/variety/lib/src/f_event.dart"},
		{"dart.nullunset/variety/f_event_wrapper.dart", "nullunset/variety/lib/src/f_event_wrapper.dart"},
		{"dart.nullunset/variety/f_its_an_enum.dart", "nullunset/variety/lib/src/f_its_an_enum.dart"},
		{"dart.nullunset/variety/f_test_base.dart", "nullunset/variety/lib/src/f_test_base.dart"},
		{"dart.nullunset/variety/f_testing_defaults.dart", "nullunset/variety/lib/src/f_testing_defaults.dart"},
		{"dart.nullunset/variety/f_testing_unions.dart", "nullunset/variety/lib/src/f_testing_unions.dart"},
		{"dart.nullunset/variety/f_health_condition.dart", "nullunset/variety/lib/src/f_health_condition.dart"},
		{"dart.nullunset/variety/f_test_lowercase.dart", "nullunset/variety/lib/src/f_test_lowercase.dart"},
		{"dart.nullunset/variety/f_foo_args.dart", "nullunset/variety/lib/src/f_foo_args.dart"},
		{"dart.nullunset/variety/f_variety_constants.dart", "nullunset/variety/lib/src/f_variety_constants.dart"},
		{"dart.nullunset/variety/f_events_scope.dart", "nullunset/variety/lib/src/f_events_scope.dart"},
		{"dart.nullunset/variety/f_foo_service.dart", "nullunset/variety/lib/src/f_foo_service.dart"},
		{"dart.nullunset/variety/variety.dart", "nullunset/variety/lib/variety.dart"},

		{"dart.nullunset/actual_base/f_actual_base_dart_constants.dart", "nullunset/actual_base_dart/lib/src/f_actual_base_dart_constants.dart"},
		{"dart.nullunset/actual_base/f_api_exception.dart", "nullunset/actual_base_dart/lib/src/f_api_exception.dart"},
		{"dart.nullunset/actual_base/f_thing.dart", "nullunset/actual_base_dart/lib/src/f_thing.dart"},
		{"dart.nullunset/actual_base/f_base_health_condition.dart", "nullunset/actual_base_dart/lib/src/f_base_health_condition.dart"},
		{"dart.nullunset/actual_base/f_base_foo_service.dart", "nullunset/actual_base_dart/lib/src/f_base_foo_service.dart"},
		{"dart.nullunset/actual_base/f_nested_thing.dart", "nullunset/actual_base_dart/lib/src/f_nested_thing.dart"},
		{"dart.nullunset/actual_base/actual_base_dart.dart", "nullunset/actual_base_dart/lib/actual_base_dart.dart"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

func TestValidDartEnums(t *testing.T) {
	options := compiler.Options{
		File:    idl("enum.frugal"),
		Gen:     "dart:use_enums",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("unexpected error", err)
	}

	files := []FileComparisonPair{
		{"dart/enum/f_testing_enums.dart", "enum_dart/lib/src/f_testing_enums.dart"},
		{"dart/enum/enum_dart.dart", "enum_dart/lib/enum_dart.dart"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

// Ensures correct import references are used when -use-vendor is set and the
// IDL has a vendored include
func TestValidDartVendor(t *testing.T) {
	options := compiler.Options{
		File:  includeVendor,
		Gen:   "dart:use_vendor",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"dart/include_vendor/f_my_scope_scope.dart", "include_vendor/lib/src/f_my_scope_scope.dart"},
		{"dart/include_vendor/f_my_service_service.dart", "include_vendor/lib/src/f_my_service_service.dart"},
		{"dart/include_vendor/f_vendored_references.dart", "include_vendor/lib/src/f_vendored_references.dart"},
		{"dart/include_vendor/include_vendor.dart", "include_vendor/lib/include_vendor.dart"},
		{"dart/include_vendor/pubspec.yaml", "include_vendor/pubspec.yaml"},
	}

	copyAllFiles(t, files)
	compareAllFiles(t, files)
}

// Ensures an error is returned when -use-vendor is set and the vendored
// include does not specify a path.
func TestValidDartVendorPathNotSpecified(t *testing.T) {
	options := compiler.Options{
		File:  includeVendorNoPath,
		Gen:   "dart:use_vendor",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err == nil {
		t.Fatal("Expected error")
	}
}

// Ensures the target IDL is generated when -use-vendor is set and it has a
// vendored namespace.
func TestValidDartVendorNamespaceTargetGenerate(t *testing.T) {
	options := compiler.Options{
		File:  vendorNamespace,
		Gen:   "dart:use_vendor",
		Out:   outputDir,
		Delim: delim,
	}
	if err := compiler.Compile(options); err != nil {
		t.Fatal("Unexpected error", err)
	}

	files := []FileComparisonPair{
		{"dart/vendor_namespace/vendor_namespace.dart", "vendor_namespace/lib/vendor_namespace.dart"},
		{"dart/vendor_namespace/f_item.dart", "vendor_namespace/lib/src/f_item.dart"},
		{"dart/vendor_namespace/f_vendored_base_service.dart", "vendor_namespace/lib/src/f_vendored_base_service.dart"},
		{"dart/vendor_namespace/f_vendor_namespace_constants.dart", "vendor_namespace/lib/src/f_vendor_namespace_constants.dart"},
		{"dart/vendor_namespace/f_my_enum.dart", "vendor_namespace/lib/src/f_my_enum.dart"},
	}
	copyAllFiles(t, files)
	compareAllFiles(t, files)
}
