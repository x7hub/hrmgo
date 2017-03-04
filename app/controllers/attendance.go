package controllers

import "github.com/revel/revel"

type Attendance struct {
    *revel.Controller
}

func (c Attendance) Index() revel.Result {
    return c.Render()
}
