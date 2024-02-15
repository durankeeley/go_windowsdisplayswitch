package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/changeDisplayMode/", changeDisplayModeHandler) 
	log.Println("Server starting on port 8999...")
	if err := http.ListenAndServe(":8999", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func changeDisplayModeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	param := strings.TrimPrefix(r.URL.Path, "/changeDisplayMode/")
	if param == "" {
		http.Error(w, "Parameter is missing in the URL", http.StatusBadRequest)
		
		return
	}

	if !isValidParameter(param) {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("DisplaySwitch.exe", "/"+param)
	if err := cmd.Run(); err != nil {
		http.Error(w, "Failed to execute command", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Display mode changed successfully"))
}

func isValidParameter(param string) bool {
	allowedParams := []string{"external", "internal", "clone", "extend"}
	for _, p := range allowedParams {
		if param == p {
			return true
		}
	}
	return false
}
