// package embed is a hack to allow go mod vendor to include the C dependencies of this package
// c.f. https://github.com/golang/go/issues/26366, https://seankhliao.com/blog/12021-05-09-go-mod-vendor-non-go-things/
package embed

import "embed"

// go:embed *.h
var _ embed.FS
