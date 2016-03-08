/*
 * Copyright 2016 ThoughtWorks, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package libgocdgolangagent

import (
	"encoding/json"
)

type CommandTest struct {
	Command     *BuildCommand
	Expectation bool
}

type BuildCommand struct {
	Name, RunIfConfig, WorkingDirectory string
	Test                                *CommandTest
	Args                                []interface{}
	SubCommands                         []*BuildCommand
}

func MakeBuildCommand(command map[string]interface{}) *BuildCommand {
	var cmd BuildCommand
	str, _ := json.Marshal(command)
	json.Unmarshal(str, &cmd)
	return &cmd
}