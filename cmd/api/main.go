package main

import (
	"log"
	"net/http"
	"time"

	thttp "github.com/arizaguaca/table/internal/http"
	"github.com/arizaguaca/table/internal/repository"
	"github.com/arizaguaca/table/internal/usecase"
)

func main() {
	// 1. Setup Repository
	tableRepo := repository.NewMemoryTableRepository()

	// 2. Setup Usecase
	timeoutContext := time.Duration(10) * time.Second
	tableUsecase := usecase.NewTableUsecase(tableRepo, timeoutContext)

	// 3. Setup Handler
	tableHandler := thttp.NewTableHandler(tableUsecase)

	// 4. Setup Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/tables", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			tableHandler.Create(w, r)
		case http.MethodGet:
			tableHandler.Fetch(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 5. Start Server
	serverAddr := ":8080"
	log.Printf("Server starting on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
