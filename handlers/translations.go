package handlers

import (
	"log"
	"net/http"
	"opentag/data"
	"opentag/helpers"
	"regexp"
	"sort"
)

type History struct {
	WordHistory     []map[string]string          `json:"history"`
	englishWordKeys []string                     `json:"-"`
	db              map[string]map[string]string `json:"-"`
}

type TranslationHandler interface {
	TranslateWord(w http.ResponseWriter, r *http.Request)
	TranslateSentence(w http.ResponseWriter, r *http.Request)
	GetHistory(w http.ResponseWriter, r *http.Request)
}

type translationHandler struct {
	logger  *log.Logger
	codec   helpers.Codec
	history *History
}

func NewTranslationHandler(logger *log.Logger, codec helpers.Codec) TranslationHandler {
	return &translationHandler{logger, codec, &History{db: map[string]map[string]string{}, englishWordKeys: []string{}, WordHistory: []map[string]string{}}}
}

// GetHistory returns all english words and sentences and their coresponding gopher ones
func (th *translationHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	// All and all this part could have been done better
	// with TreeMap but I decided to stick with this custom solution
	sort.Strings(th.history.englishWordKeys)
	resultArray := []map[string]string{}

	for _, key := range th.history.englishWordKeys {
		resultArray = append(resultArray, th.history.db[key])
	}

	th.history.WordHistory = resultArray
	th.codec.Encode(w, th.history)
}

// TranslateWord translates english word into gopher one
func (th *translationHandler) TranslateWord(w http.ResponseWriter, r *http.Request) {
	english := &data.English{}

	err := th.codec.Decode(r, english)

	if err != nil {
		th.logger.Fatal(err)
	}

	gopherWord := th.translate(english)

	if !th.doesKeyExistInDb(english.Word) {
		th.pushKeyIntoDb(english.Word, gopherWord)
	}

	th.codec.Encode(w, &data.Gopher{Word: gopherWord})
}

// TranslateSentence translates english sentence into gopher one
func (th *translationHandler) TranslateSentence(w http.ResponseWriter, r *http.Request) {
	english := &data.English{}

	err := th.codec.Decode(r, english)

	if err != nil {
		th.logger.Fatal(err)
	}

	var gopherSentence string
	regex := regexp.MustCompile(`[a-z]+`)
	words := regex.FindAllString(english.Sentence, -1)
	
	for _, word := range words {
		english.Word = word
		gopherWord := th.translate(english)
		gopherSentence += gopherWord + helpers.EmptySpace
	}
	
	lastIndex := len(english.Sentence) - 1
	sentenceSign := english.Sentence[lastIndex]
	gopherSentence = gopherSentence[:lastIndex] + string(sentenceSign)

	if !th.doesKeyExistInDb(english.Sentence) {
		th.pushKeyIntoDb(english.Sentence, gopherSentence)
	}

	th.codec.Encode(w, &data.Gopher{Sentence: gopherSentence})
}

func (th *translationHandler) translate(word *data.English) string {
	var gopherWord string
	consonantSoundCount := 0

	switch {
	case word.DoesWordBeginWithVowel():
		gopherWord = helpers.VowelLetterPrefix + word.Word

	case word.DoesWordBeginWithSequence("xr"):
		gopherWord = helpers.ConsonantLettersXrPrefix + word.Word

	case word.DoesWordBeginWithConsonantSound(&consonantSoundCount):
		gopherWord = word.Word[consonantSoundCount:] + word.Word[:consonantSoundCount] + helpers.ConsonantSoundSuffix
	}

	return gopherWord
}

func (th *translationHandler) doesKeyExistInDb(key string) bool {
	return len(th.history.db[key]) != 0
}

func (th *translationHandler) pushKeyIntoDb(key string, value string) {
	th.history.englishWordKeys = append(th.history.englishWordKeys, key)
	th.history.db[key] = map[string]string{key: value}
}
