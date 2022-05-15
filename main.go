package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
)

// CRU API with functionality for addition and subtraction

// We will use gin package
// documentation https://github.com/gin-gonic/gin#installation

// when serializing this convert it to lower case json
type post struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Img    string `json:"img"`
	Text   string `json:"text"`
	Votes  int    `json:"votes"`
}

// data structures that have our posts
var posts = []post{
	{ID: "1", Title: "Python", Author: "Me", Text: "Python is a high-level, interpreted, general-purpose programming language. Its design philosophy emphasizes code readability with the use of significant indentation.", Votes: 5, Img: ""},
	{ID: "2", Title: "C/C++", Author: "Me", Text: "C is a structural programming language, and it does not support classes and objects, while C++ is an object-oriented programming language that supports the concept of classes and objects.", Votes: 4, Img: ""},
	{ID: "3", Title: "Go", Author: "Me", Text: "Go is a statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing and CSP-style concurrency.", Votes: 5, Img: ""},
	{ID: "4", Title: "Java", Author: "Me", Text: "Java is a high-level, class-based, object-oriented programming language that is designed to have as few implementation dependencies as possible.", Votes: 3, Img: ""},
	{ID: "5", Title: "C#", Author: "Me", Text: "C# is a general-purpose, multi-paradigm programming language. C# encompasses static typing, strong typing, lexically scoped, imperative, declarative, functional, generic, object-oriented, and component-oriented programming disciplines.", Votes: 2, Img: ""},
	{ID: "6", Title: "Javascript", Author: "Me", Text: "JavaScript, often abbreviated JS, is a programming language that is one of the core technologies of the World Wide Web, alongside HTML and CSS. Over 97% of websites use JavaScript on the client side for web page behavior", Votes: 4, Img: ""},
	{ID: "7", Title: "Rust", Author: "Me", Text: "Rust is a multi-paradigm, general-purpose programming language designed for performance and safety, especially safe concurrency. It is syntactically similar to C++, but can guarantee memory safety ", Votes: 0, Img: ""},
	{ID: "8", Title: "PHP", Author: "Me", Text: "PHP is a general-purpose scripting language geared toward web development. It was originally created by Danish-Canadian programmer Rasmus Lerdorf in 1994.", Votes: 1, Img: ""},
	{ID: "9", Title: "Swift", Author: "Me", Text: "Swift is a general-purpose, multi-paradigm, compiled programming language developed by Apple Inc. and the open-source community.", Votes: 0, Img: ""},
}

// Get all posts data function
func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

// POST
// Create book
func createPost(c *gin.Context) {
	var newPost post
	if err := c.BindJSON(&newPost); err != nil {
		return
	}
	posts = append(posts, newPost)
	c.IndentedJSON(http.StatusCreated, newPost)
}

// GET by id
func post_by_id(c *gin.Context) {
	id := c.Param("id")
	post, err := get_post_by_id(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This does not exist"})
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}

// search and return post function
func get_post_by_id(id string) (*post, error) {
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
func vote_book(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing id"})
		return
	}

	post, err := get_post_by_id(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This does not exist"})
		return
	}

	if post.Votes <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This language has zero votes"})
	}

	post.Votes += 1
	c.IndentedJSON(http.StatusOK, post)
}

// SAME as previous
// Downvote for post
func downvote_book(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing id"})
		return
	}

	post, err := get_post_by_id(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "This does not exist"})
		return
	}

	if post.Votes <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "This language has zero votes"})
		return
	}

	post.Votes -= 1
	c.IndentedJSON(http.StatusOK, post)
}

// run erver and do the routing
func main() {
	router := gin.Default()
	router.Use(cors.AllowAll())
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", post_by_id)
	router.POST("/posts", createPost)
	router.PATCH("/vote_post", vote_book)
	router.PATCH("/downvote_post", downvote_book)
	router.Run("localhost:3000")
}
