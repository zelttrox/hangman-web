package hangman

var Dictionary string // Dictionnaire de mots choisi
var WordList []string // Dictionnaire des mots

var Word string           // Mot choisi au hasard
var WordProgress []string // Nombre de cases du Hangman
var CurrentWord string    // Mot incomplet avec lettres devines
var LetterAmount int      // Nombre de lettres a reveler

var HangmanPosition []string // Liste de positions du Hangman
var HangmanProgress string   // Position du Hangman

var Input string     // Letter choisie par le joueur
var WordInput string // Mot deviné par le joueur

var Attempts int              // Nombre d'essais restants
var MaxAttempts int           // Nombre total d'essais
var AttemptedLetters []string // Lettres déjà essayées
var AttemptedDisplay string   // Lettres déjà essayées au format string

var Template string // Current HTML template page

var HomePage string // Homescreen page
var MainPage string // Game page
var WinPage string  // Winning page
var LosePage string // Losing page

var IsHardcoreMode bool // Défini si le jeu est en mode hardcore ou en normal
var GuessTried bool
