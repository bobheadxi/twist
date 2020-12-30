package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"go.bobheadxi.dev/twist/internal"
)

func generate(source string, canon canonical) {
	b, err := internal.ReadFile("pkg.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("index.html").Parse(string(b))
	if err != nil {
		panic(err)
	}

	var (
		parts       = strings.Split(canon.Path, "/")
		packageName = parts[len(parts)-1]
	)
	for _, p := range append(canon.Subpackages, "") {
		target := filepath.Join(*outDir, packageName, p)
		os.MkdirAll(target, os.ModePerm)
		output := filepath.Join(target, "index.html")
		fmt.Printf("generating template in '%s' (for '%s' => '%s')\n", output, source, canon.Path)
		os.Remove(output)
		f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		if err := t.Execute(f, pkg{
			Source:    source,
			Canonical: canon.Path,
		}); err != nil {
			panic(err)
		}
		f.Sync()
		f.Close()
	}
}

func generateREADME(cfg *config) {
	// set up file
	target := filepath.Join(*outDir, "README.md")
	fmt.Printf("generating README in '%s'\n", target)
	os.Remove(target)
	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// sort packages
	keys := make([]string, len(cfg.Packages))
	var i int
	for s := range cfg.Packages {
		keys[i] = s
		i++
	}
	sort.Strings(keys)

	// render table
	table := tablewriter.NewWriter(f)
	table.SetHeader([]string{"Package", "Reference", "Source"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	for _, k := range keys {
		if cfg.Packages[k].NoREADME {
			continue
		}
		pkgPath := cfg.Packages[k].Path
		godocSource := pkgPath
		if cfg.GodocFromSource {
			godocSource = k
		}
		table.Append([]string{
			fmt.Sprintf("`%s`", pkgPath),
			fmt.Sprintf("[![godev](https://pkg.go.dev/badge/%s.svg)](https://pkg.go.dev/%s)", godocSource, godocSource),
			fmt.Sprintf("[%s](https://%s)", k, k),
		})
	}
	table.Render()
	f.WriteString("\n---\n")
	f.WriteString("\ngenerated using [twist](https://go.bobheadxi.dev/twist)\n")

	f.Sync()
	f.Close()
}
