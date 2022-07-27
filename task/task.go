package task

import(
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/ConnorRCAFInnovation/To-Do-app/database"
)

type Task struct {
	gorm.Model
	Title string `json:"title"`
	Deadline string `json:"deadline"`
	Completed bool `json:"completed"`
}

func GetTasks(c *fiber.Ctx) error{
	db := database.DBConn
	var tasks []Task
	db.Find(&tasks)
	c.JSON(tasks)
	return nil
}

func GetTask(c *fiber.Ctx) error{
	id := c.Params("id")
	db := database.DBConn
	var task Task
	db.Find(&task, id)
	c.JSON(task)
	return nil
}

func NewTask(c * fiber.Ctx) error{
	db := database.DBConn
	var task Task
	task.Title = "Do Laundry"
	task.Deadline = "July 28th"
	task.Completed = false
	db.Create(&task)
	c.JSON(task)
	return nil
}

func DeleteTask(c *fiber.Ctx) error{
	id := c.Params("id")
	db := database.DBConn
	
	var task Task
	db.First(&task, id)
	if task.Title == "" {
		return c.Status(500).SendString("no task found with given ID")
		
	}
	db.Delete(&task)
	c.SendString("Task Successfully deleted")
	return nil
}