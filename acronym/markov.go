package acronym

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const CHAIN_LAST_VALUE = "<EOL>"

type MarkovChain struct {
	Transitions map[string][]string `json:"transitions"`
	StartWords  map[string]bool     `json:"startWords"`
}

func CreateMarkovChain() *MarkovChain {

	// Initialize a new Markov Chain, return address
	return &MarkovChain{
		Transitions: make(map[string][]string),
		StartWords:  make(map[string]bool),
	}
}

func (markovChain *MarkovChain) AddCorpus(sentences []string) {

	for _, sentence := range sentences {
		markovChain.addSentence(sentence)
	}

}

// MarkovChain.GenerateSentence(startWord string, length int) -> string
func (markovChain *MarkovChain) GenerateSentence(minLength int, maxLength int) string {
	var sentence []string
	currentWord := markovChain.getRandomStartWord()
	wordCount := 0

	for i := 0; i < maxLength; i++ {

		if currentWord != CHAIN_LAST_VALUE {
			sentence = append(sentence, currentWord)
			wordCount += 1
		}

		// Pick new random next word
		nextSentence, ok := markovChain.Transitions[currentWord]
		if !ok || len(nextSentence) == 0 {
			return strings.Join(sentence, " ")
		}

		currentWord = nextSentence[rand.Intn(len(nextSentence))]

		// Check if current word is EOL
		if currentWord == CHAIN_LAST_VALUE && wordCount >= minLength {
			return strings.Join(sentence, " ")
		}
	}

	return strings.Join(sentence, " ")
}

func (markovChain *MarkovChain) PrintTransitions() {
	for key, value := range markovChain.Transitions {
		fmt.Printf("%s -> %v\n", key, value)
	}
}

func (markovChain *MarkovChain) SaveToFile(filename string) error {
	data, err := json.Marshal(markovChain)
	if err != nil {
		return fmt.Errorf("error marshalling Markov Chain: %w", err)
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile loads a Markov Chain from a file in JSON format
func LoadFromFile(filename string) (*MarkovChain, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var markovChain MarkovChain
	err = json.Unmarshal(data, &markovChain)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling Markov Chain: %w", err)
	}
	return &markovChain, nil
}

func (markovChain *MarkovChain) addSentence(sentence string) {

	words := strings.Fields(sentence)
	length := len(words)

	if length > 0 {
		// Set start word as candidate start words
		startWord := words[0]
		markovChain.StartWords[startWord] = true
	}

	for index := 0; index < length-1; index++ {
		currentWord := words[index]
		nextWord := words[index+1]

		if _, exists := markovChain.Transitions[currentWord]; !exists {
			markovChain.Transitions[currentWord] = []string{}
		}

		markovChain.Transitions[currentWord] = append(markovChain.Transitions[currentWord], nextWord)
	}

	// Append <EOL> to last word of Markov Chain
	if _, exists := markovChain.Transitions[words[length-1]]; !exists {
		markovChain.Transitions[words[length-1]] = []string{CHAIN_LAST_VALUE}
	}
	markovChain.Transitions[words[length-1]] = append(markovChain.Transitions[words[length-1]], CHAIN_LAST_VALUE)
}

func (markovChain *MarkovChain) getRandomStartWord() string {

	var startWords []string
	for word := range markovChain.StartWords {
		startWords = append(startWords, word)
	}

	if len(startWords) == 0 {
		return ""
	}

	return startWords[rand.Intn(len(startWords))]
}
