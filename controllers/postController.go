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

	// Validation
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
			"status":  "E",
		})

		return
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{
		"message": "Created post successfully",
		"status":  "S",
	})
}

func PostList(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Validate if post is empty
	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No posts found",
			"status":  "E",
		})

		return
	}

	// Response data
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get ID from query param
	id := c.Param("id")

	// Get the posts
	var posts models.Post
	initializers.DB.First(&posts, id)

	// Validation
	if posts.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
			"status":  "E",
		})

		return
	}

	// Response data
	c.JSON(http.StatusOK, gin.H{
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

	// Validation
	if body.Title == "" && body.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please insert Title or Body",
			"status":  "E",
		})

		return
	}

	// Find the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Update Data
	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
			"status":  "E",
		})

		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Update post successfully",
		"status":  "S",
	})
}

func PostDelete(c *gin.Context) {
	// Get ID from query param
	id := c.Param("id")

	// Validation
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please insert ID",
			"status":  "E",
		})

		return
	}

	var checkPost models.Post
	initializers.DB.First(&checkPost, id)

	if checkPost.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
			"status":  "E",
		})

		return
	}

	// Delete the post
	var post models.Post
	initializers.DB.Unscoped().Delete(&post, id)

	// Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete post successfully",
		"status":  "S",
	})
}
