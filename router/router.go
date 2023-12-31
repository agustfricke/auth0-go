package router

import (
	"encoding/gob"

	"github.com/agustfricke/auth0-go/auth"
	"github.com/agustfricke/auth0-go/handlers"
	"github.com/agustfricke/auth0-go/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New(auth *auth.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.Home)
	router.GET("/login", handlers.Login(auth))
	router.GET("/callback", handlers.Callback(auth))
	router.GET("/user", middleware.IsAuthenticated, handlers.User)
	router.GET("/logout", handlers.Logout)

	return router
}
