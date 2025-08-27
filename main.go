package main

import (
	authcontroller "go-restapi-gin/controllers/authcontrollers"
	usercontroller "go-restapi-gin/controllers/usercontrollers"
	"go-restapi-gin/middleware"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    models.ConnectDatabase()

    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })
    
     

    // auth
    r.POST("/api/register", authcontroller.Register)
    r.POST("/api/login", authcontroller.Login)
    r.POST("/api/logout", middleware.AuthRequired(), authcontroller.Logout)

    // endpoint public / orders
    r.GET("/api/orders", usercontroller.Index)
    r.GET("/api/orders/:id", usercontroller.Show)
    r.POST("/api/orders", usercontroller.Create)
    r.PUT("/api/orders/:id", usercontroller.Update)
    r.DELETE("/api/orders", usercontroller.Delete)

    // contoh protected route
    protected := r.Group("/api")
    protected.Use(middleware.AuthRequired())

    // Role
    protected.GET("/roles", usercontroller.IndexRole)
    protected.GET("/roles/:id", usercontroller.ShowRole)
    protected.POST("/roles", usercontroller.CreateRole)
    protected.PUT("/roles/:id", usercontroller.UpdateRole)
    protected.DELETE("/roles/:id", usercontroller.DeleteRole)

    // Menu
    protected.GET("/menus", usercontroller.IndexMenu)
    protected.GET("/menus/:id", usercontroller.ShowMenu)
    protected.POST("/menus", usercontroller.CreateMenu)
    protected.PUT("/menus/:id", usercontroller.UpdateMenu)
    protected.DELETE("/menus/:id", usercontroller.DeleteMenu)

    // Akses
    protected.GET("/akses", usercontroller.IndexAkses)
    protected.GET("/akses/:id", usercontroller.ShowAkses)
    protected.POST("/akses", usercontroller.CreateAkses)
    protected.PUT("/akses/:id", usercontroller.UpdateAkses)
    protected.DELETE("/akses/:id", usercontroller.DeleteAkses)

    //users
    protected.GET("/profile", func(c *gin.Context) {
        u, _ := c.Get("user")
        c.JSON(200, gin.H{"user": u})
    })

    // Users API
    protected.GET("/users", usercontroller.Users)
    // protected.GET("/users/:id", usercontroller.ShowUser)
    // protected.POST("/users", usercontroller.CreateUser)
    // protected.PUT("/users/:id", usercontroller.UpdateUser)
    // protected.DELETE("/users/:id", usercontroller.DeleteUser)

    r.Run()
}
