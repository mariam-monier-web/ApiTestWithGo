package model

// import (
// 	"database/sql"
// 	 "fmt"
// 	 "reflect"
// 	 )
	




type User struct {

	FirstName string
	LastName string 
	Email string 
	Phone string 

}

type UpdateResponse struct {

	Message string
	AffectedRows int64
	

}
type DeleteResponse struct {

	Message string
	AffectedRows int64

}
type CreateResponse struct {

	Message string 

}
type GetResponse struct {
	Message string 
	Data interface{}
}
type ListResponse struct {

	Message string
	TotalCount int
	Results []User  

}



