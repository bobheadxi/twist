package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/bobheadxi/twist/internal"
)

var (
	outDir = flag.String("o", ".", "path to output directory")
)

type pkg struct {
	Canonical string
	Source    string
}

//go:generate go run github.com/UnnoTed/fileb0x b0x.yml
func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		println("arguments [source] [canonical] required")
		os.Exit(1)
	}

	b, err := internal.ReadFile("pkg.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("index.html").Parse(string(b))
	if err != nil {
		panic(err)
	}
	os.MkdirAll(*outDir, os.ModePerm)
	target := filepath.Join(*outDir, "index.html")
	fmt.Printf("generating template in '%s'\n", target)
	os.Remove(target)
	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if err := t.Execute(f, pkg{
		Source:    flag.Arg(0),
		Canonical: flag.Arg(1),
	}); err != nil {
		panic(err)
	}
	f.Sync()
	f.Close()
}
