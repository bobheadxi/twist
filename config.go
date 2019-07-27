package main

type canonical struct {
	Path        string
	Subpackages []string
}

type config struct {
	Packages map[string]canonical
}

func newConfig() *config {
	return &config{Packages: map[string]canonical{
		"github.com/my/package": {
			Path:        "go.my.domain/package",
			Subpackages: []string{"subpackage1", "subpackage2"},
		},
	}}
}
