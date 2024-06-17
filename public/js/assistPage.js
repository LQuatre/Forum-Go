document.getElementById('contact-form').addEventListener('submit', function(event) {
    event.preventDefault();
    
    // Récupération des données du formulaire
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const message = document.getElementById('message').value;

    // Traitement des données (ex: validation, envoi au serveur)
    console.log('Nom:', name);
    console.log('Email:', email);
    console.log('Message:', message);

    // Afficher un message de confirmation ou traiter les erreurs
    alert('Votre message a été envoyé. Merci de nous avoir contactés !');
    
    // Réinitialiser le formulaire
    document.getElementById('contact-form').reset();
});
