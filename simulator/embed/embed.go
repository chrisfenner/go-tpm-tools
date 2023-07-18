// package embed is a hack to allow go mod vendor to include the C dependencies of this package
// c.f. https://github.com/golang/go/issues/26366, https://seankhliao.com/blog/12021-05-09-go-mod-vendor-non-go-things/
package embed

import (
	_ "github.com/google/go-tpm-tools/simulator/ms-tpm-20-ref/Samples/Google"
	_ "github.com/google/go-tpm-tools/simulator/ms-tpm-20-ref/TPMCmd/tpm/include"
	_ "github.com/google/go-tpm-tools/simulator/ms-tpm-20-ref/TPMCmd/tpm/include/Ossl"
	_ "github.com/google/go-tpm-tools/simulator/ms-tpm-20-ref/TPMCmd/tpm/include/prototypes"
	_ "github.com/google/go-tpm-tools/simulator/ms-tpm-20-ref/TPMCmd/tpm/src/support"
)
