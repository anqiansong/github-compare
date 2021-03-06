// MIT License
//
// Copyright (c) 2022 anqiansong
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

type style = string

const (
	codeFailure        = 1
	defaultEmptyString = ""
	flagFile           = "file"
	flagFileShortHand  = "f"
	flagToken          = "token"
	flagTokenShortHand = "t"
	rootCMDDesc        = "A GitHub repositories statistics command-line tool for the terminal"
	flagTokenDesc      = "github access token"
	flagTermUIDesc     = "print with term ui style(default)"
	flagJSONDesc       = "print with json style"
	flagYAMLDesc       = "print with yaml style"
	flagFileDesc       = "output to a specified file"

	styleJSON   style = "json"
	styleYAML   style = "yaml"
	styleTermUI style = "ui"
)
