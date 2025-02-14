package todolist

import (
	"fmt"
	"log"
	"net/http"
)

func Start(port int) {
	handler := http.NewServeMux()

	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
