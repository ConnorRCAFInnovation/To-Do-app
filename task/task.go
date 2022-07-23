package task

import(
	"github.com/gofiber/fiber"

)

func GetTasks(c *fiber.Ctx){
	c.Send("All task")
}

func GetTask(c *fiber.Ctx){
	c.Send("A single task")
}

func NewTask(c * fiber.Ctx){
	c.Send("A new task")
}

func DeleteTask(c *fiber.Ctx){
	c.Send("Deletes a task")
}