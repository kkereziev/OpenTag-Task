package data

import "testing"

func TestDoesWordBeginWithVowel(t *testing.T) {
	english := English{Word: "apples"}

	flag := english.DoesWordBeginWithVowel()

	if flag == false {
		t.Fatalf("The beggining of word %v should be counted as vowel", english.Word)
	}
}

func TestDoesWordBeginWithVowelWithNonVowelBeginnignOfWord(t *testing.T) {
	english := English{Word: "qwerty"}

	flag := english.DoesWordBeginWithVowel()

	if flag != false {
		t.Fatalf("The beggining of word %v should not be counted as vowel", english.Word)
	}
}

func TestDoesWordBeginWithSequence(t *testing.T) {
	english := English{Word: "test"}

	sequence := english.Word[:2]
	flag := english.DoesWordBeginWithSequence(sequence)

	if flag == false {
		t.Fatalf("The begging of the word %v should match %v", english.Word, sequence)
	}
}

func TestDoesWordBeginWithSequenceWithWrongSequence(t *testing.T) {
	english := English{Word: "test"}
	sequence := "ab"

	flag := english.DoesWordBeginWithSequence(sequence)

	if flag != false {
		t.Fatalf("The begging of the word %v should not match %v", english.Word, sequence)
	}
}

func TestDoesWordBeginWithConsonantSound(t *testing.T) {
	english := English{Word: "qwerty"}
	consonantCounter := 0

	flag := english.DoesWordBeginWithConsonantSound(&consonantCounter)

	if flag == false {
		t.Fatalf("The begging of the word %v should start with consonant sound", english.Word)
	}

	if consonantCounter != 2 {
		t.Fatalf("The counter should count the number of consonant sounds at the begging of the word")
	}
}
