package controllers

import (
    "github.com/revel/revel"
)

type Employee struct {
    *revel.Controller
}

// GET /employee/index
func (c Employee) Index() revel.Result {
    return c.Render()
}

// GET /employee/add
func (c Employee) Add() revel.Result {
    return c.Render()
}

// POST /employee/add
func (c Employee) SaveAdd() revel.Result {
    name := c.Params.Form.Get("name")
    regularizeDate := c.Params.Form.Get("regularize-date")
    birthday := c.Params.Form.Get("birthday")
    revel.INFO.Printf("Employee.SaveAdd: %s %s %s", name, regularizeDate, birthday)
    return c.Redirect(Employee.Index)
}