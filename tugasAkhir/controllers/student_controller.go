// controllers/student_controller.go

package controllers

import "../models"

func GetStudents() interface{} {
	// In a real application, this function would fetch data from the database
	// For demonstration, return dummy data from models
	return models.Students
}
