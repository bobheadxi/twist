# twist

Twist generates canonical imports for your Go packages. Since it does not require
a running server (ie in existing tools like [uber-go/sally](https://github.com/uber-go/sally)
and [rsc/go-import-redirector](https://github.com/rsc/go-import-redirector)),
Twist is particularly useful in conjunction with [GitHub Pages](https://pages.github.com/).

A canonical import path allows you to make your package import a little fancier
with a custom domain, for example:

```diff
- import "github.com/bobheadxi/zapx"
+ import "go.bobheadxi.dev/zapx"
```

## usage

```sh
go get -u go.bobheadxi.dev/twist
#          [        source         ] [     canonical     ]
twist -o x github.com/bobheadxi/zapx go.bobheadxi.dev/zapx
```

Using the example in this repo:

```sh
twist -c twist.example.yml -o x -readme
```

To set up your own configuration:

```sh
twist config
```

You'll want to commit the generated files to the GitHub Page repository of the
domain you want to use for your custom import path. For example, I used Twist to
set up my [`go.bobheadxi.dev/...`](https://go.bobheadxi.dev/) import paths using
GitHub Pages. The repository is [here](https://github.com/bobheadxi/go), and has
the following layout:

```
github.com/bobheadxi/go
|-- CNAME (go.bobheadxi.dev)
|-- README.md
|-- gobenchdata
|    +-- index.html (go.bobheadxi.dev/gobenchdata)
|-- package1
|    +-- index.html (go.bobheadxi.dev/package1)
+-- package2
     +-- index.html (go.bobheadxi.dev/package2)
```

GitHub Pages serves up the contents of the repository, allowing packages to
be served with my custom domain:

```
go get go.bobheadxi.dev/gobenchdata
```

In your Go package, you'll need to update all import paths to use the new name.

When using [Go Modules](https://github.com/golang/go/wiki/Modules), you'll also
need to update the `module` directive in your `go.mod`:

```go
module go.bobheadxi.dev/twist

go 1.12

require ( ... )
```
