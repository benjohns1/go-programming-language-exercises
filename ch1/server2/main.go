package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		var (
			cycles  = 5
			size    = 100
			nframes = 64
			delay   = 8
			phase   = 0.0
			res     = 0.001
			freq    = 3.0
			random  = false
		)
		for k, vals := range r.Form {
			var err error
			if len(vals) <= 0 {
				continue
			}
			v := vals[0]
			switch k {
			case "cycles":
				cycles, err = strconv.Atoi(v)
				if err != nil {
					log.Println(err)
				}
			case "size":
				size, err = strconv.Atoi(v)
				if err != nil {
					log.Println(err)
				}
			case "nframes":
				nframes, err = strconv.Atoi(v)
				if err != nil {
					log.Println(err)
				}
			case "delay":
				delay, err = strconv.Atoi(v)
				if err != nil {
					log.Println(err)
				}
			case "phase":
				phase, err = strconv.ParseFloat(v, 8)
				if err != nil {
					log.Println(err)
				}
			case "res":
				res, err = strconv.ParseFloat(v, 8)
				if err != nil {
					log.Println(err)
				}
			case "freq":
				freq, err = strconv.ParseFloat(v, 8)
				if err != nil {
					log.Println(err)
				}
			case "random":
				random, err = strconv.ParseBool(v)
				if err != nil {
					log.Println(err)
				}
			}
		}

		lissajous(w, cycles, size, nframes, delay, phase, res, freq, random)
	})
	port := 8080
	log.Printf("Listening on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
