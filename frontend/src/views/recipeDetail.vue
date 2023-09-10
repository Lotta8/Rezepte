<template>
  <div>
    <h1>Rezepteliste</h1>

    <!-- Numerische Aufzählung der Rezepte -->
    <ol>
      <li v-for="(recipe, index) in recipes" :key="recipe.id">
        {{ displayedIndex(index) }}. {{ recipe.bezeichnung }}
      </li>
    </ol>

    <!-- ... deine bestehende Rezeptanzeige ... -->

    <!-- Eingabefelder für ID und Menge mit Beschriftungen -->
    <h2>Rezept zum Einkaufswagen hinzufügen</h2>
    <div class="input-container">
      <label for="recipeId">Rezeptnummer:</label>
      <input id="recipeId" v-model="recipeId" placeholder="Rezept ID" />

      <label for="recipeCount">Anzahl Personen:</label>
      <input id="recipeCount" v-model="recipeCount" placeholder="Anzahl" />

      <button @click="addToCart">Zum Einkaufswagen hinzufügen</button>
    </div>

    <!-- Anzeige der hinzugefügten Rezepte -->
    <div>
      <h2>Hinzugefügte Rezepte:</h2>
      <ul>
        <li v-for="cartItem in cartItems" :key="cartItem.id">
          Rezept {{ cartItem.name }} {{ cartItem.count }}x hinzugefügt
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
/* CSS-Stile für die verbesserte Formatierung der Eingabefelder */
.input-container {
  display: flex;
  flex-direction: column;
  gap: 10px; /* Abstand zwischen den Eingabefeldern */
}

label {
  font-weight: bold;
}
</style>

<script>
import { onMounted, ref } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const recipes = ref([]);
    const recipeId = ref('');
    const recipeCount = ref('');
    const cartItems = ref([]); // Hinzugefügte Rezepte

    const fetchRecipes = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/recipe/all');
        recipes.value = response.data;
      } catch (error) {
        console.error(error);
      }
    };

    const addToCart = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/shopping-cart', {
          id: parseInt(recipeId.value),
          count: parseInt(recipeCount.value),
        });

        // Füge das hinzugefügte Rezept zur Liste der hinzugefügten Rezepte hinzu
        cartItems.value.push({
          id: parseInt(recipeId.value),
          count: parseInt(recipeCount.value),
          name: response.data.recipeName, // Annahme: Die API gibt den Namen des hinzugefügten Rezepts zurück
        });

        console.log(response.data.message); // Zeige die Erfolgsmeldung an
      } catch (error) {
        console.error(error);
      }
    };

    // Funktion zur Anzeige der numerischen Aufzählung
    const displayedIndex = (index) => {
      return index + 1;
    };

    // ... andere bestehende Funktionen ...

    onMounted(fetchRecipes);

    return {
      recipes,
      recipeId,
      recipeCount,
      addToCart,
      cartItems, // Hinzugefügte Rezepte
      displayedIndex // Funktion zur Anzeige der numerischen Aufzählung
    };
  },
};
</script>
