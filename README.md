<!--
SPDX-License-Identifier: CC-BY-NC-SA-4.0
-->

# gh-pr-list

A simple, lightweight wrapper around `gh pr list` that supports simultaneous template application and filtering.

# The Why
As [this several-year-issue](https://github.com/cli/cli/issues/8415) reports, [`jq` filters](https://jqlang.org/) and [Go templates](https://pkg.go.dev/text/template) cannot be simultaneously applied during a single call to `gh pr list`.

This extension was created to fill that niche, alongside providing other assorted utilities for listing and formatting PRs.

# Installation
```bash
gh extension install Bertie690/gh-pr-list
```

If you have a preferred formatting, you can set that as an alias:
```bash
gh alias set show-prs $'pr list \'map(select(.mergeable == "CONFLICTING"))\' \'{{range .}}{{tablerow ((autocolor (colorstate .) (printf "#%v" .number)) | hyperlink .url) (truncate 50 .title) .headRefName (timeago .updatedAt)}}{{end}}\''
```

> [!CAUTION]
> Make sure to use strong quoting when passing a template to avoid unintended shell expansion!

# Template Functions

A few helper functions are also added for use inside templates (in addition to the ones offered by `gh pr list`).
The full list is as follows:

 - `colorhex`: Colors a string based on the provided hex code
 - `colorstate`: Returns a color string passable to the `color` or `autocolor` functions based on a PR's state and draft status. Use like so: \
    `(printf "#%v" .number) | autocolor (colorstate .)`

# Licensing
This repository seeks to be [REUSE compliant](https://reuse.software/), meaning that copyright and/or licensing information for each file is stored
either in the file itself or in an associated `REUSE.toml` file.

In summary:
- All source code belonging to the project (unless otherwise noted) is licensed under [GPL-v3.0-or-later](LICENSES/GPL-3.0-or-later.txt).
- All documentation (including this README), as well as any documentation comments explicitly documenting source code, are all licensed under [CC-BY-NC-SA-4.0](LICENSES/CC-BY-NC-SA-4.0.txt)
- Auto-generated files produced by external tools or files of insigifnicant originality are not copyrighted and are licensed under [CC0-1.0](LICENSES/CC0-1.0.txt)

# Contributing
If you have any ideas for improvement, please fill out a GitHub Issue describing it.

<!-- TODO: Write contributing.md for contributing -->

This is my first open-source project, so any and all support is greatly appreciated!
