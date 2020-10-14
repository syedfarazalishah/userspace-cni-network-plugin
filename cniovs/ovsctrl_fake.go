// Copyright 2020 Intel Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cniovs

//
// Fake implementation of execCommand suitable for unit testing
//

type FakeExecCommand struct {
	Out  []byte
	Err  error
	Cmd  string
	Args []string
}

func (e *FakeExecCommand) execCommand(cmd string, args []string) ([]byte, error) {
	e.Cmd = cmd
	e.Args = args
	return e.Out, e.Err
}
