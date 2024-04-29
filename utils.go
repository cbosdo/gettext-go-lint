// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

func resolveFiles(sources []string) []string {
	var files []string

	for _, source := range sources {
		if isFile(source) && isGo(source) {
			files = append(files, source)
		} else if isDir(source) {
			list, err := os.ReadDir(source)
			if err != nil {
				fmt.Printf("Failed to read directory %s", source)
				continue
			}
			var dirFiles []string
			for _, item := range list {
				dirFiles = append(dirFiles, path.Join(source, item.Name()))
			}
			files = append(files, resolveFiles(dirFiles)...)
		}
	}

	return files
}

func isFile(filename string) bool {
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

func isDir(filename string) bool {
	info, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return info.IsDir()
}

func isGo(filename string) bool {
	return strings.HasSuffix(filename, ".go") && !strings.HasSuffix(filename, "_test.go")
}

func getEnvList(name string) []string {
	value := os.Getenv(name)
	if value == "" {
		return []string{}
	}
	return strings.Split(value, ",")
}
