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
		"profile",
		"GET",
		"/profile",
		handlers.Profile,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"closeSession",
		"POST",
		"/session/close",
		handlers.AdminCloseASession,
	},
	{
		"editProfile",
		"GET",
		"/edit_profile",
		handlers.EditProfile,
	},
	{
		"admin",
		"GET",
		"/admin",
		handlers.Admin,
	},
	{
		"adminUpdate",
		"POST",
		"/admin/update",
		handlers.AdminUpdate,
	},
	{
		"adminUpdate2",
		"POST",
		"/admin/update2",
		handlers.AdminUpdate2,
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
	{
		"disccord",
		"GET",
		"/discord",
		handlers.Discord,
	},
	{
		"authDiscord",
		"GET",
		"/auth/discord",
		handlers.AuthDiscord,
	},
	{
		"google",
		"GET",
		"/google",
		handlers.Google,
	},
	{
		"authGoogle",
		"GET",
		"/auth/google",
		handlers.AuthGoogle,
	},
	{
		"facebook",
		"GET",
		"/facebook",
		handlers.Facebook,
	},
	{
		"authFacebook",
		"GET",
		"/auth/facebook",
		handlers.AuthFacebook,
	},
	// {
	// 	"twitter",
	// 	"GET",
	// 	"/twitter",
	// 	handlers.Twitter,
	// },
	// {
	// 	"authTwitter",
	// 	"GET",
	// 	"/auth/twitter",
	// 	handlers.AuthTwitter,
	// },
	// {
	// 	"github",
	// 	"GET",
	// 	"/github",
	// 	handlers.Github,
	// },
	// {
	// 	"authGithub",
	// 	"GET",
	// 	"/auth/github",
	// 	handlers.AuthGithub,
	// },
}
