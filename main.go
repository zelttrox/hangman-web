package main

import (
	"fmt"
	hg "main/scripts"
	web "main/web"
)

func main() {
	web.CreateWebsite()
	fmt.Println("Website launched")
	hg.Init() // Initialisation du jeu
	hg.Run()  // Lancement de la boucle du jeu
}
