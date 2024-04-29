<!--
SPDX-FileCopyrightText: 2024 SUSE LLC

SPDX-License-Identifier: Apache-2.0
-->

[![REUSE status](https://api.reuse.software/badge/git.fsfe.org/reuse/api)](https://api.reuse.software/info/git.fsfe.org/reuse/api)

Check for common mistakes in gettext localizable strings in go code.

## Usage

Run the action against a go source tree to find the possible issues in localizable strings.

### Example workflow

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master
    - name: Run action
      uses: cbosdo/gettext-go-lint@master
```

### Inputs

| Input | Description |
|-------|-------------|
| `keywords` _(optional)_ | A comma-separated list of function names from which to read localizable strings] |
| `sources` _(optional)_ | A comma-separated list of files and folders containing go files to analyze |

### Outputs

The reported errors are printed on the error output and the exit code indicates the number of errors found.
There are no other outputs.

## Examples

### Using default gettext-go functions in `/code` directory

```yaml
with:
  - sources: /code
```

### Using `N` and `NL` as gettext functions in `/code1` and `/code2` directories

```yaml
with:
  - sources: /code1,/code2
  - keywords: NL,L
```
