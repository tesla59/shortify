package utils

import (
	"github.com/tesla59/golexicon"
)

func GenerateWord() string {
	return golexicon.NewLexicon().Generate()
}
