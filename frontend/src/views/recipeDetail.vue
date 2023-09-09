<template>
  <div>
    <h1>Rezepte</h1>
    <!-- ... deine bestehende Rezeptanzeige ... -->

    <!-- Eingabefelder für ID und Menge -->
    <div>
      <input v-model="recipeId" placeholder="Rezept ID" />
      <input v-model="recipeCount" placeholder="Menge" />
      <button @click="addToCart">Zum Einkaufswagen hinzufügen</button>
    </div>

    <!-- Anzeige der hinzugefügten Rezepte -->
    <div>
      <h2>Hinzugefügte Rezepte:</h2>
      <ul>
        <li v-for="cartItem in cartItems" :key="cartItem.id">
          {{ cartItem.name }} ({{ cartItem.count }}x)
        </li>
      </ul>
    </div>
  </div>
</template>

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
      // ... deine bestehende Funktion zum Laden der Rezepte ...
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

    // ... andere bestehende Funktionen ...

    return {
      recipes,
      recipeId,
      recipeCount,
      fetchRecipes,
      addToCart,
      cartItems, // Hinzugefügte Rezepte
    };
  },
};
</script>
