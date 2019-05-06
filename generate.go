package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"go.bobheadxi.dev/twist/internal"
	"github.com/olekukonko/tablewriter"
)

func generate(source, canonical string) {
	b, err := internal.ReadFile("pkg.html")
	if err != nil {
		panic(err)
	}
	t, err := template.New("index.html").Parse(string(b))
	if err != nil {
		panic(err)
	}

	parts := strings.Split(canonical, "/")
	packageName := parts[len(parts)-1]
	target := filepath.Join(*outDir, packageName)
	os.MkdirAll(target, os.ModePerm)
	target = filepath.Join(target, "index.html")
	fmt.Printf("generating template in '%s' (for '%s' => '%s')\n", target, source, canonical)
	os.Remove(target)
	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if err := t.Execute(f, pkg{
		Source:    source,
		Canonical: canonical,
	}); err != nil {
		panic(err)
	}
	f.Sync()
	f.Close()
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

	// render table
	table := tablewriter.NewWriter(f)
	table.SetHeader([]string{"Package", "Source"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	for s, c := range cfg.Packages {
		table.Append([]string{c, fmt.Sprintf("[%s](https://%s)", s, s)})
	}
	table.Render()
	f.WriteString("\n---\n")
	f.WriteString("\ngenerated using [twist](https://go.bobheadxi.dev/twist)\n")

	f.Sync()
	f.Close()
}
