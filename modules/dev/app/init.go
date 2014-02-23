package app

import (
	"github.com/robfig/revel"
)

func init() {
	revel.INFO.Printf("%s module: SourceChanged reload browser\n", "dev")
}
