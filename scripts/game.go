package hangman

import (
	"strings"
)

// Guess one letter at a time
func GuessLetter(input string) {
	Guess := false
	for i, WordLetter := range Word {
		if input == string(WordLetter) {
			Guess = true
			WordProgress[i] = string(input)
		}
	}
	if !Guess {
		AttemptProgress()
	}
	CheckWord()
}

// Guess full word, one-time use
func GuessWord(input string) {

}

// Counts down attempts and progresses hangman figure
func AttemptProgress() {
	Attempts--
	HangmanProgress = HangmanPosition[10-Attempts]
}

func CheckWord() {
	CurrentWord = strings.Join(WordProgress, "")
	if CurrentWord == Word {
		Win()
	} else if !(CurrentWord == Word) {
		Lose()
	}
}

func LoadPage(page string) {
	Template = page
}

func Play() {
	LoadPage(MainPage)
}

func Win() {
	LoadPage(WinPage)
}

func Lose() {
	LoadPage(LosePage)
}
