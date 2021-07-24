# GetGitIgnore

Tool to get Git Ignore files from [GitHub GitIgnore](https://github.com/github/gitignore)

## Install

`brew install joeds13/tap/getgitignore`

Or `brew tap joeds13/tap` and then `brew install getgitignore`

## Features

* Outputs to stdout by default
* list subcommand - lists all gitignores
* search subcommand - searches gitignores

## TODO

* arg to inteligently append to file
    * don't duplicate if header comment found
* no get args == inteligent mode - finds types of files and gets appropriate ignores
* add in headers comment with commit hash above each ignore content
* Follows structure in github/gitignore
* Use Go Releaser
* Use Github Actions
* Add to Brew
* If finds a GITHUB_TOKEN env var, try use it
* Do a case insesitive search
* Move cli logic out of gitignore package
* If it's go:
    * guess the compiled binary name and add to ignore
    * remove comment against vendor/
* remove panics
* Case insensitive search
* rename to "gitignore"?
