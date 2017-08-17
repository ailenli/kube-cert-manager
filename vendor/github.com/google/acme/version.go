// Copyright 2015 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build release

package main

import "fmt"

var (
	version = "1.1.1"
	osarch  string // set by ldflags

	cmdVersion = &command{
		run:       runVersion,
		UsageLine: "version",
		Short:     "display acme tool version",
		Long: `
The version command is available only for releases.
The binary built with "go get" will not have the version subcommand.
		`,
	}
)

func init() {
	// Insert "acme version" at the top of the commands.
	commands = append([]*command{cmdVersion}, commands...)
}

func runVersion(args []string) {
	fmt.Println(version, osarch)
}
