const profileIcon = document.querySelector('.profile-icon');

const dropdownMenu = document.querySelector('.dropdown-menu');

profileIcon.addEventListener('click', function() {
    dropdownMenu.classList.toggle('show');
});

window.addEventListener('click', function(event) {
    // Vérifie si le clic n'était pas sur l'icône de profil et si le menu déroulant est actuellement affiché
    if (!profileIcon.contains(event.target) && dropdownMenu.classList.contains('show')) {
        dropdownMenu.classList.remove('show');
    }
});
