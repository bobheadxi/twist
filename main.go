package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	outDir       = flag.String("o", ".", "path to output directory")
	cfgPath      = flag.String("c", "", "path to Twist configuration")
	renderREADME = flag.Bool("readme", false, "toggle README rendering - requires configuration")
)

type pkg struct {
	Canonical string
	Source    string
}

//go:generate go run github.com/UnnoTed/fileb0x b0x.yml
func main() {
	flag.Parse()

	// if only one argument, we are dealing with a command
	if len(flag.Args()) == 1 {
		switch arg := flag.Arg(0); arg {
		case "help":
			showHelp()
			os.Exit(0)
		case "config":
			if *cfgPath == "" {
				*cfgPath = "./twist.yml"
			}
			b, err := yaml.Marshal(newConfig())
			if err != nil {
				panic(err)
			}
			if err := ioutil.WriteFile(*cfgPath, b, os.ModePerm); err != nil {
				panic(err)
			}
			fmt.Printf("config generated in '%s'\n", *cfgPath)
			os.Exit(0)
		default:
			println("insufficient arguments provided")
			os.Exit(1)
		}
	}

	// otherwise generate
	var cfg config
	if *cfgPath == "" {
		if len(flag.Args()) == 0 {
			println("insufficient arguments provided")
			os.Exit(1)
		}
		generate(flag.Arg(0), canonical{Path: flag.Arg(1), Subpackages: flag.Args()[2:]})
	} else {
		b, err := ioutil.ReadFile(*cfgPath)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(b, &cfg); err != nil {
			panic(err)
		}
		for s, c := range cfg.Packages {
			generate(s, c)
		}
	}

	if *renderREADME {
		if cfg.Packages == nil {
			panic("no configuration found")
		}
		generateREADME(&cfg)
	}
}

func showHelp() {
	println(`twist is a tool for generating static, serverless canonical imports for Go packages.

usage:

  twist -c twist.yml
  twist [source] [canonical] [subpackages...]

other commands:

  config         generate a twist configuration file
  help           show help text

flags:
`)
	flag.PrintDefaults()
	println("\nsee https://go.bobheadxi.dev/twist for more documentation.")
}
