package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/Chinzzii/Kubernetes-Load-Testing-and-Autoscaling/utils"
)

// ComputeHandler simulates a computational task.
func ComputeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Simulate heavy computation
	result := 0.0
	for i := 1; i <= 1000000; i++ {
		result += math.Sqrt(float64(i))
	}

	elapsed := time.Since(start)
	response := map[string]interface{}{
		"result":  result,
		"elapsed": elapsed.String(),
	}
	utils.JSONResponse(w, http.StatusOK, response)
}