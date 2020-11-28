package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"net/http"
)

// createTodo add a new _todo
func CreateTodo(c *gin.Context) {
	// get the context and tries to conform it to the model
	var todo models.CreateTodoInput
	if err := c.ShouldBindJSON(&todo); err != nil || len(todo.Title) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// generate the data model to send to the database and get the object id
	input := models.Todos{Title: todo.Title, Completed: false}

	models.DB.Save(&input)

	c.JSON(http.StatusCreated, input)
}

// fetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var todos []models.Todos

	if err := models.DB.Order("id asc").Find(&todos).Error; err != nil || len(todos) <= 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todos)
}

// fetchSingleTodo fetch a single _todo
func FetchSingleTodo(c *gin.Context) {
	var todo models.Todos

	if fetch := models.DB.Where("id = ?", c.Param("id")).First(&todo); fetch.Error != nil || fetch.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todo)
}

// updateTodo update a _todo
func UpdateTodo(c *gin.Context) {

	// Get model if exist
	var todo models.Todos
	if fetch := models.DB.Where("id = ?", c.Param("id")).First(&todo); fetch.Error != nil || fetch.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Assign to the input the value of the searched _todo, then overwrite it with the input from the context
	// With that the update will not contains unexpected zero values
	// P.S. if you want pass the second level and so on, you need deep copy!
	input := todo
	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// gorm not set when we use zero-value (input.Completed = false), to avoid this I use map
	models.DB.Model(&todo).Updates(map[string]interface{}{"title": input.Title, "completed": input.Completed})

	c.JSON(http.StatusOK, todo)
}

// deleteTodo remove a _todo
func DeleteTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todos
	if fetch := models.DB.Where("id = ?", c.Param("id")).First(&todo); fetch.Error != nil || fetch.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	models.DB.Delete(&todo)

	c.JSON(http.StatusOK, todo)
}
