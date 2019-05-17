package main

import (
	"log"
	"net/http"
	"os/exec"
)

// pull and restart
func Relaunch() {
	cmd := exec.Command("sh", "./relaunch.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>deploy server</h1>"))
	Relaunch()
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":5000", nil)
}
