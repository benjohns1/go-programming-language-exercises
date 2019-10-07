package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/benjohns1/go-programming-language-exercises/ch3/server/surface"
)

func main() {
	http.HandleFunc("/", surfaceHandler)
	http.ListenAndServe(":8080", nil)
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		errorResponse(w, err, 404)
		return
	}
	var (
		width, height = 600, 320
		cells         = 100
		xyrange       = 30.0
	)
	for k, vals := range r.Form {
		var err error
		if len(vals) <= 0 {
			continue
		}
		v := vals[0]
		switch k {
		case "width":
			width, err = strconv.Atoi(v)
			if err != nil {
				errorResponse(w, err, http.StatusBadRequest)
				return
			}
		case "height":
			height, err = strconv.Atoi(v)
			if err != nil {
				errorResponse(w, err, http.StatusBadRequest)
				return
			}
		case "cells":
			cells, err = strconv.Atoi(v)
			if err != nil {
				errorResponse(w, err, http.StatusBadRequest)
				return
			}
		case "xyrange":
			xyrange, err = strconv.ParseFloat(v, 64)
			if err != nil {
				errorResponse(w, err, http.StatusBadRequest)
				return
			}
		}
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.SvgSurface(w, width, height, cells, xyrange)
}

func errorResponse(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	log.Println(err)
	fmt.Fprintf(w, err.Error())
}
