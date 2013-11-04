package controllers

import (
	"fmt"
	"time"

	"github.com/robfig/revel"
)

// Action invoke info.
type ActionRTInfo struct {
	StartTime time.Time
	EndTime   time.Time
}

type XRuntimeableController struct {
	*revel.Controller
	RTInfo ActionRTInfo
}

func (c *XRuntimeableController) DoActionBegin() revel.Result {
	c.RTInfo.StartTime = time.Now()
	return nil
}

func (c *XRuntimeableController) DoActionEnd() revel.Result {
	c.RTInfo.EndTime = time.Now()
	costTime := c.RTInfo.EndTime.Sub(c.RTInfo.StartTime)
	c.Response.Out.Header().Add("X-Runtime", fmt.Sprintf("%v", costTime))
	return nil
}

func init() {
	revel.InterceptMethod((*XRuntimeableController).DoActionBegin, revel.BEFORE)
	revel.InterceptMethod((*XRuntimeableController).DoActionEnd, revel.AFTER)
}
