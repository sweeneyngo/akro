package server

import (
	"akro/acronym"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Run(port string) {

	// Generate Markov Chain if doesn't exist
	if !fileExists("model.json") {
		generateModel()
	}

	// Load Markov Chain
	mc, err := acronym.LoadFromFile("model.json")
	if err != nil {
		log.Fatalf("Error loading Markov Chain: %v", err)
	}

	router := http.NewServeMux()
	router.HandleFunc("/generate", handleRequest(mc))

	cors := applyCORS(router)

	fmt.Printf("Server started at http://localhost:%s\n", port)
	if err := http.ListenAndServe(port, cors); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func applyCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleRequest(mc *acronym.MarkovChain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		minLength, err := strconv.Atoi(query.Get("minLength"))
		if err != nil || minLength <= 0 {
			http.Error(w, "Invalid minLength", http.StatusBadRequest)
			return
		}

		maxLength, err := strconv.Atoi(query.Get("maxLength"))
		if err != nil || maxLength <= 0 {
			http.Error(w, "Invalid maxLength", http.StatusBadRequest)
			return
		}

		noiseLevel, err := strconv.Atoi(query.Get("noiseLevel"))
		if err != nil || noiseLevel < 0 {
			http.Error(w, "Invalid noiseLevel", http.StatusBadRequest)
			return
		}

		if minLength > maxLength {
			http.Error(w, "minLength cannot be greater than maxLength", http.StatusBadRequest)
			return
		}

		sentence := mc.GenerateSentence(minLength, maxLength)
		capLevel := rand.Intn(len(strings.Fields(sentence)))
		password, err := acronym.CreatePassword(sentence, capLevel, noiseLevel)

		if err != nil {
			http.Error(w, "Undefined error, can't generate password", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"sentence": sentence,
			"password": password,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func generateModel() {
	corpus, err := acronym.CreateCorpus("data")

	if err != nil {
		fmt.Println("error creating corpus:", err)
		return
	}

	mc := acronym.CreateMarkovChain()
	mc.AddCorpus(corpus)
	mc.SaveToFile("model.json")
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
