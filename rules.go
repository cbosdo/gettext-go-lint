// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"regexp"
	"strings"

	"github.com/cbosdo/gettext-go-lint/rules"
)

type Rule interface {
	Name() string
	Check(str string) bool
	Description() string
}

var allRules []Rule

var multiUnnamed *regexp.Regexp

func init() {
	multiUnnamed = regexp.MustCompile("%[a-zA-Z0-9.]+")
	allRules = []Rule{
		rules.NewBaseRule(
			"multi-unnamed-variables",
			`Multiple unnamed placeholders in a string.
Replace %s with %[n]s where n indicates the argument position starting from 1.`,
			func(str string) bool {
				return len(multiUnnamed.FindAllString(str, -1)) > 1
			},
		),
		rules.NewBaseRule(
			"ellipsis",
			"Replace the three dots with the ellipsis characters '…'",
			func(str string) bool {
				return strings.Contains(str, "...")
			},
		),
	}
}
