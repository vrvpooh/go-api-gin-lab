package models

type Student struct {
	// id and name must not be empty
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Major string `json:"major"`
	// gpa must be between 0.00 and 4.00
	GPA float64 `json:"gpa" binding:"min=0,max=4"`
}
