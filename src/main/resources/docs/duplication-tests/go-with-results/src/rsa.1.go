// (c) Copyright 2016 Hewlett Packard Enterprise Development LP
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

package rules

import (
	"fmt"
	"go/ast"

	"github.com/securego/gosec"
)

type weakKeyStrength struct {
	gosec.MetaData
	calls gosec.CallList
	bits  int
}

func (w *weakKeyStrength) Match(n ast.Node, c *gosec.Context) (*gosec.Issue, error) {
	if callExpr := w.calls.ContainsCallExpr(n, c); callExpr != nil {
		if bits, err := gosec.GetInt(callExpr.Args[1]); err == nil && bits < (int64)(w.bits) {
			return gosec.NewIssue(c, n, w.ID(), w.What, w.Severity, w.Confidence), nil
		}
	}
	return nil, nil
}