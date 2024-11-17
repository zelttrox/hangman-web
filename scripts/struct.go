package hangman

var Dictionary string     // Dictionnaire de mots choisi
var WordList []string     // Dictionnaire des mots
var WordChosen string     //
var Word string           // Mot choisi au hasard
var WordProgress []string // Nombre de cases du Hangman
var CurrentWord string    // Mot incomplet avec lettres devines
var FullWord string
var LetterAmount int      // Nombre de lettres a reveler
var OneShot bool

var HangmanPosition []string // Liste de positions du Hangman
var HangmanProgress string   // Position du Hangman
var LinesToAdd int           // Lignes du Hangman à afficher
var HangmanLen int           // Taille d'une frame du Hangman

var Attempts int              // Nombre d'essais restants
var MaxAttempts int           // Nombre total d'essais
var AttemptedLetters []string // Lettres déjà essayées
var LetterTried bool          // Est vrai si la lettre entrée a déjà été 

var IsHardcoreMode bool // Défini si le jeu est en mode hardcore ou en normal

var Input string   // Lettre choisie par le joueur
var IsRunning bool // Est vrai si le jeu est lancé
