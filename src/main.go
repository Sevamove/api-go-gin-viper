/*
Simple respresentation of how to implement a CRUD functionality in API using Go.

Run command:

	$ go run main.go # start listening and serving

Open another terminal window and run this commands:

	$ curl localhost:8080/add-funder --header "Content-Type: application json" -d @single_funder.json --request "POST"
	$ curl localhost:8080/add-funders --header "Content-Type: application json" -d @numerous_funders.json --request "POST"

	$ curl localhost:8080/fetch-funders --request "GET"
	$ curl localhost:8080/fetch-funder?id=1 --request "GET"

	$ curl localhost:8080/update-funder?id=1 --request "PATCH"

	$ curl localhost:8080/delete-funder?id=1 --request "DELETE"
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"api/src/db"
	"api/src/helper"
	"api/src/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/*******************************************************************************
Declaring global variables.
*******************************************************************************/
var (
	host string
	port string
	url  string
	// Total IDs in db.
	funderIds helper.Counter
	// Initializing mapping between funder ID and Funder.
	idToFunder = make(map[uint64]model.Funder)
)

/*******************************************************************************
Run the programme.
*******************************************************************************/
func main() {

	getConfigData()

	// Working with gin.
	router := gin.Default()

	router.POST("/add-funder", addFunder)
	router.POST("/add-funders", addFunders)

	router.GET("/fetch-funders", fetchFunders)
	router.GET("/fetch-funder", fetchFunder)

	router.PATCH("/update-funder", updateFunder)
	router.DELETE("/delete-funder", deleteFunder)

	router.Run(url)
}

/*******************************************************************************
Use viper to get values from config.yaml
*******************************************************************************/
func getConfigData() {
	// Setting viper.
	helper.SetConfig("config" /*path*/, "config" /*name*/, "yaml" /*file extension*/)

	host = viper.GetString("server.host")
	port = viper.GetString("server.port")
	url = host + ":" + port
}

/*******************************************************************************
POST request in json type for 1 funder.
In order to add more than 1 funder, see addFunders().
*******************************************************************************/
func addFunder(c *gin.Context) {
	var newFunder model.Funder
	newFunder.Id = funderIds.Increment()
	newFunder.DateFunded = helper.GetCurrentTime()
	newFunder.Funded = true

	if err := c.BindJSON(&newFunder); err != nil {
		return
	}

	db.Funders = append(db.Funders, newFunder)

	idToFunder[newFunder.Id] = newFunder
	c.IndentedJSON(http.StatusCreated, idToFunder[newFunder.Id])
}

/*******************************************************************************
POST request in json type for list of funders.
*******************************************************************************/
func addFunders(c *gin.Context) {
	var newFunders []model.Funder

	if err := c.BindJSON(&newFunders); err != nil {
		return
	}

	for _, newFunder := range newFunders {
		newFunder.Id = funderIds.Increment()
		newFunder.DateFunded = helper.GetCurrentTime()
		newFunder.Funded = true

		db.Funders = append(db.Funders, newFunder)

		idToFunder[newFunder.Id] = newFunder
		c.IndentedJSON(http.StatusCreated, idToFunder[newFunder.Id])
	}

}

/*******************************************************************************
Return all funders in json format to the client.
For individual ID see FetchFunder()
*******************************************************************************/
func fetchFunders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.Funders)
}

/*******************************************************************************
Return a specified with ID funder in json format to the client.
*******************************************************************************/
func fetchFunder(c *gin.Context) {
	strId, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": "Missing id query param."})
		return
	}

	id, _ := strconv.ParseUint(strId, 10, 64)

	_, index, err := getFunder(id, c)

	if err != nil {
		errorMessage := fmt.Sprintf("ID %v: the funder is invalid.", id)
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, db.Funders[index])
}

/*******************************************************************************
Update a specified with ID funder with a new key/value pair value.
Increases funder.Amount value.
*******************************************************************************/
func updateFunder(c *gin.Context) {
	strId, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": "Missing id query param."})
		return
	}

	id, _ := strconv.ParseUint(strId, 10, 64)
	_, index, err := getFunder(id, c)

	if err != nil {
		errorMessage := fmt.Sprintf("ID %v: the funder is invalid.", id)
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	db.Funders[index].Amount += uint64(0.0716 * math.Pow(10, 18))
	idToFunder[id] = db.Funders[index]
	c.IndentedJSON(http.StatusOK, db.Funders[index])
}

/*******************************************************************************
Delete a specified with ID funder from the db.Funders
*******************************************************************************/
func deleteFunder(c *gin.Context) {
	strId, _ := c.GetQuery("id")
	id, _ := strconv.ParseUint(strId, 10, 64)

	_, index, err := getFunder(id, c)
	if err != nil {
		errorMessage := fmt.Sprintf("ID %v: the funder is invalid.", id)
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	db.Funders = append(db.Funders[:index], db.Funders[index+1:]...)

	message := fmt.Sprintf("ID %v: deleted.\n", id)
	c.IndentedJSON(http.StatusOK, message)
}

/*******************************************************************************
Return a specified with ID funder's data,
index in db.Funders, and a bool value if the funder valid.
*******************************************************************************/
func getFunder(id uint64, c *gin.Context) (model.Funder, uint64, error) {
	funderIndex, err := helper.IndexOf(id)

	isValid := db.Funders[funderIndex] == idToFunder[id]
	if isValid == true {
		return db.Funders[funderIndex], funderIndex, err
	} else {
		return db.Funders[funderIndex], 0, errors.New("-1")
	}
}
