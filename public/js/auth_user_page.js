const usernameElement = document.getElementById('username');
        const emailElement = document.getElementById('email');
        const phoneElement = document.getElementById('phone');

        // Exemple de données d'utilisateur
        const userData = {
            username: '',
            email: '',
        };

        usernameElement.textContent = userData.username;
        emailElement.textContent = userData.email;