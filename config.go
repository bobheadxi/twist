package main

type canonical struct {
	Path        string
	Subpackages []string
	NoREADME    bool
}

type config struct {
	GodocFromSource bool
	Packages        map[string]canonical
}

func newConfig() *config {
	return &config{Packages: map[string]canonical{
		"github.com/my/package": {
			Path:        "go.my.domain/package",
			Subpackages: []string{"subpackage1", "subpackage2"},
		},
	}}
}
