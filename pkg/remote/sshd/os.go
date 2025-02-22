// Copyright 2022 The envd Authors
// Copyright 2022 The okteto remote Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sshd

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

var (
	errNoShell = fmt.Errorf("failed to find any shell in the PATH")

	shells = []string{
		"zsh",
		"bash",
		"sh",
	}
)

// GetShell returns the shell in $PATH.
func GetShell() (string, error) {
	for _, shell := range shells {
		if path, err := exec.LookPath(shell); err == nil {
			logrus.Infof("%s exists at %s", shell, path)
			return shell, nil
		}
		logrus.Debugf("%s does not exist", shell)
	}

	return "", errNoShell
}
