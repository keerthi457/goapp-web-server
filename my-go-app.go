package main

import (
  "fmt"
  "net/http"
  "sync/atomic"
  "log"
  "os"
)

var calls int64
func HelloWorld(w http.ResponseWriter, r *http.Request) {
      for calls <10{
        count := atomic.AddInt64(&calls, 1)
        fmt.Fprintf(w, "You've called me %d times", count)
        file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
        if err != nil {
           log.Fatal(err)
        }

        log.SetOutput(file)

        log.Println("You've called me", count, "times")
        //fmt.Print(count)
           break;
        }
}
func main() {
  fmt.Printf("Started server at http://localhost%v.\n", 5000)
  http.HandleFunc("/", HelloWorld)
  http.ListenAndServe(":5000", nil)
}