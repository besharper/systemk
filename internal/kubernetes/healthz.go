package kubernetes

import (
	"fmt"
	"net/http"
)

func InstallHealthzHandler(mux *http.ServeMux) {
	path := "/healhtz"
	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s check passed\n", path)
	})
}
