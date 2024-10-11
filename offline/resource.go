package offline

import _ "embed"

//go:embed offline.zip
var Resource []byte

//go:embed dist.tar.gz
var Dist []byte
