all:
	go generate ./...
	go install

example-config: all
	twist config

example-redirects: all
	twist -c twist.example.yml -o x -readme
