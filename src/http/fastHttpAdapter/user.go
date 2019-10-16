package fastHttpAdapter

import (	
	"model"
	"dataBase"
	"github.com/valyala/fasthttp" 
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
	"fmt"
	"strings"
	
	
)


type Users []model.User

func listUsers (ctx *fasthttp.RequestCtx){
	
	users  :=dataBase.ListUsers();
	response := model.ListResponse {
		Message : "get user successfully ",
		TotalCount: len(users),
		Results: users,
	}

	JSONResponse(ctx, 200, response)


}


func registerUser(ctx *fasthttp.RequestCtx){
	var createUser model.User
	decoder := json.NewDecoder(strings.NewReader(string(ctx.PostBody())))
	decoder.Decode(&createUser)
	
	dataBase.RegisterUser(createUser);
	response := model.CreateResponse {
		Message : "create user successfully ",
	}
	JSONResponse(ctx, 200, response)

}


func getUser(ctx *fasthttp.RequestCtx){
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET")
	email := fmt.Sprintf("%v", ctx.UserValue("email")) 
	var user model.User
	user = dataBase.GetUser(email);
	response := model.GetResponse {
		Message : "get user successfully ",
		Data: user,
	}

	JSONResponse(ctx, 200, response)



}

func deleteUser(ctx *fasthttp.RequestCtx){
	deleteConditions := model.User{
		Email: fmt.Sprintf("%v", ctx.UserValue("email")) ,
	}
	
	affectedRows:=  dataBase.DeleteUser(deleteConditions);
	response := model.DeleteResponse {
		Message : "deleted user successfully ",
		AffectedRows: affectedRows,
	}
	JSONResponse(ctx, 200, response)
	

}

func updateUser(ctx *fasthttp.RequestCtx){
	var updatedObject model.User
	decoder := json.NewDecoder(strings.NewReader(string(ctx.PostBody())))
	 decoder.Decode(&updatedObject)
	
	var updateConditions model.User 
	updateConditions = model.User{
		Email:  fmt.Sprintf("%v", ctx.UserValue("email")) ,
	}

	affectedRows:=  dataBase.UpdateUser(updateConditions, updatedObject );
	response := model.UpdateResponse {
		Message : "updated user successfully" ,
		AffectedRows: affectedRows,
	}

	JSONResponse(ctx, 200, response)

	

}








func HandleRequests (myRouter *fasthttprouter.Router){
	
	
	myRouter.GET("/users", listUsers)
	myRouter.POST("/users", registerUser)
	myRouter.GET("/users/:email", getUser)
	myRouter.DELETE("/users/:email", deleteUser)
	myRouter.PUT("/users/:email", updateUser)

}
