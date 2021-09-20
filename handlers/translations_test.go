package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"opentag/helpers"
	"os"
	"strings"
	"testing"
)

var th TranslationHandler

func init() {
	th = NewTranslationHandler(log.New(os.Stdout, "opentag-task ", log.LstdFlags), helpers.NewCodec())
}

func TestTranslateWord(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/word", strings.NewReader(`{"english_word": "apple"}`))
	w := httptest.NewRecorder()

	th.TranslateWord(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Word translation didn't return %v", http.StatusOK)
	}

	expected := `{"gopher_word":"gapple"}`

	response, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("Reading response body: %v, want %v", err, expected)
	}

	if !strings.Contains(string(response), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(response), expected)
	}
}

func TestTranslateSentence(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/sentence", strings.NewReader(`{"english_sentence": "absolutely marvalous!"}`))
	w := httptest.NewRecorder()

	th.TranslateSentence(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Sentence translation didn't return %v", http.StatusOK)
	}

	expected := `{"gopher_sentence":"gabsolutely arvalous!"}`

	response, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("Reading response body: %v, want %v", err, expected)
	}

	if !strings.Contains(string(response), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(response), expected)
	}
}

func TestGetHistory(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/history", nil)
	w := httptest.NewRecorder()
	th.GetHistory(w, r)

  // already in the db due to our previous two tests
	expected := `{"history":[{"absolutely marvalous!":"gabsolutely arvalous!"},{"apple":"gapple"}]}`
	response, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("Reading response body: %v, want %v", err, "")
	}

	if !strings.Contains(string(response), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(response), expected)
	}
}
