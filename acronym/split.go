package acronym

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CreateCorpus(directory string) ([]string, error) {

	var result strings.Builder
	files, err := os.ReadDir(directory)

	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", directory, err)
	}

	for _, file := range files {

		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			path := filepath.Join(directory, file.Name())
			content, err := readTextFile(path)
			if err != nil {
				return nil, fmt.Errorf("error reading file %s: %w", path, err)
			}
			result.WriteString(content + "")
		}
	}

	return splitTextIntoSentences(result.String()), nil
}

func readTextFile(fileName string) (string, error) {

	content, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func splitTextIntoSentences(text string) []string {

	// Define a regular expression for sentence splitting
	re := regexp.MustCompile(`[.!?]`)

	rawSentences := re.Split(text, -1)
	removePunct := regexp.MustCompile(`[":,\-;“”‘’(){}[\]_+=*½°™—\d&]`)
	removePatternsRe := regexp.MustCompile(`\[\d+\]`) // Remove patterns like [1], [2], [3]

	var sentences []string
	for _, sentence := range rawSentences {

		sentence = removePunct.ReplaceAllString(sentence, " ")
		sentence = removePatternsRe.ReplaceAllString(sentence, " ")
		sentence = strings.TrimSpace(sentence)

		if len(sentence) > 0 {
			sentences = append(sentences, sentence)
		}
	}

	return sentences
}
