// API

/*
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	// I accessed to http://localhost:8080. But it was "Not Working".
	// Typed "ip a" to inspect the IP of Docker Machine
	// http://172.19.0.2:8080/ -> "write tcp 10.239.47.181:46561: write: network is unreachable"
}

*/

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
