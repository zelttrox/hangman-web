package main

import (
hangman "main/scripts"
//web "main/web"
)

func main() {
	hangman.Init()  //Initialisation du jeu
	hangman.Run()   //Lancement de la boucle du jeu
}
