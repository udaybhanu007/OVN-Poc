package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	id   string
	name string
}

type Users struct {
	users [5]User
}

/*var users = map[string]string{
	"101": "Jayanthi",
	"102": "Jayanthi2",
	"103": "Jayanthi3",
}*/

var users Users

func main() {

	users = generateUserDetails()
	fmt.Println(users.users)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/getUser/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("Url:", r.FormValue("id"))
		//u, err := url.Parse(r.URL)
		// if err == nil {
		fmt.Println("Url Path:", r.URL.Path)
		path := strings.SplitAfter(r.URL.Path, "/")
		fmt.Println("Param:", path)
		id := path[len(path)-1]
		var details User
		for _, value := range users.users {
			if value.id == id {
				details = value
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

		//details, ok := users[id]
		/*if ok != true {
			fmt.Println("User Not Found with ID : ", id)
			fmt.Fprintln(w, "User Not Found with ID : ", id)
		} else {
			fmt.Printf("user details => %s\n", details)
			fmt.Fprintln(w, "user found : ", details)
		}*/
	})
	http.ListenAndServe(":8090", mux2)
}

func generateUserDetails() Users {
	users := Users{}
	for index := 0; index < 5; index++ {
		id := strconv.Itoa(100 + index)
		user := User{id, "Jayanthi" + strconv.Itoa(index)}
		users.users[index] = user
	}
	return users
}

/*func getUserDetails(id string) (string, error) {
	return users[id], errors.New("User Not Found with ID : " + id)
}*/
