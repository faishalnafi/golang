// databases/database.go

package databases

import "github.com/yourusername/tugasakhir/models"

// GetDummyStudents mengembalikan data siswa untuk keperluan demonstrasi
func GetDummyStudents() []models.Student {
	// Data dummy untuk demonstrasi
	students := []models.Student{
		{NISN: "12345678", Name: "John Doe", Address: "123 Jalan Utama", ParentName: "Jane Doe", Class: "XII-A"},
		{NISN: "87654321", Name: "Jane Smith", Address: "456 Jalan Oak", ParentName: "John Smith", Class: "XII-B"},
		// Tambahkan data dummy lainnya jika diperlukan
	}
	return students
}

// Simulated database operations
func GetStudents() interface{} {
	// Di aplikasi nyata, fungsi ini akan berinteraksi dengan database yang sebenarnya
	// Untuk demonstrasi, kembalikan data dummy
	return GetDummyStudents()
}
