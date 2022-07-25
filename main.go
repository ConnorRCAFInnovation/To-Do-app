package main

import(
	"github.com/gofiber/fiber"
	"github.com/ConnorRCAFInnovation/To-Do-app/task"
	
)

func helloWorld(c *fiber.Ctx){
	c.Send("Hello world!")
}

func setupRoutes(app *fiber.App){
	app.Get("api/v1/task", task.GetTasks)
	app.Get("api/v1/task/:id", task.GetTask)
	app.Post("api/v1/task", task.NewTask)
	app.Delete("api/v1/task/:id", task.DeleteTask)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(3000)

}