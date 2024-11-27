package hangman

import (
	"fmt"
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
	if !GuessTried {
		if input == Word {
			Win()
		} else {
			AttemptProgress(2)
		}
		GuessTried = true
	}
}

// Counts down attempts and progresses hangman figure
func AttemptProgress(progress int) {
	Attempts = Attempts - progress
	HangmanProgress = HangmanPosition[10-Attempts]
}

// Check if the word progress matches the word to guess
func CheckWord() {
	CurrentWord = strings.Join(WordProgress, "")
	fmt.Println("current word: ", CurrentWord)
	fmt.Println("word to find: ", Word)
	if CurrentWord == Word {
		Win()
	} else if CurrentWord != Word && Attempts <= 0 {
		Lose()
	}
}

// Add the guessed letter to the letters attempted
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

// Load a page on the website
func LoadPage(page string) {
	Template = page
	fmt.Println("webpage is now ", page)
}

// Load home page
func Home() {
	LoadPage(HomePage)
}

// load game
func Play() {
	LoadPage(MainPage)
}

// Load win screen
func Win() {
	LoadPage(WinPage)
	fmt.Println("Win")
}

// Load lose screen
func Lose() {
	LoadPage(LosePage)
	fmt.Println("Lose")
}

// Sets letter and word input back to blank
func ResetInput() {
	Input = ""
	WordInput = ""
}

// Main function to run the game
func Run() {
	if len(WordInput) > 0 {
		GuessWord(WordInput)
	} else if len(Input) > 0 {
		GuessLetter(Input)
	}
	ResetInput()
}
