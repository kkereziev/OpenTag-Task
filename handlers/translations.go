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
	EnglishWordKeys []string                     `json:"e,omitempty"`
	Db              map[string]map[string]string `json:"b,omitempty"`
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
	return &translationHandler{logger, codec, &History{Db: map[string]map[string]string{}, EnglishWordKeys: []string{}, WordHistory: []map[string]string{}}}
}

func (th *translationHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	a := &History{WordHistory: []map[string]string{}}
	sort.Strings(th.history.EnglishWordKeys)
	res := []map[string]string{}
	for _, key := range th.history.EnglishWordKeys {
		res = append(res, th.history.Db[key])
	}

	a.WordHistory = res
	th.codec.Encode(w, a)
}

func (th *translationHandler) TranslateWord(w http.ResponseWriter, r *http.Request) {
	englishWord := &data.English{}

	err := th.codec.Decode(r, englishWord)

	if err != nil {
		th.logger.Fatal(err)
	}

	gopherWord := th.translate(englishWord)

	if len(th.history.Db[englishWord.Word]) == 0 {
		th.history.EnglishWordKeys = append(th.history.EnglishWordKeys, englishWord.Word)
		th.history.Db[englishWord.Word] = map[string]string{englishWord.Word: gopherWord}
	}

	th.codec.Encode(w, &data.Gopher{Word: gopherWord})
}

func (th *translationHandler) TranslateSentence(w http.ResponseWriter, r *http.Request) {
	englishWord := &data.English{}

	err := th.codec.Decode(r, englishWord)

	if err != nil {
		th.logger.Fatal(err)
	}

	regex := regexp.MustCompile(`[a-z]+`)
	words := regex.FindAllString(englishWord.Sentence, -1)
	var gopherSentence string
	lastIndex := len(englishWord.Sentence) - 1
	sign := englishWord.Sentence[lastIndex]

	for _, word := range words {
		englishWord.Word = word
		gopherWord := th.translate(englishWord)
		gopherSentence += gopherWord + " "
	}

	lastIndex = len(gopherSentence) - 1

	gopherSentence = gopherSentence[:lastIndex] + string(sign)
	if len(th.history.Db[englishWord.Word]) == 0 {
		th.history.EnglishWordKeys = append(th.history.EnglishWordKeys, englishWord.Sentence)
		th.history.Db[englishWord.Sentence] = map[string]string{englishWord.Sentence: gopherSentence}
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
