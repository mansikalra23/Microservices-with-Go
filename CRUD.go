// CRUD operations with Gin Gonic

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User - Structure for user information
type User struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

var Users []User

func FindUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Users})
}

func FindUser(c *gin.Context) {
	key := c.Param("id")

	for _, user := range Users {
		if user.Id == key {
			c.JSON(http.StatusOK, gin.H{"data": user})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "Record not found"})
}

func CreateUser(c *gin.Context) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := User{Id: input.Id, Name: input.Name}
	Users = append(Users, user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var input User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	key := c.Param("id")
	for index, user := range Users {
		if user.Id == key {
			Users = append(Users[:index], Users[index+1:]...)
			var user User
			user = User{Id: key, Name: input.Name}
			Users = append(Users, user)
			c.JSON(http.StatusOK, gin.H{"data": user})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "Record not found."})
}

func DeleteUser(c *gin.Context) {
	key := c.Param("id")
	for index, user := range Users {
		if user.Id == key {
			Users = append(Users[:index], Users[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"data": "Record deleted."})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": "Record not found."})
}

func handleRequests() {
	r := gin.Default()
	r.GET("/users", FindUsers)
	r.GET("/users/:id", FindUser)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)
	r.Run(":8000")
}

func main() {
	var n int
	fmt.Print("Number of records : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var id, name string
		fmt.Printf("Enter roll no of %d user : ", i+1)
		fmt.Scan(&id)
		fmt.Printf("Enter name of %d user : ", i+1)
		fmt.Scan(&name)
		s := User{Id: id, Name: name}
		Users = append(Users, s)
	}
	fmt.Printf("Records entered : %s\n", Users)

	handleRequests()
}
