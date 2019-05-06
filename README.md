# twist

generate canonical imports for your Go packages. useful in conjunction with
[GitHub Pages](https://pages.github.com/), since it does not require a running
server (ie in existing tools like [uber-go/sally](https://github.com/uber-go/sally)
and [rsc/go-import-redirector](https://github.com/rsc/go-import-redirector)).

## usage

```
go get github.com/bobheadxi/twist
twist github.com/bobheadxi/zapx bobheadxi.dev/zapx
```
