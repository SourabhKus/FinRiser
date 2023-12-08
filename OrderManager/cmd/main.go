package main

import (
	"fmt"
	"net"
	//"net/http"
)

func HostHttpServer() {

}

func ConnectToTicker(add string){
	conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    // Send data to the server
    // ...

    // Read and process data from the server
    // ...
	
	/*var wg sync.WaitGroup
	wg.Add(1) /// Added for the HttpServer
	
	go func () {
		defer wg.Done()
	
		//wrappedMux := middleware.NewMiddleware(mux)
		server := http.Server{
			Addr:    ":" + srvSettings.ServicePort,
			//Handler: wrappedMux,
		}
		
		if err = server.ListenAndServe(); err != nil {
			fmt.Println("func", "main", "error", err)
		}
	}*/
}


