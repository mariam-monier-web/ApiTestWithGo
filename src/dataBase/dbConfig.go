package dataBase
import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"enviroment"
)


	
		
	
	  

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "admin"
// 	password = "Iotblue55$"
// 	dbname   = "goDataBase"
//   )
  
var db *sql.DB

func DataBaseConnection()(*sql.DB){


	v1, err := enviroment.ReadConfig()
	  if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	  }
	
	  port := v1.GetInt("port")
	  hostname := v1.GetString("hostname")
	  username := v1.GetString("username")
	  password := v1.GetString("password")
	  dbname := v1.GetString("dbname")
	



	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
	hostname, port, username, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
		panic(err)
		}

		


		err = db.Ping()
		if err != nil {
		panic(err)
		
		}


	fmt.Println("Successfully connected!")
	return db



}