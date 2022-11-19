package main

import (
	"fmt"
	_ "latihan1/docs"
	"os"
	"time"

	"latihan1/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"xorm.io/xorm"
)

// @title           Hacktiv8 Swagger
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4444
// @BasePath  /api/v1
func main() {
	/*
		var firstName string = "Pitra"
		var lastName string
		lastName = "Wibisana"

		fmt.Printf("Nama lengkap : %s %s\n", firstName, lastName)

		name := new(string)
		fmt.Println(name)
		fmt.Println(*name)

		*name = "Hello"
		fmt.Println(name)
		fmt.Println(*name)

		students := make(map[string]string, 0)
		fmt.Println(students)
	*/

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed load .env : ", err.Error())
	}

	router := gin.Default()
	router.Use(GlobalMiddleware())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
	}

	todos := v1.Group("/todos")
	{
		todos.GET("", controllers.GetAll)
		todos.GET(":id", controllers.GetByID)
		todos.POST("", controllers.CreateTodo)
		todos.PUT(":id", controllers.UpdateByID)
		todos.DELETE(":id", controllers.DeleteByID)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":4444")
}

func GlobalMiddleware() gin.HandlerFunc {
	var dbConnection string = os.Getenv("DB_CONNECTION")
	var dbHost string = os.Getenv("DB_HOST")
	var dbPort string = os.Getenv("DB_PORT")
	var dbName string = os.Getenv("DB_DATABASE")
	var dbUser string = os.Getenv("DB_USERNAME")
	var dbPass string = os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	dbEngine, err := xorm.NewEngine(dbConnection, dsn)
	if err != nil {
		fmt.Println("dsn: ", dsn)
		fmt.Println("Failed Connect DB : ", err.Error())
	} else {
		fmt.Println("DB Connected")
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err == nil {
			dbEngine.SetTZDatabase(loc)
			dbEngine.SetTZLocation(loc)
		}
	}

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {

			}
		}()

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, app-token, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		c.Set("dbEngine", dbEngine)
	}
}
