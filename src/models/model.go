package models

// Employee represents an employee record in the system.
type Employee struct {
	ID        int32  `bson:"id" json:"id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	JobTitle  string `bson:"job_title" json:"job_title"`
	Email     string `bson:"email" json:"email"`
	Gender    string `bson:"gender" json:"gender"`
}
