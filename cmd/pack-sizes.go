package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"pack-sizes/internal/config"
	"pack-sizes/internal/endpoints"
	"pack-sizes/pkg/split"
)

func Start() {
	c := config.Get()
	r := mux.NewRouter()

	// by default use greedy algorithm which is mostly correct as dynamic tends to be slow
	// *always correct for default buckets set
	f := split.Greedy
	if c.Options.Depth != 0 {
		f = split.Dynamic
		fmt.Printf("using Dynamic algorithm with depth: %d\n", c.Options.Depth)
	} else {
		fmt.Println("using Greedy algorithm")
	}

	r.HandleFunc("/split", endpoints.MakeSplitEndpoint(c.Buckets, f, c.Options))

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", c.Host, c.Port), r)
	if err != nil {
		fmt.Printf("error starting server: %s", err)
	}
}
