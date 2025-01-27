package templates

var wrapper = `
package {{.Package}}

import (
	_ "embed"
)

`
