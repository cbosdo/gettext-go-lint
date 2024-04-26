// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

type Flags struct {
	Keywords []string
}

func main() {
	var flags Flags

	name := path.Base(os.Args[0])
	cmd := &cobra.Command{
		Use:          name + " file.go ...",
		Long:         "Check localizable strings for common issues",
		Version:      "0.0.1",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			os.Exit(lint(&flags, args))
		},
	}

	cmd.Flags().StringArrayVarP(&flags.Keywords, "keyword", "k", []string{}, "defines the functions matching the localizable strings")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
