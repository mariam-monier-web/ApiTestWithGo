package fastHttpAdapter

import (
	"github.com/valyala/fasthttp" 
	// "log"
	"encoding/json"
	
	
)



func JSONResponse(ctx *fasthttp.RequestCtx, code int, output interface{}) {
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	ctx.SetContentType("application/json")
	// Our response code
	ctx.SetStatusCode(code)
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Write(response)
}


//  func RequestHandler(ctx *fasthttp.RequestCtx) {
	

// 		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
// 		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET")
// 		// ctx.Response.Header.Set("Access-Control-Max-Age", "0")
// 		// ctx.Response.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
// 		// ctx.Response.Header.Set("Pragma", "no-cache")
// 		// ctx.Response.Header.Set("Expires", "Wed, 21 Oct 2015 07:28:00 GMT")

		
// }



	


