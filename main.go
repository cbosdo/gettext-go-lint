// SPDX-FileCopyrightText: 2024 SUSE LLC
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
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
			sources := getEnvList("INPUT_SOURCES")
			sources = append(sources, args...)
			if len(sources) == 0 {
				pwd, err := os.Getwd()
				if err != nil {
					log.Fatalf("No sources provided and failed to get current working directory: %s", err)
				}
				if pwd == "/" {
					log.Fatalln("Cannot add / as sources")
				}
				sources = append(sources, pwd)
			}

			keywords := append(flags.Keywords, getEnvList("INPUT_KEYWORDS")...)
			os.Exit(lint(keywords, sources))
		},
	}

	cmd.AddCommand(newDescribeCommand())

	var defaultKeywords []string
	for _, domainPrefix := range []string{"", "D"} {
		for _, prefix := range []string{"", "P", "N", "PN"} {
			defaultKeywords = append(defaultKeywords, domainPrefix+prefix+"Gettext")
		}
	}
	cmd.Flags().StringSliceVarP(&flags.Keywords, "keyword", "k", defaultKeywords, "defines the functions matching the localizable strings")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func newDescribeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "describe <rule>",
		Long: "Describe a rule using the Identifier shown in the error output",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, rule := range rules {
				if rule.Name == args[0] {
					fmt.Println(rule.Description)
					os.Exit(0)
				}
			}
			os.Exit(1)
		},
	}
	return cmd
}
