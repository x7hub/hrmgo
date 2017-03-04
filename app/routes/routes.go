// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tEmployee struct {}
var Employee tEmployee


func (_ tEmployee) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Employee.Index", args).Url
}

func (_ tEmployee) Add(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Employee.Add", args).Url
}

func (_ tEmployee) SaveAdd(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Employee.SaveAdd", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tAttendance struct {}
var Attendance tAttendance


func (_ tAttendance) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Attendance.Index", args).Url
}


type tDashboard struct {}
var Dashboard tDashboard


func (_ tDashboard) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Dashboard.Index", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


