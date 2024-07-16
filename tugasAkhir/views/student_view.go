// views/student_view.go

package views

import (
	"fmt"

	"github.com/yourusername/tugasakhir/models"
)

func DisplayStudents(students []models.Student) {
	fmt.Println("Student List:")
	for _, student := range students {
		fmt.Printf("NISN: %s, Name: %s, Address: %s, Parent: %s, Class: %s\n", student.NISN, student.Name, student.Address, student.ParentName, student.Class)
	}
}
