module github.com/Bertie690/gh-pr-list/magefiles

go 1.25.1

// Replace directives aren't usually recommended to be left in prod builds, but this entire directory
// is solely used for workflows so it's fine
replace github.com/Bertie690/gh-pr-list => ../

require (
	github.com/Bertie690/gh-pr-list v1.0.4
	github.com/fatih/color v1.18.0
	github.com/magefile/mage v1.15.0
	golang.org/x/text v0.29.0
)

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/cli/go-gh/v2 v2.12.2 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/muesli/termenv v0.16.0 // indirect
	github.com/nsf/jsondiff v0.0.0-20230430225905-43f6cf3098c1 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/term v0.35.0 // indirect
)
