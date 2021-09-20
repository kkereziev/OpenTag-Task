package data

import "opentag/helpers"

const (
	QLetter = 'q'
	ULetter = 'u'
)

type English struct {
	Word     string `json:"english_word"`
	Sentence string `json:"english_sentence"`
}

func (e *English) DoesWordBeginWithVowel() bool {
	return helpers.BinarySearch(helpers.Vowels[:], rune(e.Word[0]))
}

func (e *English) DoesWordBeginWithSequence(sequence string) bool {
	sequenceLength := len(sequence)
	partOfWord := e.Word[:sequenceLength]

	return sequence == partOfWord
}

func (e *English) DoesWordBeginWithConsonantSound(consonantCounter *int) bool {
	beginsWithConsonant := false

	for i, letter := range e.Word {
		if helpers.BinarySearch(helpers.Vowels[:], letter) {
			break
		}

		beginsWithConsonant = true
		*consonantCounter++

		if letter == QLetter && e.Word[i+1] == ULetter {
			*consonantCounter++
			break
		}
	}

	return beginsWithConsonant
}
