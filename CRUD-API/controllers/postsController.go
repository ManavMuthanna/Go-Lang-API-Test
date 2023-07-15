package controllers

import (
	"net/http"

	"github.com/ManavMuthanna/CRUD-API/initializers"
	"github.com/ManavMuthanna/CRUD-API/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	//Get the posts
	var posts models.Post
	initializers.DB.First(&posts, id)
	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"post": posts,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//Find the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//respond
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get the id
	id := c.Param("id")

	//Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.Status(200)
}
