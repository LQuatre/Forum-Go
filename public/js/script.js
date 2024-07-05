document.addEventListener('DOMContentLoaded', function() {
    const profileIcon = document.getElementById('profile-icon');
    const dropdownMenu = document.getElementById('dropdown-menu');
  
    profileIcon.addEventListener('click', function(event) {
      event.stopPropagation();
      dropdownMenu.classList.toggle('show');
    });
  
    document.addEventListener('click', function(event) {
      if (!profileIcon.contains(event.target)) {
        dropdownMenu.classList.remove('show');
      }
    });
  });
  