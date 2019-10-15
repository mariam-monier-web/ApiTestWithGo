package fastHttpAdapter

import (
	"github.com/valyala/fasthttp" 
	// "github.com/qiangxue/fasthttp-routing"
	// "log"
	"encoding/json"
	
)



func JSONResponse(ctx *fasthttp.RequestCtx, code int, output interface{}) {
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	ctx.SetContentType("application/json")
	// Our response code
	ctx.SetStatusCode(code)

	ctx.Write(response)
}


// func Start (){

	
	
// 	router := routing.New()
	
// 	router.Get("/", ListUsers)
	
// 	log.Fatal(fasthttp.ListenAndServe(":8080", router.HandleRequest))
// }