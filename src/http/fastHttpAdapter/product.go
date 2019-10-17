package fastHttpAdapter

import (	
	"model"
	"github.com/valyala/fasthttp" 
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
	// "fmt"
	"strings"
	"dataBase/gromAdapter"
	"dataBase/gromAdapter/dbModules"
	
	
)




func listProducts (ctx *fasthttp.RequestCtx){
	// type Products []dbModules.Product
	// products  := 
	 gromAdapter.ListProducts();
	response := model.CreateResponse {
		Message : "get product successfully ",
		// TotalCount: len(products),
		// Results: products,
	}

	JSONResponse(ctx, 200, response)


}


func createProduct(ctx *fasthttp.RequestCtx){
	var productObject dbModules.Product
	decoder := json.NewDecoder(strings.NewReader(string(ctx.PostBody())))
	decoder.Decode(&productObject)
	
	gromAdapter.CreateProduct(productObject);
	response := model.CreateResponse {
		Message : "create product successfully ",
	}
	JSONResponse(ctx, 200, response)

}


// func getProduct(ctx *fasthttp.RequestCtx){
// 	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET")
// 	id := fmt.Sprintf("%v", ctx.ProductValue("id")) 
// 	var product model.Product
// 	product = gromAdapter.GetProduct(id);
// 	response := model.GetResponse {
// 		Message : "get product successfully ",
// 		Data: product,
// 	}

// 	JSONResponse(ctx, 200, response)



// }

// func deleteProduct(ctx *fasthttp.RequestCtx){
// 	deleteConditions := model.Product{
// 		Email: fmt.Sprintf("%v", ctx.ProductValue("id")) ,
// 	}
	
// 	affectedRows:=  gromAdapter.DeleteProduct(deleteConditions);
// 	response := model.DeleteResponse {
// 		Message : "deleted product successfully ",
// 		AffectedRows: affectedRows,
// 	}
// 	JSONResponse(ctx, 200, response)
	

// }

// func updateProduct(ctx *fasthttp.RequestCtx){
// 	var updatedObject model.Product
// 	decoder := json.NewDecoder(strings.NewReader(string(ctx.PostBody())))
// 	 decoder.Decode(&updatedObject)
	
// 	var updateConditions model.Product 
// 	updateConditions = model.Product{
// 		Email:  fmt.Sprintf("%v", ctx.ProductValue("id")) ,
// 	}

// 	affectedRows:=  gromAdapter.UpdateProduct(updateConditions, updatedObject );
// 	response := model.UpdateResponse {
// 		Message : "updated product successfully" ,
// 		AffectedRows: affectedRows,
// 	}

// 	JSONResponse(ctx, 200, response)

	

// }








func HandleProductRequests (myRouter *fasthttprouter.Router){
	
	
	myRouter.GET("/products", listProducts)
	myRouter.POST("/products", createProduct)
	// myRouter.GET("/products/:id", getProduct)
	// myRouter.DELETE("/products/:id", deleteProduct)
	// myRouter.PUT("/products/:id", updateProduct)

}
