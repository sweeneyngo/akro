package acronym

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateCorpus(t *testing.T) {
	// Create a temporary directory for testing
	dir := t.TempDir()

	// Create test files
	files := []struct {
		name    string
		content string
	}{
		{"file1.txt", "This is the first sentence. And this is the second."},
		{"file2.txt", "Another file with sentences! Here is one more."},
		{"ignore.csv", "This should be ignored."},
	}

	for _, file := range files {
		err := os.WriteFile(filepath.Join(dir, file.name), []byte(file.content), 0644)
		if err != nil {
			t.Fatalf("failed to write file %s: %v", file.name, err)
		}
	}

	sentences, err := CreateCorpus(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedSentences := []string{
		"This is the first sentence",
		"And this is the second",
		"Another file with sentences",
		"Here is one more",
	}
	for i, sentence := range expectedSentences {
		if i >= len(sentences) || sentences[i] != sentence {
			t.Errorf("expected sentence %d to be %q, got %q", i, sentence, sentences[i])
		}
	}
}

func TestCreateCorpus_ErrorReadingDirectory(t *testing.T) {
	// Test with a directory that does not exist
	dir := "nonexistent_directory"

	_, err := CreateCorpus(dir)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestReadTextFile(t *testing.T) {

	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile_*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	expectedContent := "Hello, world!"
	_, err = file.WriteString(expectedContent)
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	file.Close()

	content, err := readTextFile(file.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if content != expectedContent {
		t.Errorf("expected %q, got %q", expectedContent, content)
	}
}

func TestReadTextFile_FileDoesNotExist(t *testing.T) {

	_, err := readTextFile("nonexistent_file.txt")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSplitTextIntoSentences(t *testing.T) {
	text := "This is the first sentence. This is the second! Is this a question? Yes, it is. This one has dashes- and quotes “like these”."
	expected := []string{
		"This is the first sentence",
		"This is the second",
		"Is this a question",
		"Yes it is",
		"This one has dashes and quotes like these",
	}

	sentences := splitTextIntoSentences(text)
	if len(sentences) != len(expected) {
		t.Fatalf("expected %d sentences, got %d", len(expected), len(sentences))
	}

	for index, expSentence := range expected {
		if sentences[index] != expSentence {
			t.Errorf("expected sentence %d to be %q, got %q", index, expSentence, sentences[index])
		}
	}
}

func TestSplitTextIntoSentences_EmptyString(t *testing.T) {
	text := ""
	expected := []string{}

	sentence := splitTextIntoSentences(text)
	if len(sentence) != len(expected) {
		t.Fatalf("expected %d sentences, got %d", len(expected), len(sentence))
	}
}
