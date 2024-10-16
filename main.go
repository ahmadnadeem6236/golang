package main

import "fmt"

func main() {

	var studyItem1 = "Watch golang tutorial"
	var studyItem2 = "Study about microservices"

	var workItem1 = "Create golang Todo App"
	var workItem2 = "Create microservices app for demonstartion."

	var taskItems = []string{studyItem1, studyItem2, workItem1, workItem2}

	fmt.Println("### GoLang TodoList App ###")

	fmt.Println("List of my Todos")
	fmt.Println()
	
	for index, task := range taskItems{
		// fmt.Println(index + 1 , ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
	
	
	//fmt.Println("Tasks", taskItems)


}