package dataBase
import (
	"model"
	"database/sql"
	"fmt"
	"strings"
	"reflect"
	// "log"
	
)

const (
	tableName     = "users"
  )


  func ListUsers () []model.User {
	
	db := DataBaseConnection()
	sqlStatement :="SELECT * FROM "+tableName
	var user model.User
	var users []model.User
	
	rows, err := db.Query(sqlStatement)
	
	if err != nil {
		// handle this error better than this
		panic(err)
	  }
	  defer rows.Close()
	
	for rows.Next() {
		err = rows.Scan(&user.FirstName, &user.LastName , &user.Email, &user.Phone)
		if err != nil {
		  // handle this error
		  panic(err)
		}
		users = append(users, user)
		
	  }
	  // get any error encountered during iteration
	  err = rows.Err()
	  if err != nil {
		panic(err)
	  }
	  defer db.Close()
	 return users
}




func GetUser (email string ) model.User  {
	
		db := DataBaseConnection()
		sqlStatement :=`SELECT * FROM public.users  WHERE email=$1;`
		var user model.User
		row := db.QueryRow(sqlStatement, email)
		err :=row.Scan(&user.Email, &user.FirstName, &user.LastName, &user.Phone)
		switch err {
					case sql.ErrNoRows:
						fmt.Println("No rows were returned!")
					case nil:
						return user
					default:
						panic(err)
					}
		 
					defer db.Close()
					return user
}



func RegisterUser (UserObject model.User)   {
	
	db := DataBaseConnection()

	 createColumsValues := reflect.ValueOf(&UserObject).Elem()
	 var columSqlStatement, valueSqlStatement  string
	 for i := 0; i < createColumsValues.NumField(); i++ {
		 
		 varName := createColumsValues.Type().Field(i).Name
		 varValue := fmt.Sprintf("%v", createColumsValues.Field(i).Interface())
		 
		 if (len(varValue) != 0 ) {
			 if(len(columSqlStatement)== 0 ){
				 columSqlStatement = "INSERT INTO "+ tableName +" (" + strings.ToLower(varName)
			 }else {
				 columSqlStatement = columSqlStatement +", "+ strings.ToLower(varName)
			 }
			 
		 }
		 
	 }
	 
	columSqlStatement = columSqlStatement+ ")"


	for i := 0; i < createColumsValues.NumField(); i++ {
		varValue := fmt.Sprintf("%v", createColumsValues.Field(i).Interface())
		if (len(varValue) != 0 ) {
			if(len(valueSqlStatement)== 0 ){
				valueSqlStatement = "VALUES ('" + varValue + "'"
			}else {
				valueSqlStatement = valueSqlStatement +", '"+ varValue + "'"
			}
			
		}
		
	}
	valueSqlStatement = valueSqlStatement +")"
	sqlStatement := columSqlStatement + valueSqlStatement
	_, err := db.Exec(sqlStatement)
		if err != nil {
		panic(err)
		}
	defer db.Close()
}


func UpdateUser (updateConditions model.User, updateObject model.User) int64 {
	db := DataBaseConnection()
	createColumsValues := reflect.ValueOf(&updateObject).Elem()
	createCondition := reflect.ValueOf(&updateConditions).Elem()
	var columSqlStatement , conditionSqlStatement string
    for i := 0; i < createColumsValues.NumField(); i++ {
		
		varName := createColumsValues.Type().Field(i).Name
		varValue := fmt.Sprintf("%v", createColumsValues.Field(i).Interface())
		
		if (len(varValue) != 0 ) {
			if(len(columSqlStatement)== 0 ){
				columSqlStatement = "UPDATE "+ tableName +" SET " + strings.ToLower(varName) + " = '" + varValue + "' "
			}else {
				columSqlStatement = columSqlStatement +", "+ strings.ToLower(varName) + "= '" + varValue + "'"
			}
			
		}
		
	}

	
	for i := 0; i < createCondition.NumField(); i++ {
		
		varName := createCondition.Type().Field(i).Name
		varValue := fmt.Sprintf("%v", createCondition.Field(i).Interface())
		if(len(varValue) != 0){
			if( len(conditionSqlStatement)== 0 ){
				conditionSqlStatement = " WHERE "+ strings.ToLower(varName) + "= '" + varValue + "'"
			}else{
				conditionSqlStatement = conditionSqlStatement +" and "+ strings.ToLower(varName) + "= '" + varValue + "'"
			}
			
		}
				
			
		}
		

	sqlStatement  :=  columSqlStatement + conditionSqlStatement 
		
	res, err := db.Exec(sqlStatement )
	if err != nil {
	panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
	panic(err)
	}
	defer db.Close()
	return count

}

func DeleteUser (deleteConditions model.User) int64{
	db := DataBaseConnection()
	createCondition := reflect.ValueOf(&deleteConditions).Elem()
	var   conditionSqlStatement string
    

	
	for i := 0; i < createCondition.NumField(); i++ {
		
		varName := createCondition.Type().Field(i).Name
		varValue := fmt.Sprintf("%v", createCondition.Field(i).Interface())
		if(len(varValue) != 0){
			if( len(conditionSqlStatement)== 0 ){
				conditionSqlStatement = strings.ToLower(varName) + "= '" + varValue + "'"
			}else{
				conditionSqlStatement = conditionSqlStatement +" and "+ strings.ToLower(varName) + "= '" + varValue + "'"
			}
			
		}
				
			
		}
		

	sqlStatement  :=  "DELETE FROM public.users WHERE  "+ conditionSqlStatement 
	res, err := db.Exec(sqlStatement )
	if err != nil {
	panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
	panic(err)
	}
	defer db.Close()
	return count


}