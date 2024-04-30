// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package rules

import (
	"fmt"
	"regexp"
	"strings"
)

type similar struct {
	collected map[string]stringRecord
	matcher   *regexp.Regexp
}

type stringRecord struct {
	msgid    string
	location string
}

func Similar() Rule {
	return similar{
		collected: map[string]stringRecord{},
		matcher:   regexp.MustCompile("[^a-z0-9%]"),
	}
}

func (r similar) Name() string {
	return "similar"
}

func (r similar) Description() string {
	return "Strings that could be duplicates of another one."
}

func (r similar) Check(str string, filePos string) (bool, string) {
	cleanedStr := r.matcher.ReplaceAllString(strings.ToLower(str), "")
	for key, record := range r.collected {
		if cleanedStr == key && str != record.msgid {
			return true, fmt.Sprintf(" similar to %s in %s", record.msgid, record.location)
		}
	}
	r.collected[cleanedStr] = stringRecord{
		msgid:    str,
		location: filePos,
	}
	return false, ""
}
