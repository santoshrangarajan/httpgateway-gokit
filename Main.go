package main

//package main
import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

////https://github.com/tensor-programming/go-kit-tutorial
////https://sagikazarmark.hu/blog/getting-started-with-go-kit/

// Endpoints

func main() {

	//logger := log.NewLogfmtLogger(os.Stderr)

	var bksvc BookService

	bksvc = bookService{}
	//bksvc = loggingMiddleware(logger, bksvc)

	authornameHandler := httptransport.NewServer(
		makeAuthornameEndpoint(bksvc),
		decodeAuthornameRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(bksvc),
		decodeCountRequest,
		encodeResponse,
	)

	isavailableHandler := httptransport.NewServer(
		makeIsavailableEndpoint(bksvc),
		decodeIsAvailableRequest,
		encodeResponse,
	)

	//log.Println("starting....")
	http.Handle("/authorname", authornameHandler)
	http.Handle("/count", countHandler)
	http.Handle("/isavailable", isavailableHandler)

	//http.ListenAndServe(":8080", nil)

	log.Fatal(http.ListenAndServe(":8080", nil))

	//log.Println("Go-Kit POC")
}

/////// end point
/////// transport
/////// encapsulation of objects inside service
