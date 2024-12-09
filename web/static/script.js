// Fonction pour générer les boutons de lettres
function generateAlphabetButtons() {

    var wordProgress = document.getElementById('word-progress').innerText
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

    // Disable the button to prevent multiple clicks
    button.disabled = true;

    // Send the letter to the server via AJAX
    fetch('/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `letter=${letter}`
    })

    .then(response => response.json())

    .then(data => {
        // Update the image and progress elements with the new data
        document.getElementById('hangman-image').src = data.Image;
        document.getElementById('word-progress').textContent = data.Progress

        var attemptsLeft = data.Attempts
        checkProgress()
    })

    .catch(error => console.error('Error:', error));
}

function checkProgress() {
    
    wordProgress = document.getElementById('word-progress').innerText
    console.log(wordProgress)

    if (!wordProgress.includes('_') || attemptsLeft <= 0) {
        location.reload()
    }
}

generateAlphabetButtons();
