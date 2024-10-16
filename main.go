package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"Study Go", "Go to the gym", "Eat something healthy"}

// func printTasks(taskItems []string) {
// 	fmt.Println("List of my Todos:")
// 	for index, task := range taskItems {
// 		fmt.Printf("%d. %s\n", index+1, task)
// 	}
// }

// func addTask(taskItems []string, newTask string) []string {
// 	taskItems = append(taskItems, newTask)
// 	return taskItems
// }

func helloUser(w http.ResponseWriter, r *http.Request) {
	var greeting = "Hello user!"
	fmt.Fprintf(w, greeting)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	for index, task := range taskItems {
		fmt.Fprintf(w, "%d. %s\n", index+1, task)
	}
}

func main() {
	fmt.Println("#### Welcome to out TODO App! ####")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)

	// var taskItems = []string{"Study Go", "Go to the gym", "Eat something healthy"}

	// taskItems = addTask(taskItems, "Win")
	// printTasks(taskItems)
}
