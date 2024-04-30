// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/cbosdo/gettext-go-lint/rules"
)

var allRules []rules.Rule

func init() {
	allRules = []rules.Rule{
		rules.MultiUnnamedVariables(),
		rules.Ellipsis(),
		rules.Similar(),
	}
}
