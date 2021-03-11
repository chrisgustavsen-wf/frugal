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

	"github.com/stretchr/testify/assert"

	"github.com/Workiva/frugal/compiler"
)

func TestInvalid(t *testing.T) {
	suite := [][2]string{
		{"invalid.frugal", "parser: syntax error"},
		{"duplicate_services.frugal", "Services foo and Foo conflict"},
		{"duplicate_scopes.frugal", "Scopes foo and Foo conflict"},
		{"duplicate_methods.frugal", "Methods Ping and ping conflict"},
		{"duplicate_operations.frugal", "Operations boo and Boo conflict"},
		{"duplicate_arg_ids.frugal", "Duplicate field"},
		{"duplicate_field_ids.frugal", "Duplicate field"},
		{"bad_namespace.frugal", "annotation not compatible with * namespace"},
		{"bad_op_type.frugal", "Invalid operation type invalid.type for test.test"},
	}
	for _, test := range suite {
		options := compiler.Options{
			File:  idl(test[0]),
			Gen:   "go",
			Out:   outputDir,
			Delim: delim,
		}
		err := compiler.Compile(options)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), test[1])
	}
}
