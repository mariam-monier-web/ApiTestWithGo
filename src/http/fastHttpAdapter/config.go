package fastHttpAdapter

import (
	"github.com/valyala/fasthttp" 
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
	"log"
	
	
)



func JSONResponse(ctx *fasthttp.RequestCtx, code int, output interface{}) {
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	ctx.SetContentType("application/json")
	// Our response code
	ctx.SetStatusCode(code)
	ctx.Write(response)
}


func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		// ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		// ctx.Response.Header.Set("Access-Control-Allow-Headers", "authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "HEAD,GET,POST,PUT,DELETE,OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")

		next(ctx)
	}
}









func StartServer (){

	myRouter := fasthttprouter.New()
	HandleUserRequests(myRouter)
	HandleProductRequests(myRouter)
	
	if err := fasthttp.ListenAndServe(":8181", CORS(myRouter.Handler)); err != nil {
        log.Fatalf("Error in ListenAndServe: %s", err)
    }

}





	


