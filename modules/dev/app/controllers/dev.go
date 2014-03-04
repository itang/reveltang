package controllers

import (
	"time"

	. "github.com/itang/reveltang/controllers"
	"github.com/revel/revel"
)

var serverSince = time.Now().UnixNano() / 1000000.0

type Dev struct {
	*revel.Controller
	XRuntimeableController
}

func (c Dev) SourceChanged(clientSince int64) revel.Result {
	if serverSince > clientSince {
		return c.RenderText("true")
	}
	return c.RenderText("false")
}
