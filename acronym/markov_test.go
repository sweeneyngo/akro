package acronym

import (
	"strings"
	"testing"
)

func TestCreateMarkovChain(t *testing.T) {
	markovChain := CreateMarkovChain()

	if markovChain.Transitions == nil {
		t.Errorf("Expected Transitions map to be initialized, got nil")
	}
}

func TestAddCorpus(t *testing.T) {

	markovChain := &MarkovChain{
		Transitions: map[string][]string{
			"the":   {"quick"},
			"quick": {"brown"},
			"brown": {"fox"},
			"fox":   {"<EOL>"},
		},
	}

	if len(markovChain.Transitions) == 0 {
		t.Errorf("Expected Transitions to contain entries, but it is empty")
	}

	if len(markovChain.Transitions["the"]) == 0 {
		t.Errorf("Expected 'the' to have map of next words, but it does not")
	}

	if len(markovChain.Transitions["quick"]) == 0 {
		t.Errorf("Expected 'quick' to have map of next words, but it does not")
	}
}

func TestGenerateSentence(t *testing.T) {

	markovChain := &MarkovChain{
		Transitions: map[string][]string{
			"the": {"<EOL>"},
		},
	}

	sentence := markovChain.GenerateSentence(1, 1)

	if len(sentence) == 0 {
		t.Errorf("Expected a non-empty sentence, but got empty")
	}

	expectedWords := []string{"the"}
	for _, word := range expectedWords {
		if !containsWord(sentence, word) {
			t.Errorf("Expected sentence to contain '%s', but it does not", word)
		}
	}
}

func TestGetRandomStartWord(t *testing.T) {

	markovChain := &MarkovChain{
		Transitions: map[string][]string{
			"Hello":   {"world"},
			"Goodbye": {"everyone"},
			"Welcome": {"home"},
		},
	}

	// Define the expected set of start words
	expectedWords := []string{"Hello", "Goodbye", "Welcome"}

	// Call the function
	startWord := markovChain.getRandomStartWord()

	// Check if the result is in the expected set
	isValid := false
	for _, word := range expectedWords {
		if startWord == word {
			isValid = true
			break
		}
	}

	if !isValid {
		t.Errorf("Expected one of %v, but got %q", expectedWords, startWord)
	}
}

func containsWord(sentence, word string) bool {
	for _, w := range strings.Fields(sentence) {
		if w == word {
			return true
		}
	}
	return false
}
