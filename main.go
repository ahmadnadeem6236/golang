package main

import "fmt"

func main() {

	var studyItem1 = "Watch golang tutorial"
	var studyItem2 = "Study about microservices"

	var workItem1 = "Create golang Todo App"
	var workItem2 = "Create microservices app for demonstartion."

	var taskItems = []string{studyItem1, studyItem2, workItem1, workItem2}

	fmt.Println("### GoLang TodoList App ###")

	taskItems = addTask(taskItems, "Go for run")
	taskItems = addTask(taskItems, "Practice go lang")

	getTask(taskItems)

	
}

func getTask(taskItems []string)  {
	fmt.Println("List of my Todos")
	fmt.Println()
	
	for index, task := range taskItems{
		// fmt.Println(index + 1 , ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
	
}


func addTask(taskItems []string,task string) []string {
	var updatedItem = append(taskItems, task)
	return updatedItem
	
}