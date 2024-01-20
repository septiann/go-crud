package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sept.dev/septian/go-crud/initializers"
	"sept.dev/septian/go-crud/models"
)

func PostCreate(c *gin.Context) {
	// Get data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	if body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Title is required",
			"status":  "E",
		})

		return
	}

	// Create data
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)

		return
	}

	// Return data
	c.JSON(200, gin.H{
		"message": "Created post successfully",
		"status":  "S",
	})
}

func PostList(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Response data
	c.JSON(200, gin.H{
		"data": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get ID from query param
	id := c.Param("id")

	// Get the posts
	var posts models.Post
	initializers.DB.First(&posts, id)

	// Response data
	c.JSON(200, gin.H{
		"data": posts,
	})
}

func PostUpdate(c *gin.Context) {
	// Get ID from query param
	id := c.Param("id")

	// Get data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Update Data
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Response
	c.JSON(200, gin.H{
		"message": "Update post successfully",
		"status":  "S",
	})
}

func PostDelete(c *gin.Context) {
	// Get ID from query param
	id := c.Param("id")

	// Delete the post
	var post models.Post
	initializers.DB.Delete(&post, id)

	// Response
	c.JSON(200, gin.H{
		"message": "Delete post successfully",
		"status":  "S",
	})
}
