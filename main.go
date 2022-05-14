package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


// CRU API with functionality for addition and subtraction



// We will use gin package
// documentation https://github.com/gin-gonic/gin#installation

// when serializing this convert it to lower case json
type post struct {
	ID 		string `json:"id"`
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	Img 	string `json:"img"`
	Text 	string `json:"text"`
	Votes 	int `json:"votes"`
}

// data structures that have our posts
var posts = []post{
	{ID:"1", Title: "Python",Author: "Me",Text: "Python is a general purpose language", Votes: 3, Img: ""},
	{ID:"2", Title: "C",Author: "Me",Text: "C is a general purpose language", Votes: 4, Img: ""},
	{ID:"3", Title: "Go",Author: "Me",Text: "Go is a general purpose language", Votes: 12, Img: ""},
}


// Get all posts data function
func getPosts(c *gin.Context){
	c.IndentedJSON(http.StatusOK,posts)
}


// POST
// Create book
func createPost(c *gin.Context){
	var newPost post
	if err := c.BindJSON(&newPost); err != nil {
		return
	}
	posts = append(posts,newPost)
	c.IndentedJSON(http.StatusCreated, newPost)
}


// GET by id
func post_by_id(c *gin.Context){
	id := c.Param("id")
	post, err := get_post_by_id(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"This does not exist"})
		return 
	}

	c.IndentedJSON(http.StatusOK, post)
}

// search and return post function
func get_post_by_id(id string) (*post, error){
	// iterate all posts
	for i, p := range posts {
		if p.ID == id {
			// return book and no error
			return &posts[i], nil
		}
	}
	return nil, errors.New("Post not found")
}

// Vote for post
func vote_book(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Missing id"})
		return 
	}

	post, err := get_post_by_id(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"This does not exist"})
		return 
	}

	if post.Votes <= 0 {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"This language has zero votes"})
	}

	post.Votes +=1
	c.IndentedJSON(http.StatusOK, post)
}


// SAME as previous
// Downvote for post
func downvote_book(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Missing id"})
		return 
	}

	post, err := get_post_by_id(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"This does not exist"})
		return 
	}

	if post.Votes <= 0 {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"This language has zero votes"})
		return
	}

	post.Votes -=1
	c.IndentedJSON(http.StatusOK, post)
}





// run erver and do the routing
func main(){
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", post_by_id)
	router.POST("/posts", createPost)
	router.PATCH("/vote_post", vote_book)
	router.PATCH("/downvote_post", downvote_book)
	router.Run("localhost:3000")
}
