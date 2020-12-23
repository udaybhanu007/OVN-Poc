package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserInfo struct {
	id   string
	name string
}

type UserList struct {
	users [5]UserInfo
}

var userList UserList

func main() {
	userList = generateUserList()
	fmt.Println(userList.users)

	//mux2 := http.NewServeMux()
	http.HandleFunc("/getUser/", getHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	r.ParseForm()
	fmt.Println("Url:", r.FormValue("id"))
	fmt.Println("Url Path:", r.URL.Path)
	path := strings.SplitAfter(r.URL.Path, "/")
	fmt.Println("Param:", path)
	id := path[len(path)-1]
	fmt.Println("ID:", id)
	var details UserInfo
	for _, value := range userList.users {
		fmt.Println("current row:", value)
		if value.id == id {
			details = value
			fmt.Println("details: ", details)
			break
		}
	}

	if details.id == "" {
		fmt.Println("User Not Found with ID : ", id)
		fmt.Fprintln(w, "User Not Found with ID : ", id)
	} else {
		fmt.Printf("user details => %s\n", details)
		fmt.Fprintln(w, "user found : ", details)
	}
}

func generateUserList() UserList {
	users := UserList{}
	for index := 0; index < 5; index++ {
		id := strconv.Itoa(100 + index)
		user := UserInfo{id, "Jayanthi" + strconv.Itoa(index)}
		users.users[index] = user
	}
	return users
}
