package models

// Employee represents an employee record in the system.
type Employee struct {
	ID        int32  `bson:"id" json:"id" validate:"omitempty,number"`
	FirstName string `bson:"first_name" json:"first_name" validate:"required,min=1,max=100"`
	LastName  string `bson:"last_name" json:"last_name" validate:"required,min=1,max=100"`
	JobTitle  string `bson:"job_title" json:"job_title" validate:"required,min=1,max=100"`
	Email     string `bson:"email" json:"email" validate:"required,email"`
	Gender    string `bson:"gender" json:"gender" validate:"required,oneof=Male Female Other"`
}
