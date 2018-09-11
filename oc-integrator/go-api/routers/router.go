package routers

import (
	"openshift_exp/oc-integrator/go-api/common"
	"openshift_exp/oc-integrator/go-api/controllers"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	// if gin.Mode() == gin.ReleaseMode {
	// 	app.Use(sessions.Sessions("mysession", models.RedisStore))
	// } else {
	// 	app.Use(sessions.Sessions("mysession", models.MemStore))
	// }

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "oc-int zone",
		Key:           []byte(common.JWT_SECRET),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: controllers.LoginJWT,
		Authorizator:  controllers.UserAuthor,
		Unauthorized:  controllers.UnAuthor,
		LoginResponse: controllers.LoginResponse,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "OC-Int",
		TimeFunc:      time.Now,
	}

	conf := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	conf.AllowAllOrigins = true
	app.Use(cors.New(conf))

	// app.Use(static.Serve("/static", common.GetBFS("assets/static")))

	// r := multitemplate.New()
	// bytes, err := common.Asset("assets/index.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// t, err := template.New("index").Parse(string(bytes))
	// fmt.Println(t, err)

	// r.Add("index", t)
	// app.HTMLRender = r
	// app.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index", gin.H{})
	// })
	app.POST("/login", authMiddleware.LoginHandler)

	userAPI := app.Group("/user")
	userAPI.Use(authMiddleware.MiddlewareFunc())
	{
		userAPI.GET("/info", controllers.GetUserInfo)
		userAPI.POST("/logout", controllers.LogoutHandler)
		userAPI.GET("/list", controllers.GetUsers)
	}

	// adminAPI := app.Group("/api")
	// adminAPI.Use(authMiddleware.MiddlewareFunc())
	// {
	// 	adminAPI.GET("/site", controllers.GetSite)
	// }

	return app
}
