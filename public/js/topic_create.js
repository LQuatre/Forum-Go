// Attendez que le DOM soit entièrement chargé
document.addEventListener('DOMContentLoaded', function() {
    // Ajoutez un gestionnaire d'événement au bouton "Créer un nouveau topic"
    document.getElementById('btnNewTopic').addEventListener('click', function() {
        // Affichez le formulaire pour créer un nouveau topic
        document.getElementById('newTopicForm').style.display = 'block';
    });
});
