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
	"time"

	"github.com/Workiva/frugal/compiler"
	"github.com/Workiva/frugal/compiler/globals"
)

func TestValidJavaWithAsync(t *testing.T) {
	nowBefore := globals.Now
	defer func() {
		globals.Now = nowBefore
	}()
	globals.Now = time.Date(2015, 11, 24, 0, 0, 0, 0, time.UTC)

	options := compiler.Options{
		File:  frugalGenFile,
		Gen:   "java:async",
		Out:   outputDir + "/async",
		Delim: delim,
	}
	suite := ComparisonList{
		{"java/variety_async/FFoo.java", "async/variety/java/FFoo.java"},
	}
	suite.Run(t, options)
}

func TestValidJavaFrugalCompiler(t *testing.T) {
	defer globals.Reset()
	nowBefore := globals.Now
	defer func() {
		globals.Now = nowBefore
	}()
	globals.Now = time.Date(2015, 11, 24, 0, 0, 0, 0, time.UTC)

	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "java",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"java/variety/AwesomeException.java", "variety/java/AwesomeException.java"},
		{"java/variety/Event.java", "variety/java/Event.java"},
		{"java/variety/EventWrapper.java", "variety/java/EventWrapper.java"},
		{"java/variety/ItsAnEnum.java", "variety/java/ItsAnEnum.java"},
		{"java/variety/TestBase.java", "variety/java/TestBase.java"},
		{"java/variety/TestingDefaults.java", "variety/java/TestingDefaults.java"},
		{"java/variety/TestingUnions.java", "variety/java/TestingUnions.java"},
		{"java/variety/HealthCondition.java", "variety/java/HealthCondition.java"},
		{"java/variety/varietyConstants.java", "variety/java/varietyConstants.java"},
		{"java/variety/TestLowercase.java", "variety/java/TestLowercase.java"},
		{"java/variety/FooArgs.java", "variety/java/FooArgs.java"},
		{"java/variety/EventsPublisher.java", "variety/java/EventsPublisher.java"},
		{"java/variety/EventsSubscriber.java", "variety/java/EventsSubscriber.java"},
		{"java/variety/FFoo.java", "variety/java/FFoo.java"},

		{"java/actual_base/api_exception.java", "actual_base/java/api_exception.java"},
		{"java/actual_base/baseConstants.java", "actual_base/java/baseConstants.java"},
		{"java/actual_base/thing.java", "actual_base/java/thing.java"},
		{"java/actual_base/base_health_condition.java", "actual_base/java/base_health_condition.java"},
		{"java/actual_base/FBaseFoo.java", "actual_base/java/FBaseFoo.java"},
		{"java/actual_base/nested_thing.java", "actual_base/java/nested_thing.java"},
	}
	suite.Run(t, options)
}

func TestValidJavaBoxedPrimitives(t *testing.T) {
	defer globals.Reset()
	nowBefore := globals.Now
	defer func() {
		globals.Now = nowBefore
	}()
	globals.Now = time.Date(2015, 11, 24, 0, 0, 0, 0, time.UTC)

	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "java:boxed_primitives",
		Out:     outputDir + "/boxed_primitives",
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"java/boxed_primitives/FFoo.java", "boxed_primitives/variety/java/FFoo.java"},
		{"java/boxed_primitives/TestingDefaults.java", "boxed_primitives/variety/java/TestingDefaults.java"},
	}
	suite.Run(t, options)
}

// Ensures correct import references are used when -use-vendor is set and the
// IDL has a vendored include.
func TestValidJavaVendor(t *testing.T) {
	nowBefore := globals.Now
	defer func() {
		globals.Now = nowBefore
	}()
	globals.Now = time.Date(2015, 11, 24, 0, 0, 0, 0, time.UTC)

	options := compiler.Options{
		File:    includeVendor,
		Gen:     "java:use_vendor",
		Out:     outputDir + "/valid_vendor",
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"java/valid_vendor/FMyService.java", "valid_vendor/include_vendor/java/FMyService.java"},
		{"java/valid_vendor/MyScopePublisher.java", "valid_vendor/include_vendor/java/MyScopePublisher.java"},
		{"java/valid_vendor/MyScopeSubscriber.java", "valid_vendor/include_vendor/java/MyScopeSubscriber.java"},
		{"java/valid_vendor/VendoredReferences.java", "valid_vendor/include_vendor/java/VendoredReferences.java"},
		{"java/valid_vendor/InvalidData.java", "valid_vendor/InvalidData.java"},
	}
	suite.Run(t, options)

	filesNotToGenerate := []string{
		"valid_vendor/vendor_namespace/java/Item.java",
		"valid_vendor/vendor_namespace/java/vendor_namespaceConstants.java",
		"valid_vendor/vendor_namespace/java/MyEnum.java",
		"valid_vendor/vendor_namespace/java/FVendoredBase.java",
	}

	assertFilesNotExist(t, filesNotToGenerate)

}

func TestValidJavaVendorButNotUseVendor(t *testing.T) {
	nowBefore := globals.Now
	defer func() {
		globals.Now = nowBefore
	}()
	globals.Now = time.Date(2015, 11, 24, 0, 0, 0, 0, time.UTC)

	options := compiler.Options{
		File:    includeVendor,
		Gen:     "java",
		Out:     outputDir + "/vendored_but_no_use_vendor",
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"java/vendored_but_no_use_vendor/FMyService.java", "vendored_but_no_use_vendor/include_vendor/java/FMyService.java"},
		{"java/vendored_but_no_use_vendor/MyScopePublisher.java", "vendored_but_no_use_vendor/include_vendor/java/MyScopePublisher.java"},
		{"java/vendored_but_no_use_vendor/MyScopeSubscriber.java", "vendored_but_no_use_vendor/include_vendor/java/MyScopeSubscriber.java"},
		{"java/vendored_but_no_use_vendor/VendoredReferences.java", "vendored_but_no_use_vendor/include_vendor/java/VendoredReferences.java"},
		{"java/vendored_but_no_use_vendor/InvalidData.java", "vendored_but_no_use_vendor/InvalidData.java"},
		{"java/vendored_but_no_use_vendor/Item.java", "vendored_but_no_use_vendor/vendor_namespace/java/Item.java"},
		{"java/vendored_but_no_use_vendor/vendor_namespaceConstants.java", "vendored_but_no_use_vendor/vendor_namespace/java/vendor_namespaceConstants.java"},
		{"java/vendored_but_no_use_vendor/MyEnum.java", "vendored_but_no_use_vendor/vendor_namespace/java/MyEnum.java"},
		{"java/vendored_but_no_use_vendor/FVendoredBase.java", "vendored_but_no_use_vendor/vendor_namespace/java/FVendoredBase.java"},
	}
	suite.Run(t, options)
}

func TestValidJavaVendorNoPathUsesDefinedNamespace(t *testing.T) {
	nowBefore := globals.Now
	defer func() {
		globals.Now = nowBefore
	}()
	globals.Now = time.Date(2015, 11, 24, 0, 0, 0, 0, time.UTC)

	options := compiler.Options{
		File:    includeVendorNoPath,
		Gen:     "java:use_vendor",
		Out:     outputDir + "/valid_vendor_no_path",
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"java/valid_vendor_no_path/FMyService.java", "valid_vendor_no_path/include_vendor_no_path/java/FMyService.java"},
		{"java/valid_vendor_no_path/MyScopePublisher.java", "valid_vendor_no_path/include_vendor_no_path/java/MyScopePublisher.java"},
		{"java/valid_vendor_no_path/MyScopeSubscriber.java", "valid_vendor_no_path/include_vendor_no_path/java/MyScopeSubscriber.java"},
		{"java/valid_vendor_no_path/VendoredReferences.java", "valid_vendor_no_path/include_vendor_no_path/java/VendoredReferences.java"},
		{"java/valid_vendor_no_path/InvalidData.java", "valid_vendor_no_path/InvalidData.java"},
	}
	suite.Run(t, options)

	filesNotToGenerate := []string{
		"valid_vendor_no_path/vendor_namespace/java/Item.java",
		"valid_vendor_no_path/vendor_namespace/java/vendor_namespaceConstants.java",
		"valid_vendor_no_path/vendor_namespace/java/MyEnum.java",
		"valid_vendor_no_path/vendor_namespace/java/FVendoredBase.java",
	}

	assertFilesNotExist(t, filesNotToGenerate)
}

func TestValidJavaSuppressDeprecatedLogging(t *testing.T) {
	nowBefore := globals.Now
	defer func() {
		globals.Now = nowBefore
	}()
	globals.Now = time.Date(2015, 11, 24, 0, 0, 0, 0, time.UTC)

	options := compiler.Options{
		File:  frugalGenFile,
		Gen:   "java:suppress_deprecated_logging",
		Out:   outputDir + "/deprecated_logging",
		Delim: delim,
	}
	suite := ComparisonList{
		{"java/deprecated_logging/FFoo.java", "deprecated_logging/variety/java/FFoo.java"},
	}
	suite.Run(t, options)
}
