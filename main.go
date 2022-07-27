package main

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ConnorRCAFInnovation/To-Do-app/task"
	"github.com/ConnorRCAFInnovation/To-Do-app/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
)

func helloWorld(c *fiber.Ctx){
	c.SendString("Hello world!")
}

func setupRoutes(app *fiber.App){
	app.Get("api/v1/task", task.GetTasks)
	app.Get("api/v1/task/:id", task.GetTask)
	app.Post("api/v1/task", task.NewTask)
	app.Delete("api/v1/task/:id", task.DeleteTask)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		panic("Failed to connect to the database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&task.Task{})
	fmt.Println("Database Migrated")	

}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")

}