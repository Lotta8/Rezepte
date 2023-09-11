<template>
  <div>
    <div class="container text-center mt-5">
      <h1 class="display-1 header">Omas Rezepte</h1>
      <div class="mt-4">
        <router-link :to="{ name: 'rezepte' }" class="btn btn-primary">Rezepte</router-link>
        <router-link :to="{ name: 'rezeptListe' }" class="btn btn-primary ml-3">Rezeptliste</router-link>
        <router-link :to="{ name: 'einkaufswagen' }" class="btn btn-primary ml-3">Einkaufswagen</router-link>
      </div>
      <router-view></router-view>
    </div>
    <footer class="text-center mt-5">
      <p>&copy; 2023 Omas Rezepte by Angela und Jan</p><br>
      <p>Projektarbeit GoLang _O-WWI-20-T-b-GoLang </p>
    </footer>
  </div>
</template>

<style>
/* Anpassungen für den Header */
.header {
  background-color: #888888;
  color: white;
  padding: 20px;
  border-radius: 10px 10px 0 0;
}

/* Auseinandernehmen der Buttons */
.btn-primary {
  background-color: #007BFF;
  color: white;
  font-weight: bold;
  margin-right: 20px; /* Abstand zwischen den Buttons */
}

.btn-primary:last-child {
  margin-right: 0; /* Kein Abstand nach dem letzten Button */
}

.btn-primary:hover {
  background-color: #888888;
}

/* Stil für die Fußzeile */
footer {
  background-color: #f8f9fa;
  padding: 20px;
}
</style>
