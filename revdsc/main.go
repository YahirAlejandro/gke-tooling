package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func validate_cidr(cidr string) bool {
	if cidr == "" {
		fmt.Printf("Env var is not set or was sent as an empty string!\n")
		log.Fatal("envvar is not set")
	}

	re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	if !re.MatchString(cidr) {
		fmt.Printf("is %v a valid ip?\n", cidr)
		log.Fatal("Invalid IP")
	}

	return true
}

func lookup_all_cidr_hosts(cidr string) {
	octet := strings.Split(cidr, ".")

	for i := 2; i <= 254; i++ {
		lookup_ip := fmt.Sprintf("%s.%s.%s.%v", octet[0], octet[1], octet[2], i)
		fmt.Printf("Will look up for: %s\n", lookup_ip)
		addr, err := net.LookupAddr(lookup_ip)
		fmt.Println(addr, err)
	}
}

func get_pods(w http.ResponseWriter, r *http.Request) {
	pod_cidr := os.Getenv("POD_CIDR")
	pod_cidr = strings.Trim(pod_cidr, " ")

	if validate_cidr(pod_cidr) {
		fmt.Fprintf(w, "Will use %s as searching address space...\n", pod_cidr)
	}

	lookup_all_cidr_hosts(pod_cidr)
}

func main() {

	landing_instructions := `Welcome,
	There should be 3 env vars set in the YAML definition of this pod.
	Each should be accessible through different URIs, for instance: "/pods" should return all hosts in the pods range.
	If correctly set, it will be the same as running "nmap -sL 1.2.3.0/24".
	This is using Go's standard libary net package to perform a reverse lookup on each IP from the CIDR.`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", landing_instructions)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/pods", get_pods)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
