package hangman

import (
	"bufio"
	"flag"
	"math/rand"
	"os"
	"strings"
	web "main/web"
)

// Import du dictionnaire dans un tableau de string
func GetWordList(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		WordList = append(WordList, scanner.Text())
	}
	return WordList, scanner.Err()
}

// Choix aléatoire du mot à partir de la liste
func GetWord() {
	Word = WordList[rand.Intn(len(WordList))]
	Word = strings.ToUpper(Word)
}

// Scan des arguments choisis dans le terminal
func GetArg() {
	argHardcore := flag.Bool("hard", false, "Hardcore mode")
	flag.Parse()

	if *argHardcore {
		IsHardcoreMode = true
	}
}

// Initialisation des cases du Hangman
func InitWordProgress() {
	for range Word {
		WordProgress = append(WordProgress, " _")
	}
}

func InitHangmanProgress() {
	HangmanPosition = []string{
		"web/static/images/HangmanPositions/Hangman_1.png",
		"web/static/images/hangman-positions/Hangman_2.png",
		"web/static/images/hangman-positions/Hangman_3.png",
		"web/static/images/hangman-positions/Hangman_4.png",
		"web/static/images/hangman-positions/Hangman_5.png",
		"web/static/images/hangman-positions/Hangman_6.png",
		"web/static/images/hangman-positions/Hangman_7.png",
		"web/static/images/hangman-positions/Hangman_8.png",
		"web/static/images/hangman-positions/Hangman_9.png",
		"web/static/images/hangman-positions/Hangman_10.png",
	}
}

// Initialisation of pages

func InitPages() {
	web.MainPage = ""
	web.WinPage = ""
	web.LosePage = ""
}

// Initialisation of fonctions
func Init() {

	IsRunning = false
	OneShot = true

	GetArg()

	Dictionary = "files/league.txt"
	MaxAttempts = 10

	GetWordList(Dictionary)
	GetWord()

	InitWordProgress()
}
