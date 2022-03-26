package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/mattn/go-sqlite3"
)

type Item struct {
	ID int64	`json:"id"`
	Title  string  `json:"title"`
	Description string  `json:"desc"`
	Url  string `json:"url"`
	Tags []string `json:"tags"`
}
type Items map[int64]Item

type errorResponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

var DB string = "data.db"
var S *Storage = newStorage(DB)
var items Items = *S.loadAllItems()
// items slice to seed record item data.
// var items = []item{
// 	{
// 		Title: "南+",
// 		Description: "",
// 		Url: "https://south-plus.net",
// 		Tags: []string{"r18", "community"},
// 	},
// 	{
// 		Title: "r/selfhosted",
// 		Description: "",
// 		Url: "https://reddit.ccllssd.com/r/selfhosted",
// 		Tags: []string{"community", "selfhosted"},
// 	},
// 	{
// 		Title: "Rip@Lip (水原優) クレーム性処理女子社員2  ~謝罪出張~ [中国翻訳] [DL版]",
// 		Description: "",
// 		Url: "https://nhentai.net/g/395281/",
// 		Tags: []string{"r18:manga", "toread"},
// 	},
// }

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	
	router.GET("/items", getItems)
	router.GET("/items/:keywords", getItemByTitle)
	router.POST("/items", postItems)

	router.Run("0.0.0.0:8080")
}

// getItems responds with the list of all items as JSON.
func getItems(c *gin.Context) {
	items = *S.loadAllItems()
	c.IndentedJSON(http.StatusOK, items)
}

// postItems adds an item from JSON received in the request body.
func postItems(c *gin.Context) {
	var newItem Item

	// Call BindJSON to bind the received JSON to
	// newItem.
	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	var err error
	newItem.ID, err = S.saveItem(newItem)

	// return error code and message if S.saveItem returned non-nil error
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		log.Printf("database constraint error:\n\t%s", err.Error())

		// if a general error occured, return 400 and a general message
		responseCode := http.StatusBadRequest
		response := errorResponse{
			Error: err.Error(),
			Message: "error encounted when adding item",
		}
		// if adding an already existed item, return 409 and specified message
		if sqliteErr.Code == sqlite3.ErrNo(sqlite3.ErrConstraint) {
			response.Message = "posted item already exists in database"
			responseCode = http.StatusConflict
		}
		c.IndentedJSON(responseCode, response)
		return
	}
	// Add the new item to the slice.
	items[newItem.ID] = newItem
	c.IndentedJSON(http.StatusCreated, newItem)
}

// getItemByID locates the item whose ID value matches the id
// parameter sent by the client, then returns that item as a response.
func getItemByTitle(c *gin.Context) {
	keywords := c.Param("keywords")
	mathchedItems := make([]Item, 0)
	// Loop over the list of items, looking for
	// an item whose ID value matches the parameter.
	for _, it := range items {
		// check for keyword in item title
		if strings.Contains(strings.ToLower(it.Title), strings.ToLower(keywords)) {
			mathchedItems = append(mathchedItems, it)
			continue
		}
		// check for keyword in item url
		if strings.Contains(strings.ToLower(it.Url), strings.ToLower(keywords)) {
			mathchedItems = append(mathchedItems, it)
		}
		// check for keyword in item tag
		for _, t := range it.Tags {
			if strings.Contains(strings.ToLower(t), strings.ToLower(keywords)) {
				mathchedItems = append(mathchedItems, it)
				continue
			}
		}
	}

	c.IndentedJSON(http.StatusOK, mathchedItems)

	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
