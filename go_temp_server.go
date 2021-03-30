package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("vcgencmd", "measure_temp")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, out.String())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
