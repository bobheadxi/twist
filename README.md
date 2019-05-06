# twist

Generate canonical imports for your Go packages. Since it does not require a running
server (ie in existing tools like [uber-go/sally](https://github.com/uber-go/sally)
and [rsc/go-import-redirector](https://github.com/rsc/go-import-redirector)),
this is particularly useful in conjunction with [GitHub Pages](https://pages.github.com/).

A canonical import path allows you to make your package import a little fancier
with a custom domain, for example:

```diff
- import "github.com/bobheadxi/zapx"
+ import "bobheadxi.dev/zapx"
```

## usage

```sh
go get github.com/bobheadxi/twist
#             [ source ]          [ canonical ]
twist github.com/bobheadxi/zapx bobheadxi.dev/zapx
```
