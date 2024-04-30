// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package rules

import "regexp"

type multiUnnamedVariables struct {
	BaseRule
	matcher *regexp.Regexp
}

func (r multiUnnamedVariables) Check(str string) bool {
	return len(r.matcher.FindAllString(str, -1)) > 1
}

func MultiUnnamedVariables() Rule {
	return multiUnnamedVariables{
		BaseRule: BaseRule{
			name: "multi-unnamed-variables",
			description: `Multiple unnamed placeholders in a string.
Replace %s with %[n]s where n indicates the argument position starting from 1.`,
		},
		matcher: regexp.MustCompile("%[a-zA-Z0-9.]+"),
	}

}
