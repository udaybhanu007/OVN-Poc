package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("mysql", "root:lavanya123@tcp(127.0.0.1:3306)/golang")

	fmt.Println("Connecting to mysql DB")

	if err != nil {

		fmt.Println("Not connected")
		panic(err.Error())

	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []Post

	result, err := db.Query("SELECT id, title from posts")
	if err != nil {
		fmt.Println("Error in selecting table")
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var post Post

		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			fmt.Println("Error in fetching data")
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmt, err := db.Prepare("INSERT INTO posts(id,title) VALUES(?,?)")

	if err != nil {
		fmt.Println("Error in inserting data")
		panic(err.Error())
	}

	defer stmt.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in reading data")
		panic(err.Error())
	}

	keyVal := make(map[string]string)

	json.Unmarshal(body, &keyVal)

	title := keyVal["title"]
	id := keyVal["id"]
	_, err = stmt.Exec(id, title)

	if err != nil {
		fmt.Println("Error in creating new post")
		panic(err.Error())
	}

	fmt.Fprintf(w, "New post was created")
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	result, err := db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
	if err != nil {
		fmt.Println("Error in fetching the post")
		panic(err.Error())
	}

	defer result.Close()

	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Title)
		if err != nil {
			fmt.Println("Error is traversing")
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	stmt, err := db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Error in update statement")
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in updating data")
		panic(err.Error())
	}
	keyVal := make(map[string]string)

	json.Unmarshal(body, &keyVal)

	newTitle := keyVal["title"]

	_, err = stmt.Exec(newTitle, params["id"])
	if err != nil {
		fmt.Println("Error in updating")
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		fmt.Println("Error in delete statement")
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])

	if err != nil {
		fmt.Println("Error in deleting data")
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}
