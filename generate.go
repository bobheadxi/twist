package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/bobheadxi/twist/internal"
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
