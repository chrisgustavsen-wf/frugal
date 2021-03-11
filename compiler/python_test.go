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
	"path/filepath"
	"testing"

	"github.com/Workiva/frugal/compiler"
)

func TestValidPython(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "py",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"python/variety/__init__.py", "variety/python/__init__.py"},
		{"python/variety/constants.py", "variety/python/constants.py"},
		{"python/variety/ttypes.py", "variety/python/ttypes.py"},
		{"python/variety/f_Events_publisher.py", "variety/python/f_Events_publisher.py"},
		{"python/variety/f_Events_subscriber.py", "variety/python/f_Events_subscriber.py"},
		{"python/variety/f_Foo.py", "variety/python/f_Foo.py"},

		{"python/actual_base/__init__.py", "actual_base/python/__init__.py"},
		{"python/actual_base/constants.py", "actual_base/python/constants.py"},
		{"python/actual_base/ttypes.py", "actual_base/python/ttypes.py"},
		{"python/actual_base/f_BaseFoo.py", "actual_base/python/f_BaseFoo.py"},
	}
	suite.Run(t, options)
}

func TestValidPythonTornado(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "py:tornado",
		Out:     outputDir + "/tornado",
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"python.tornado/variety/__init__.py", "tornado/variety/python/__init__.py"},
		{"python.tornado/variety/constants.py", "tornado/variety/python/constants.py"},
		{"python.tornado/variety/ttypes.py", "tornado/variety/python/ttypes.py"},
		{"python.tornado/variety/f_Events_publisher.py", "tornado/variety/python/f_Events_publisher.py"},
		{"python.tornado/variety/f_Events_subscriber.py", "tornado/variety/python/f_Events_subscriber.py"},
		{"python.tornado/variety/f_Foo.py", "tornado/variety/python/f_Foo.py"},

		{"python.tornado/actual_base/__init__.py", "tornado/actual_base/python/__init__.py"},
		{"python.tornado/actual_base/constants.py", "tornado/actual_base/python/constants.py"},
		{"python.tornado/actual_base/ttypes.py", "tornado/actual_base/python/ttypes.py"},
		{"python.tornado/actual_base/f_BaseFoo.py", "tornado/actual_base/python/f_BaseFoo.py"},
	}
	suite.Run(t, options)
}

func TestValidPythonAsyncIO(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "py:asyncio",
		Out:     outputDir + "/asyncio",
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"python.asyncio/variety/__init__.py", "asyncio/variety/python/__init__.py"},
		{"python.asyncio/variety/constants.py", "asyncio/variety/python/constants.py"},
		{"python.asyncio/variety/ttypes.py", "asyncio/variety/python/ttypes.py"},
		{"python.asyncio/variety/f_Events_publisher.py", "asyncio/variety/python/f_Events_publisher.py"},
		{"python.asyncio/variety/f_Events_subscriber.py", "asyncio/variety/python/f_Events_subscriber.py"},
		{"python.asyncio/variety/f_Foo.py", "asyncio/variety/python/f_Foo.py"},

		{"python.asyncio/actual_base/__init__.py", "asyncio/actual_base/python/__init__.py"},
		{"python.asyncio/actual_base/constants.py", "asyncio/actual_base/python/constants.py"},
		{"python.asyncio/actual_base/ttypes.py", "asyncio/actual_base/python/ttypes.py"},
		{"python.asyncio/actual_base/f_BaseFoo.py", "asyncio/actual_base/python/f_BaseFoo.py"},
	}
	suite.Run(t, options)
}

func TestPythonPackagePrefix(t *testing.T) {
	options := compiler.Options{
		File:    idl("service_inheritance.frugal"),
		Gen:     "py:package_prefix=generic_package_prefix.",
		Out:     outputDir,
		Delim:   delim,
		Recurse: false,
	}
	suite := ComparisonList{
		{"python/package_prefix/f_Foo.py", "service_inheritance/f_Foo.py"},
		{"python/package_prefix/ttypes.py", "service_inheritance/ttypes.py"},
		{"python/package_prefix/constants.py", "service_inheritance/constants.py"},
	}
	suite.Run(t, options)
}

func TestPythonExtendServiceSameFile(t *testing.T) {
	options := compiler.Options{
		File:  idl("service_extension_same_file.frugal"),
		Gen:   "py:asyncio",
		Out:   outputDir,
		Delim: delim,
	}
	suite := ComparisonList{
		{"python.asyncio/service_extension_same_file/f_BasePinger.py", "service_extension_same_file/python/f_BasePinger.py"},
		{"python.asyncio/service_extension_same_file/f_Pinger.py", "service_extension_same_file/python/f_Pinger.py"},
	}
	suite.Run(t, options)
}

func TestPythonAbsoluteOutputPath(t *testing.T) {
	absoluteOutputDir, err := filepath.Abs("testdata/out/absolute_path")
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "py",
		Out:     absoluteOutputDir,
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"python/variety/__init__.py", absoluteOutputDir + "/variety/python/__init__.py"},
		{"python/variety/constants.py", absoluteOutputDir + "/variety/python/constants.py"},
		{"python/variety/ttypes.py", absoluteOutputDir + "/variety/python/ttypes.py"},
		{"python/variety/f_Events_publisher.py", absoluteOutputDir + "/variety/python/f_Events_publisher.py"},
		{"python/variety/f_Events_subscriber.py", absoluteOutputDir + "/variety/python/f_Events_subscriber.py"},
		{"python/variety/f_Foo.py", absoluteOutputDir + "/variety/python/f_Foo.py"},

		{"python/actual_base/__init__.py", absoluteOutputDir + "/actual_base/python/__init__.py"},
		{"python/actual_base/constants.py", absoluteOutputDir + "/actual_base/python/constants.py"},
		{"python/actual_base/ttypes.py", absoluteOutputDir + "/actual_base/python/ttypes.py"},
		{"python/actual_base/f_BaseFoo.py", absoluteOutputDir + "/actual_base/python/f_BaseFoo.py"},
	}
	suite.Run(t, options)
}
