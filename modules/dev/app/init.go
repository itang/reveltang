package app

import (
	"github.com/revel/revel"
)

func init() {
	revel.INFO.Printf("%s module: SourceChanged reload browser\n", "dev")
}
