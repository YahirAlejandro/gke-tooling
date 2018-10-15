package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
MY_POD_NAME
MY_NODE_NAME
MY_POD_IP
*/

func main() {

	my_pod_name := os.Getenv("MY_POD_NAME")
	my_pod_ip := os.Getenv("MY_POD_IP")
	my_node_name := os.Getenv("MY_NODE_NAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(my_pod_name))
		fmt.Fprintf(w, string(my_pod_ip))
		fmt.Fprintf(w, string(my_node_name))

	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
