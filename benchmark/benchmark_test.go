package server

import (
	"akro/acronym"
	"testing"
)

func BenchmarkGenerateSentence(b *testing.B) {
	// Load your Markov Chain model
	mc, err := acronym.LoadFromFile("../model.json")
	if err != nil {
		b.Fatalf("Error loading Markov Chain: %v", err)
	}

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		mc.GenerateSentence(5, 15) // Adjust min and max length as needed
	}
}

func BenchmarkCreatePassword(b *testing.B) {
	// Load your Markov Chain model
	mc, err := acronym.LoadFromFile("../model.json")
	if err != nil {
		b.Fatalf("Error loading Markov Chain: %v", err)
	}

	// Generate a sentence for the password creation benchmark
	sentence := mc.GenerateSentence(5, 15) // Generate a sample sentence
	capLevel := 1                          // Adjust as needed
	noiseLevel := 0                        // Adjust as needed

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		_, err := acronym.CreatePassword(sentence, capLevel, noiseLevel)
		if err != nil {
			b.Fatalf("Error generating password: %v", err)
		}
	}
}
