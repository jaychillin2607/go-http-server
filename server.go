package server

import (
	"fmt"
	"io"
	"net/http"
)

func PlayerServer(w io.Writer, r *http.Request) {
	fmt.Fprint(w, 20)
}
