package acronym

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

const CHAIN_LAST_VALUE = "<EOL>"

type MarkovChain struct {
	Transitions map[string][]string `json:"transitions"`
	StartWords  []string            `json:"startWords"`
}

func CreateMarkovChain() *MarkovChain {

	// Initialize a new Markov Chain, return address
	return &MarkovChain{
		Transitions: make(map[string][]string),
		StartWords:  make([]string, 0),
	}
}

func (markovChain *MarkovChain) AddCorpus(sentences []string) {

	for _, sentence := range sentences {
		markovChain.addSentence(sentence)
	}

}

// func (markovChain *MarkovChain) addStartWords() {
// 	markovChain.memStartWords = make([]string, 0, len(markovChain.StartWords))
// 	for word := range markovChain.StartWords {
// 		markovChain.memStartWords = append(markovChain.memStartWords, word)
// 	}
// }

// MarkovChain.GenerateSentence(startWord string, length int) -> string
func (markovChain *MarkovChain) GenerateSentence(minLength int, maxLength int) string {
	sentence := make([]string, 0, maxLength)
	currentWord := markovChain.getRandomStartWord()
	wordCount := 0

	for i := 0; i < maxLength; i++ {

		if currentWord == CHAIN_LAST_VALUE {
			break
		}

		sentence = append(sentence, currentWord)
		wordCount += 1

		// Pick new random next word
		nextSentence, ok := markovChain.Transitions[currentWord]
		length := len(nextSentence)
		if !ok || length == 0 {
			break
		}

		currentWord = nextSentence[rand.Intn(length)]

		// Check if current word is EOL
		if currentWord == CHAIN_LAST_VALUE && wordCount >= minLength {
			break
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
		if unicode.IsUpper(rune(startWord[0])) && !contains(markovChain.StartWords, startWord) {
			markovChain.StartWords = append(markovChain.StartWords, startWord)
		}
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

	length := len(markovChain.StartWords)
	if length == 0 {
		return ""
	}
	index := rand.Intn(length)
	return markovChain.StartWords[index]
}

func contains(slice []string, word string) bool {
	for _, w := range slice {
		if w == word {
			return true
		}
	}
	return false
}
