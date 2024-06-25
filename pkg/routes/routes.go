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
		"/thread/comment",
		handlers.CommentThread,
	},
	{
		"Category",
		"GET",
		"/categories",
		handlers.Categories,
	},
	{
		"newCategory",
		"GET",
		"/categories/new",
		handlers.NewCategory,
	},
	{
		"createCategory",
		"POST",
		"/categories/create",
		handlers.CreateCategory,
	},
	{
		"deleteCategory",
		"GET",
		"/categories/delete",
		handlers.DeleteCategory,
	},
	{
		"goCategory",
		"GET",
		"/categories/category",
		handlers.GoCategory,
	},
	{
		"Topics",
		"GET",
		"/topics",
		handlers.Topics,
	},
	{
		"newTopic",
		"GET",
		"/topics/new",
		handlers.NewTopic,
	},
	{
		"createTopic",
		"POST",
		"/topics/create",
		handlers.CreateTopic,
	},
	{
		"deleteTopic",
		"GET",
		"/topics/delete",
		handlers.DeleteTopic,
	},
	{
		"goTopic",
		"GET",
		"/topics/topic",
		handlers.GoTopic,
	},
	{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
}
