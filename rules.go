// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"regexp"
	"strings"
)

type Rule struct {
	Name        string
	Fn          func(str string) bool
	Description string
}

var rules []Rule

var multiUnnamed *regexp.Regexp

func init() {
	multiUnnamed = regexp.MustCompile("%[a-zA-Z0-9.]+")
	rules = []Rule{
		{
			Name: "multi-unnamed-variables",
			Fn: func(str string) bool {
				return len(multiUnnamed.FindAllString(str, -1)) > 1
			},
			Description: `Multiple unnamed placeholders in a string.
Replace %s with %[n]s where n indicates the argument position starting from 1.`,
		},
		{
			Name: "ellipsis",
			Fn: func(str string) bool {
				return strings.Contains(str, "...")
			},
			Description: "Replace the three dots with the ellipsis characters '…'",
		},
	}
}
