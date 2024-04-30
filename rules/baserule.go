// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package rules

type CheckFn func(str string, filePos string) (bool, string)

type BaseRule struct {
	name        string
	fn          CheckFn
	description string
}

func NewBaseRule(name string, description string, fn CheckFn) BaseRule {
	return BaseRule{name: name, description: description, fn: fn}
}

func (r BaseRule) Name() string {
	return r.name
}

func (r BaseRule) Description() string {
	return r.description
}

func (r BaseRule) Check(str string, filePos string) (bool, string) {
	return r.fn(str, filePos)
}
