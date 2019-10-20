package gromAdapter
 
import (
	
	_ "github.com/lib/pq"
	"dataBase/gromAdapter/dbModules"
	

	
)


  


func CreateProduct (productObject dbModules.Product)   {
	
	db :=DataBaseConnection()
	db.Create(&productObject)

	  defer db.Close()
}


func ListProducts ()[]dbModules.Product {
	db :=DataBaseConnection()
	
	var products []dbModules.Product
	 db.Debug().Find(&products)
	
	  defer db.Close()
	  return products ; 


}



func GetProduct (id int) dbModules.Product {
	db :=DataBaseConnection()
	var product dbModules.Product
	db.Debug().First(&product, id)
	defer db.Close()
	return product ; 


}


func UpdateProduct (id int, updateObject  dbModules.Product) int64 {

	db :=DataBaseConnection()
	var product dbModules.Product
	affectedRows := db.Debug().Model(&product).Where("ID = ?", id).Update(updateObject).RowsAffected
	defer db.Close()
	return affectedRows ; 


	
}


func DeleteProduct (id int) int64 {

	db :=DataBaseConnection()
	var product dbModules.Product
	affectedRows := db.Debug().Where("ID = ?", id).Delete(&product).RowsAffected
	defer db.Close()
	return affectedRows ; 


	
}