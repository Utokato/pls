package offline

import (
	"embed"
)

//go:embed offline.zip
var Resource []byte

//go:embed dist
var Dist embed.FS
