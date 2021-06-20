package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joeds13/getgitignore/gitignore"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

const (
	usage = `usage: %s
Get GitIgnore Files

`
)

func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getFile := getCmd.Bool("file", false, "TODO: file usage")
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		getCmd.Usage()
		// listCmd.Usage()
		// searchCmd.Usage()
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		getCmd.Parse(os.Args[2:])
		if len(os.Args[2:]) != 1 {
			flag.Usage()
			os.Exit(1)
		}

		ignore, err := gitignore.Get(os.Args[2])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		if *getFile {
			// TODO handle appending to gitignore
			err := ioutil.WriteFile(".gitignore", ignore, 0644)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("%s\n", ignore)
		}

	case "list":
		listCmd.Parse(os.Args[2:])
		if len(listCmd.Args()) > 0 {
			flag.Usage()
			os.Exit(1)
		}
		ignores := gitignore.List()
		for _, file := range ignores {
			fmt.Printf("%s\n", file)
		}

	case "search":
		searchCmd.Parse(os.Args[2:])
		if len(searchCmd.Args()) > 1 {
			flag.Usage()
			os.Exit(1)
		}
		ignores, err := gitignore.Search(searchCmd.Args()[0])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		for _, file := range ignores {
			fmt.Printf("%s\n", file)
		}

	case "version":
		fmt.Printf("Version: %s\nCommit: %s\nBuilt at: %s\n", version, commit, date)

	default:
		flag.Usage()
		os.Exit(1)
	}
}
