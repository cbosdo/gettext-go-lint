// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"regexp"

	"github.com/cbosdo/gettext-go-lint/rules"
)

var allRules []rules.Rule

var multiUnnamed *regexp.Regexp

func init() {
	multiUnnamed = regexp.MustCompile("%[a-zA-Z0-9.]+")
	allRules = []rules.Rule{
		rules.MultiUnnamedVariables(),
		rules.Ellipsis(),
	}
}
