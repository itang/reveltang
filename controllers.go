package reveltang

import (
	"fmt"
	"time"

	"github.com/robfig/revel"
)

type XRuntimeable struct {
	*revel.Controller

	StartTime time.Time
	CostTime  time.Duration
	EndTime   time.Time
}

func (c *XRuntimeable) DoBegin() revel.Result {
	c.StartTime = time.Now()
	return nil
}

func (c *XRuntimeable) DoEnd() revel.Result {
	c.EndTime = time.Now()
	c.CostTime = c.EndTime.Sub(c.StartTime)
	c.Response.Out.Header().Add("X-Runtime", fmt.Sprintf("%v", c.CostTime))
	return nil
}

func init() {
	revel.InterceptMethod((*XRuntimeable).DoBegin, revel.BEFORE)
	revel.InterceptMethod((*XRuntimeable).DoEnd, revel.AFTER)
}
