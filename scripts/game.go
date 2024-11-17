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
		AttemptProgress(1)
		AddToAttempted(input)
	}
	CheckWord()
}

// Guess full word, one-time use
func GuessWord(input string) {
	if OneShot {
		if input == Word {
			Win()
		} else {
			AttemptProgress(2)
		}
		OneShot = !OneShot
	}
}

// Counts down attempts and progresses hangman figure
func AttemptProgress(progress int) {
	Attempts = Attempts - progress
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

func AddToAttempted(letter string) {
	isLetterTried := false
	for i := 0; i < len(AttemptedLetters); i++ {
		if letter == AttemptedLetters[i] {
			isLetterTried = true
		}
	}
	if !isLetterTried {
		AttemptedLetters = append(AttemptedLetters, letter)
		AttemptedDisplay = strings.Join(AttemptedLetters, ", ")
		isLetterTried = false
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
