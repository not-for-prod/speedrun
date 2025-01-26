
package sql

import (
	_ "embed"
)

//go:embed create.sql
var Create string

//go:embed get.sql
var Get string

//go:embed update.sql
var Update string

//go:embed delete.sql
var Delete string
