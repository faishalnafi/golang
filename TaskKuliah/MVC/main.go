// main.go

package main

import (
	"./controllers"
	"./views"
)

func main() {
	// Contoh memanggil controller untuk mendapatkan dan menampilkan data siswa
	students := controllers.GetStudents()
	views.DisplayStudents(students)
}
