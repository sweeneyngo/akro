package acronym

import (
	"strings"
	"testing"
)

func TestCreatePassword(t *testing.T) {
	tests := []struct {
		sentence   string
		noiseLevel int
		capLevel   int
		wantErr    bool
	}{
		{
			"hello world",
			3,
			2,
			false,
		},
		{
			"go programming",
			0,
			1,
			false,
		},
		{
			"test case",
			5,
			0,
			false,
		},
		{
			"",
			1,
			1,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.sentence, func(t *testing.T) {
			password, err := CreatePassword(tt.sentence, tt.capLevel, tt.noiseLevel)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Password: %s", password)
		})
	}
}

func TestCreateAcronym(t *testing.T) {
	tests := []struct {
		sentence []string
		want     []string
	}{
		{
			[]string{"hello", "world"},
			[]string{"h", "w"},
		},
		{
			[]string{"go", "Programming"},
			[]string{"g", "p"},
		},
		{
			[]string{"test", "Case", "phony"},
			[]string{"t", "c", "p"},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.sentence, " "), func(t *testing.T) {
			acronym := CreateAcronym(tt.sentence)
			if !equal(acronym, tt.want) {
				t.Errorf("CreateAcronym() = %v, want %v", acronym, tt.want)
			}
		})
	}
}

func TestSetRandomCaps(t *testing.T) {
	tests := []struct {
		acronym  []string
		capLevel int
		wantLen  int
		wantErr  bool
	}{
		{
			[]string{"a", "b", "c"},
			2,
			3,
			false,
		},
		{
			[]string{"x", "y"},
			0,
			2,
			false,
		},
		{
			[]string{"a", "b", "c"},
			4,
			0,
			true,
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.acronym, ""), func(t *testing.T) {
			acronym, err := setRandomCaps(test.acronym, test.capLevel)
			if (err != nil) != test.wantErr {
				t.Errorf("setRandomCaps() error = %v, wantErr %v", err, test.wantErr)
			}

			if len(acronym) != test.wantLen {
				t.Errorf("setRandomCaps() = %v, want %v", len(acronym), test.wantLen)
			}
		})
	}
}

func TestAddNoise(t *testing.T) {
	tests := []struct {
		acronym    []string
		noiseLevel int
		wantLen    int
		wantErr    bool
	}{
		{
			[]string{"a", "b", "c"},
			2,
			5,
			false,
		},
		{
			[]string{"x", "y"},
			0,
			2,
			false,
		},
		{
			[]string{},
			1,
			1,
			false,
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.acronym, ""), func(t *testing.T) {
			acronym, err := addNoise(test.acronym, test.noiseLevel)
			if (err != nil) != test.wantErr {
				t.Errorf("addNoise() error = %v, wantErr %v", err, test.wantErr)
			}

			if len(acronym) != test.wantLen {
				t.Errorf("addNoise() = %v, want %v", len(acronym), test.wantLen)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		content []string
		index   int
		value   string
		want    []string
		wantErr bool
	}{
		{
			[]string{"a", "b", "c"},
			1,
			"x",
			[]string{"a", "x", "b", "c"},
			false,
		},
		{
			[]string{"a", "b"},
			2,
			"z",
			[]string{"a", "b", "z"},
			false,
		},
		{
			[]string{"a"},
			0,
			"z",
			[]string{"z", "a"},
			false,
		},
		{
			[]string{"a"},
			-1,
			"z",
			nil,
			true,
		},
		{
			[]string{"a"},
			2,
			"z",
			nil,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			content, err := insert(test.content, test.index, test.value)
			if (err != nil) != test.wantErr {
				t.Errorf("insert() error = %v, wantErr %v", err, test.wantErr)
			}

			if !equal(content, test.want) {
				t.Errorf("insert() = %v, want %v", content, test.want)
			}
		})
	}
}

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for index := range a {
		if a[index] != b[index] {
			return false
		}
	}

	return true
}
