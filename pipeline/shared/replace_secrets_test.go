// Copyright 2023 Woodpecker Authors
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

package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSecretsReplacer(t *testing.T) {
	tc := []struct {
		log     string
		secrets []string
		expect  string
	}{{
		log:     "start log\ndone",
		secrets: []string{""},
		expect:  "start log\ndone",
	}, {
		log:     `this IS secret: password`,
		secrets: []string{"password", " IS "},
		expect:  `this ******** secret: ********`,
	}, {
		log:     "start log\ndone\nnow\nan\nmulti line secret!! ;)",
		secrets: []string{"an\nmulti line secret!!"},
		expect:  "start log\ndone\nnow\n********\n******** ;)",
	}, {
		log:     "start log\ndone\nnow\nan\nmulti line secret!! ;)",
		secrets: []string{"Test\nwith\n\ntwo new lines"},
		expect:  "start log\ndone\nnow\nan\nmulti line secret!! ;)",
	}, {
		log:     "start log\ndone\nnow\nan\nmulti line secret!! ;)\nwith\ntwo\n\nnewlines",
		secrets: []string{"an\nmulti line secret!!", "two\n\nnewlines"},
		expect:  "start log\ndone\nnow\n********\n******** ;)\nwith\n********\n\n********",
	}}

	for _, c := range tc {
		rep := NewSecretsReplacer(c.secrets)
		result := rep.Replace(c.log)
		assert.EqualValues(t, c.expect, result)
	}
}
