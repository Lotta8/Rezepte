<template>
  <div class="container">
    <h1 class="display-1">Omas Rezepte</h1>
    <router-link :to="{ name: 'rezepte' }">Rezepte</router-link><br>
    <router-link :to="{ name: 'rezeptListe' }">Rezeptliste</router-link><br>
    <router-link :to="{ name: 'einkaufswagen' }">Einkaufswagen</router-link>
    <router-view></router-view>
  </div>
</template>


