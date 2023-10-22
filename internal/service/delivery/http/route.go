package http

import "github.com/avtara/carehub/pkg"

func (so *svObject) initRoute() {
	so.service.GET("/", so.handlerHelloWorld)

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

	category := so.service.Group("/categories")
	{
		category.GET("", so.HandlerGetAllCategory)
		category.GET("/:id", so.HandlerGetCategoryByID)

	}

	complain := so.service.Group("/complains")
	{
		complain.GET("", pkg.AdminMiddleware(so.handlerGetAllCompain))
		complain.GET("/:id", pkg.AdminMiddleware(so.handlerGetComplainByID))
	}
}
