package handlers

import (
	"net/http"

	"github.com/Chinzzii/Kubernetes-Load-Testing-and-Autoscaling/utils"
)

// DataHandler simulates fetching data.
func DataHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "Here is some data",
		"items":   []string{"item1", "item2", "item3"},
	}
	utils.JSONResponse(w, http.StatusOK, data)
}
