// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package rules

import "strings"

func Ellipsis() Rule {
	return NewBaseRule(
		"ellipsis",
		"Replace the three dots with the ellipsis characters '…'",
		func(str string) bool {
			return strings.Contains(str, "...")
		},
	)
}
