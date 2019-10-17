package gromAdapter
 
import (
	
	_ "github.com/lib/pq"
	"fmt"
	"dataBase/gromAdapter/dbModules"
	

	
)


  


func CreateProduct (productObject dbModules.Product)   {
	
	db :=DataBaseConnection()
	db.Create(&productObject)

	  defer db.Close()
}


func ListProducts () {
	db :=DataBaseConnection()
	var products []dbModules.Product
	result := db.Find(&products).Scan(&products)
fmt.Println("==================",*result )
	  defer db.Close()
	//    return result;


}