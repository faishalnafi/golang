// controllers/student_controller.go

package controllers

import (
	"../databases"
)

func GetStudents() interface{} {
	// Di aplikasi nyata, fungsi ini akan mengambil data dari database
	// Untuk demonstrasi, ambil data dari fungsi dummy database
	return databases.GetStudents()
}
