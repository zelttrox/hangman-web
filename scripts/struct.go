package hangman

var Dictionary string     // Dictionary used for the game
var WordList []string     // Word dictionaries
var Word string           // Word randomly chosen
var WordProgress []string // Advancement of letters guessed in the word
var CurrentWord string    // WordProgress transformed into string
var LetterAmount int      // Amount of letters to reveal

var HangmanPosition []string // List of states of Hangman (images)
var HangmanProgress string   // Currcent state of the Hangman

var Attempts int       // Amount of attempts left
var MaxAttempts int    // Max amount of attempts
var Attempted []string // Letters already tried

var Template string // Current HTML template page
var HomePage string // Homescreen page
var MainPage string // Game page
var WinPage string  // Winning page
var LosePage string // Losing page

var IsHardcoreMode bool // Defines if the game is in hardcore mode

var Input string     // Letter chosen by the player
var WordInput string // Word guessed by the player
