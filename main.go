package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type CommandRequest struct {
	Command string `json:"command"`
}

func main() {
	http.HandleFunc("/api/cmd", handleCmd)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCmd(w http.ResponseWriter, r *http.Request) {

	var request CommandRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.Command == "" {
		http.Error(w, "Command not found", http.StatusBadRequest)
		return
	}

	cmd := exec.Command(request.Command)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing command: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("Output is", string(output))

	w.Write(output)
}
