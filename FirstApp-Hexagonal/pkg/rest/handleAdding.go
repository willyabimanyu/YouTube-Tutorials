package rest

import (
	"net/http"

	"github.com/willyabimanyu/FirstApp-Hexagonal/pkg/adding"
)

func handleAdding(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
