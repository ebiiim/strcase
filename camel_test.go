/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strcase

import (
	"testing"
)

func TestToCamelIncludePeriods(t *testing.T) {
	cases := []struct {
		in      string
		isLower bool
		out     string
	}{
		{"test_case", false, "TestCase"},
		{"test", false, "Test"},
		{"TestCase", false, "TestCase"},
		{" test  case ", false, "TestCase"},
		{"", false, ""},
		{"many_many_words", false, "ManyManyWords"},
		{"AnyKind of_string", false, "AnyKindOfString"},
		{"odd-fix", false, "OddFix"},
		{"numbers2And55with000", false, "Numbers2And55With000"},
		{"foo-bar", true, "fooBar"},
		{"TestCase", true, "testCase"},
		{"", true, ""},
		{"AnyKind of_string", true, "anyKindOfString"},
		{"AnyKind.of.string", false, "AnyKindOfString"},
		{"AnyKind.of.string", true, "anyKindOfString"},
	}
	for _, c := range cases {
		c := c
		result := ToCamelIncludePeriods(c.in, c.isLower)
		if result != c.out {
			t.Error("'" + result + "' != '" + c.out + "'")
		}
	}
}

func TestToCamel(t *testing.T) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToCamel(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func TestToLowerCamel(t *testing.T) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToLowerCamel(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
