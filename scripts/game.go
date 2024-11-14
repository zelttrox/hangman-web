package hangman

// Guess one letter at a time
func GuessLetter(input string) {
	Guess := false
	for i, WordLetter := range Word {
		if input == string(WordLetter) {
			Guess = true
			WordProgress[i] = string(input)
		}
	}
	if (!Guess){
		AttemptProgress()
	}
	//CheckWord()
}

// Guess full word, one-time use
func GuessWord(input string) {

}

// Counts down attempts and progresses hangman figure
func AttemptProgress() {
	Attempts--
	HangmanProgress = HangmanPosition[10-Attempts]
}

// func CheckWord() {
// 	if WordProgress == Word {
// 		Win()
// 	} else if !(WordProgress == Word) {
// 		Lose()
// 	}
// }

func LoadPage(page string) {

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