package models

type Item struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

type Employee struct {
    ID        int32 `bson:"id"`
    FirstName string `bson:"first_name"`
    LastName  string `bson:"last_name"`
    JobTitle  string `bson:"job_title"`
    Email     string `bson:"email"`
    Gender    string `bson:"gender"`
}