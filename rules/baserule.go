// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package rules

type BaseRule struct {
	name        string
	fn          func(str string) bool
	description string
}

func NewBaseRule(name string, description string, fn func(str string) bool) BaseRule {
	return BaseRule{name: name, description: description, fn: fn}
}

func (r BaseRule) Name() string {
	return r.name
}

func (r BaseRule) Description() string {
	return r.description
}

func (r BaseRule) Check(str string) bool {
	return r.fn(str)
}
