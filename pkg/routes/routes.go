package routes

import (
	"net/http"

	"jilt.com/m/pkg/handlers"
)

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"newThread",
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	{
		"createThread",
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},

	{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},
	{
		"Category",
		"GET",
		"/category",
		handlers.Categories,
	},
	{
		"newCategory",
		"GET",
		"/category/new",
		handlers.NewCategory,
	},
	{
		"createCategory",
		"POST",
		"/category/create",
		handlers.CreateCategory,
	},
	{
		"goCategory",
		"GET",
		"/category/category",
		handlers.GoCategory,
	},
	{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
}
