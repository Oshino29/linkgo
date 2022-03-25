package main

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

type item struct {
	Title  string  `json:"title"`
	Description string  `json:"desc"`
	Url  string `json:"url"`
	Tags []string `json:"tags"`
}

// items slice to seed record item data.
var items = []item{
	{
		Title: "南+",
		Description: "",
		Url: "https://south-plus.net",
		Tags: []string{"r18", "community"},
	},
	{
		Title: "r/selfhosted",
		Description: "",
		Url: "https://reddit.ccllssd.com/r/selfhosted",
		Tags: []string{"community", "selfhosted"},
	},
	{
		Title: "Rip@Lip (水原優) クレーム性処理女子社員2  ~謝罪出張~ [中国翻訳] [DL版]",
		Description: "",
		Url: "https://nhentai.net/g/395281/",
		Tags: []string{"r18:manga", "toread"},
	},
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.GET("/items/:keywords", getItemByTitle)
	router.POST("/items", postItems)

	router.Run("0.0.0.0:8080")
}

// getItems responds with the list of all items as JSON.
func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

// postItems adds an item from JSON received in the request body.
func postItems(c *gin.Context) {
	var newItem item

	// Call BindJSON to bind the received JSON to
	// newItem.
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	// Add the new item to the slice.
	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

// getItemByID locates the item whose ID value matches the id
// parameter sent by the client, then returns that item as a response.
func getItemByTitle(c *gin.Context) {
	keywords := c.Param("keywords")
	mathchedItems := make([]item, 0)
	// Loop over the list of items, looking for
	// an item whose ID value matches the parameter.
	for _, it := range items {
		if strings.Contains(strings.ToLower(it.Title), strings.ToLower(keywords)) {
			mathchedItems = append(mathchedItems, it)
			continue
		}
		for _, t := range it.Tags {
			if strings.Contains(strings.ToLower(t), strings.ToLower(keywords)) {
				mathchedItems = append(mathchedItems, it)
			}
		}
	}

	c.IndentedJSON(http.StatusOK, mathchedItems)

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
