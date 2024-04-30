// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package rules

type Rule interface {
	Name() string
	Check(str string, filePos string) (bool, string)
	Description() string
}
