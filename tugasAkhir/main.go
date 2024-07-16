// main.go

package main

import (
	"./controllers"
	"./views"
)

func main() {
	// Example of calling the controller to get and display student data
	students := controllers.GetStudents()
	views.DisplayStudents(students)
}
