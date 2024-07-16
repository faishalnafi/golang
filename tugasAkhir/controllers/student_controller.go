// controllers/student_controller.go

package controllers

import (
	"github.com/yourusername/tugasakhir/models"
	"github.com/yourusername/tugasakhir/views"
)

func GetStudents() {
	// In a real application, you would fetch data from the database.
	students := models.Students
	views.DisplayStudents(students)
}
