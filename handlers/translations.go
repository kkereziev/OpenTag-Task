package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"opentag/data"
	"opentag/helpers"
)

type History struct {
	WordHistory []map[string]string `json:"history"`
}

type TranslationHandler interface {
	TranslateWord(w http.ResponseWriter, r *http.Request)
}

type translationHandler struct {
	logger *log.Logger
	codec  helpers.Codec
}

func NewTranslationHandler(logger *log.Logger, codec helpers.Codec) TranslationHandler {
	return &translationHandler{logger, codec}
}

// TODO: history implement
// func (th *translationHandler) Translate(w http.ResponseWriter, r *http.Request) {
// 	a := &English{}

// 	err := json.NewDecoder(r.Body).Decode(a)

// 	if err != nil {
// 		th.logger.Fatal(err)
// 	}

// 	q := make(map[string]map[string]string)
// 	q[a.Word] = map[string]string{
// 		a.Word: a.Word,
// 	}
// 	q["Ivan"] = map[string]string{
// 		"Ivan": "Petkan",
// 	}

// 	re := []map[string]string{q[a.Word], q["Ivan"]}
// 	json.NewEncoder(w).Encode(&History{re})
// }

func (th *translationHandler) TranslateWord(w http.ResponseWriter, r *http.Request) {
	englishWord := &data.English{}

	err := json.NewDecoder(r.Body).Decode(englishWord)

	if err != nil {
		th.logger.Fatal(err)
	}

	gopherWord := th.translate(englishWord)

	th.codec.Encode(w, &data.Gopher{Word: gopherWord})
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
