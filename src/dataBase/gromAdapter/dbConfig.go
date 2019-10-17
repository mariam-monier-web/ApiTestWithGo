package gromAdapter

 import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"enviroment"
	// "dataBase/gromAdapter/dbModules"
	"fmt"
 )


 


  
  
  func DataBaseConnection() *gorm.DB {


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


	db, err := gorm.Open("postgres",psqlInfo)
	if err != nil {
	  panic("failed to connect database")
	}
	
  
	// // Migrate the schema
	// db.AutoMigrate(&Product{})
  
	// Create
	// db.Create(&dbModules.DbUser{
	// 	firstName :"mariam",
	// 	lastName:"monier",
	// 	email:"mmonier@iotblue.net",
	// 	phone:"0142345452478754",
		
	
	// })
  
	// // Read
	// var product Product
	// db.First(&product, 1) // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212
  
	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)
  
	// // Delete - delete product
	// db.Delete(&product)
	return db
  }