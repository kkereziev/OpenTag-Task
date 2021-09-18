package data

import "opentag/helpers"

type English struct {
	Word string `json:"english_word"`
}

type Gopher struct {
	Word string `json:"gopher_word"`
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
	flag := false

	for i, word := range e.Word {
		if helpers.BinarySearch(helpers.Vowels[:], word) {
			break
		}

		flag = true
		*consonantCounter++

		if word == 'q' && e.Word[i+1] == 'u' {
		  *consonantCounter++
      break;
		}
	}

	return flag
}
