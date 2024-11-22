// Fonction pour générer les boutons de lettres
function generateAlphabetButtons() {
    const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    const alphabetContainer = document.getElementById('alphabet');

    // Création d'un bouton pour chaque lettre
    alphabet.split('').forEach(letter => {
        const button = document.createElement('input');
        button.innerHTML = letter;
        button.classList.add('letter-btn');
        button.type = "button";
        button.name = "letter";
        button.value = letter;
        button.addEventListener('click', function () {
            handleLetterClick(button, letter);
        });
        alphabetContainer.appendChild(button);
    });
}

// Fonction qui gère ce qui se passe lorsqu'on clique sur une lettre
function handleLetterClick(button, letter) {
    // Désactiver le bouton pour éviter de cliquer plusieurs fois
    button.disabled = true;

    // Affichage d'un message de la lettre choisie
    const message = document.getElementById('message');
    message.textContent = `Vous avez cliqué sur la lettre: ${letter}`;

    // Envoi de la lettre au serveur via AJAX
    fetch('/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `letter=${letter}`
    })
    .then(response => response.text())
    .then(data => {
        console.log(data);
        // Vous pouvez ajouter ici la logique pour mettre à jour l'interface utilisateur
    })
    .catch(error => console.error('Error:', error));
}

// Initialisation du jeu et génération des boutons
generateAlphabetButtons();