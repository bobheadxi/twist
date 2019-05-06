package main

type config struct {
	Packages map[string]string
}

func newConfig() *config { return &config{Packages: make(map[string]string)} }
