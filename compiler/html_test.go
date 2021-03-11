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

func TestHTML(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "html",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"html/style.css", "style.css"},
		{"html/index.html", "index.html"},
		{"html/base.html", "base.html"},
		{"html/variety.html", "variety.html"},
	}
	suite.Run(t, options)
}

func TestHTMLStandalone(t *testing.T) {
	options := compiler.Options{
		File:    frugalGenFile,
		Gen:     "html:standalone",
		Out:     outputDir,
		Delim:   delim,
		Recurse: true,
	}
	suite := ComparisonList{
		{"html/standalone/index.html", "index.html"},
		{"html/standalone/base.html", "base.html"},
		{"html/standalone/variety.html", "variety.html"},
	}
	suite.Run(t, options)
}
