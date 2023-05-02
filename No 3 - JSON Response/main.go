package main

import (
	"datadiri"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func liatDatadiri(w http.ResponseWriter, r *http.Request) {
	nama := r.FormValue("Nama")

	datadiri := datadiri.ShowDataDiri(nama)

	str, err := json.Marshal(datadiri)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print(err)
	} else {
		io.WriteString(w, string(str))
	}
}

func semuaDatadiri(w http.ResponseWriter, r *http.Request) {

	semuadata := datadiri.ShowSemuaData()

	str, err := json.Marshal(semuadata)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print(err)
	} else {
		io.WriteString(w, string(str))
	}
}

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/nama", liatDatadiri)
	mux.HandleFunc("/semuadata", semuaDatadiri)

	http.ListenAndServe(":8000", mux)
}
