package main

import (
	"dd-server/api"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Handle all request.
//func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
//	// TODO should be configurable?
//	w.WriteHeader(http.StatusOK)
//	if status == http.StatusNotFound {
//		record := Record{
//			Url:    r.URL.String(),
//			Header: r.Header,
//			Method: r.Method,
//		}
//		if body, err := decodeRequestBody(r); err != nil {
//			record.Body = fmt.Sprintf("Parse error: %s", err.Error())
//		} else {
//			record.Body = body
//		}
//
//		log.Println("--------------------------------------------------------------------------------")
//		b, _ := json.Marshal(record)
//		log.Println(string(b))
//	}
//}

// Default handler, do nothing.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello world!`))
}

func main() {

	port := flag.String("port", "8080", "sets service port number")
	log_param := flag.String("log", "", "save all request to a log file")

	flag.Parse()

	// save logs to file if set.
	if *log_param != "" {
		f, err := os.OpenFile(*log_param, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			panic(fmt.Sprintf("error opening file: %s", err.Error()))
		}
		defer f.Close()

		log.SetOutput(f)

	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/v1/series/", api.SeriesHandler)
	http.HandleFunc("/intake/", api.IntakeHandler)

	err := http.ListenAndServe("0.0.0.0:"+*port, nil)
	if err != nil {
		fmt.Printf("I think something here could work, but not this")
	}
}
