// models/student.go

package models

type Student struct {
	NISN       string
	Name       string
	Address    string
	ParentName string
	Class      string
}

// Dummy data for demonstration
var Students = []Student{
	{NISN: "12345678", Name: "John Doe", Address: "123 Main St", ParentName: "Jane Doe", Class: "XII-A"},
	{NISN: "87654321", Name: "Jane Smith", Address: "456 Oak Ave", ParentName: "John Smith", Class: "XII-B"},
	// Add more dummy data as needed
}
