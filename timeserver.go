package main

import (
    "fmt"
    "net/http"
	"time"
	"strconv"
	"os"
    "bufio"
)

func weekday(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, time.Now().Weekday().String())
	fmt.Fprintf(w, "\n")
}

func tux(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Fprintf(w, "\n")
}

func index(w http.ResponseWriter, req *http.Request){
	file, err := os.Open("index.html")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "ERR\n")
		return;
    }
    defer file.Close()
	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	
	scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)    // use scanwords
    for scanner.Scan() {
		fmt.Fprintf(w, scanner.Text())
		fmt.Fprintf(w, " ")
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
		fmt.Fprintf(w, "ERR\n")
		return;
    }
}

func headers(w http.ResponseWriter, req *http.Request) {
	req.Header.Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Timeserver - DEBUG - Headers</title>
	<style>table, th, td {border:1px solid black;}</style>
</head>
<body><p><a href="/">Back</a></p><h1>Header</h1><table><tr><th>name</th><th>value</th></tr>`)

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "<tr><th>%v</th><th>%v</th></tr>", name, h)
        }
    }

	fmt.Fprintf(w, "</table></body></head>")
}

func main() {

    http.HandleFunc("/time/weekday", weekday)
	http.HandleFunc("/time/unix", tux)
    http.HandleFunc("/http/echo/headers", headers)
	http.HandleFunc("/", index)

    http.ListenAndServe(":8090", nil)
	fmt.Print("Timeserver is running\n")
}
