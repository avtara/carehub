package http

import "github.com/avtara/carehub/pkg"

func (so *svObject) initRoute() {
	so.service.GET("/", pkg.JwtMiddleware(so.handlerHelloWorld))

	auth := so.service.Group("/auth")
	{
		auth.GET("/google", so.handlerAuthenticationGoogle)
		auth.GET("/google/callback", so.handlerCallbackGoogle)
		auth.POST("/login", so.handlerAuthenticationLogin)
	}

	static := so.service.Group("/assets")
	{
		static.Static("/", "assets")
	}

	user := so.service.Group("/users")
	{
		user.PUT("/edit", pkg.JwtMiddleware(so.handlerEditProfileUser))
	}

}
