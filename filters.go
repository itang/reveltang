package reveltang

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/revel"
)

func DoFilterChain(c *revel.Controller, fc []revel.Filter) {
	if len(fc) > 0 {
		fc[0](c, fc[1:])
	}
}

func XRuntimeHeaderFilter(c *revel.Controller, fc []revel.Filter) {
	if !isStaticResourceRequest(c) {
		start := time.Now()
		DoFilterChain(c, fc)
		end := time.Now()
		cost := end.Sub(start)
		c.Response.Out.Header().Add("X-Runtime", fmt.Sprintf("%v", cost))
	} else {
		DoFilterChain(c, fc)
	}
}

// -----------------------------------------------------------------------

var statics = []string{".js", ".css", ".gif", ".jpg", ".png", ".map"}

func isStaticResourceRequest(c *revel.Controller) bool {
	for _, v := range statics {
		if strings.HasSuffix(c.Request.RequestURI, v) {
			return true
		}
	}
	return false
}
