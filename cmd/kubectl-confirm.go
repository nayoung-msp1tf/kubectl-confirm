/*
Copyright 2022 Brian Pursley.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/brianpursley/kubectl-confirm/pkg/cmd"
)

func main() {
	if containsDeleteCommand(cmd) {
		fmt.Println("명령어에 delete 명령이 포함되어 있습니다.")
		
		cmd := cmd.NewConfirmCommand()
		err := cmd.Execute()
		if err != nil {
			_ = fmt.Errorf("confirm failed: %s", err)
		}	
	}
}
