package acronym

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func CreatePassword(sentence string, capLevel int, noiseLevel int) (string, error) {

	words := strings.Fields(sentence)

	acronym := CreateAcronym(words)
	capAcronym, err := setRandomCaps(acronym, capLevel)

	if err != nil {
		fmt.Println("error randomizing caps:", err)
		return "", err
	}

	noisyAcronym, err := addNoise(capAcronym, noiseLevel)

	if err != nil {
		fmt.Println("error randomizing noise:", err)
		return "", err
	}

	return strings.Join(noisyAcronym, ""), nil
}

func CreateAcronym(sentence []string) []string {

	var acronym []string
	for _, word := range sentence {
		if len(word) > 0 {
			acronym = append(acronym, strings.ToLower(string(word[0])))
		}
	}

	return acronym
}

func setRandomCaps(acronym []string, capLevel int) ([]string, error) {

	if capLevel < 0 || capLevel > len(acronym) {
		return nil, errors.New("index out of bounds")
	}

	// Randomize list of indices, prevent repeating index
	indices := rand.Perm(len(acronym))

	for index := 0; index < capLevel; index++ {
		pos := indices[index]
		acronym[pos] = strings.ToUpper(acronym[pos])
	}

	return acronym, nil
}

func addNoise(acronym []string, noiseLevel int) ([]string, error) {

	characters := "0123456789!@#$%^&*"

	for index := 0; index < noiseLevel; index++ {
		pos := rand.Intn(len(acronym) + 1)
		char := string(characters[rand.Intn(len(characters))])

		var err error
		acronym, err = insert(acronym, pos, char)

		if err != nil {
			fmt.Println("error inserting character:", err)
			return nil, err
		}

	}
	return acronym, nil
}

func insert(content []string, index int, value string) ([]string, error) {

	if index < 0 || index > len(content) {
		return nil, errors.New("index out of bounds")
	}

	if index == len(content) {
		return append(content, value), nil
	}

	newContent := append(content[:index+1], content[index:]...) // [a, b, c*, c, d]
	newContent[index] = value

	return newContent, nil
}
