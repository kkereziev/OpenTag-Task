package handlers

import (
	"fmt"
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

	req := httptest.NewRequest(http.MethodPost, "/word", strings.NewReader(`{"english_word": "apple"}`))
	w := httptest.NewRecorder()

	th.TranslateWord(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Word translation didn't return %v", http.StatusOK)
	}

	expected := `{"gopher_word":"gapple"}`

	response, err := ioutil.ReadAll(w.Body)

	if err != nil {
		t.Errorf("reading response body: %v, want %v", err, expected)
	}

  fmt.Println()
	if !strings.Contains(string(response), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(response), expected)
	}
}
