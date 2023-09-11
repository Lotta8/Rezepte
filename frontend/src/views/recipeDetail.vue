<template>
  <div>
    <h1>Rezepteliste</h1>

    <h2>Alle Rezepte</h2>
    <div>
      <ul>
        <li v-for="(recipe, index) in recipes" :key="recipe.id">
          {{ displayedIndex(index) }} {{ recipe.bezeichnung }}
        </li>
      </ul>
    </div>

    <h2>Rezept suchen</h2>
    <!-- Eingabefeld für Rezept-ID -->
    <label for="recipeSearchId">Rezept-ID:</label>
    <input id="recipeSearchId" v-model="recipeSearchId" placeholder="Rezept ID" />
    <button @click="searchRecipe">Suchen</button>

    <!-- Anzeige der Rezeptdetails -->
    <div v-if="recipeDetails">
      <h3>{{ recipeDetails.name }}</h3>
      <ul>
        <li v-for="ingredient in recipeDetails.ingredients" :key="ingredient.zutat.id">
          {{ ingredient.menge }} {{ ingredient.einheit }} {{ ingredient.zutat.bezeichnung }}
        </li>
      </ul>
    </div>

    <h2>Rezept zum Einkaufswagen hinzufügen</h2>
    <div class="input-container">
      <label for="recipeId">Rezeptnummer:</label>
      <input id="recipeId" v-model="recipeId" placeholder="Rezept ID" />

      <label for="recipeCount">Anzahl Personen:</label>
      <input id="recipeCount" v-model="recipeCount" placeholder="Anzahl" />

      <button @click="addToCart">Zum Einkaufswagen hinzufügen</button>
    </div>

    <h2>Hinzugefügte Rezepte:</h2>
    <ul>
      <li v-for="cartItem in cartItems" :key="cartItem.id">
        Rezept {{ cartItem.bezeichnung }} {{ cartItem.count }}x hinzugefügt
      </li>
    </ul>
  </div>
</template>

<!-- ... rest of your code ... -->


<script>
import { onMounted, ref } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const recipes = ref([]);
    const recipeId = ref('');
    const recipeCount = ref('');
    const cartItems = ref([]);
    const recipeSearchId = ref('');
    const recipeDetails = ref(null);

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

        cartItems.value.push({
          id: parseInt(recipeId.value),
          count: parseInt(recipeCount.value),
          bezeichnung: response.data.recipeName,
        });

        console.log(response.data.message);
      } catch (error) {
        console.error(error);
      }
    };

    const searchRecipe = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/api/recipe/${recipeSearchId.value}`);
        recipeDetails.value = response.data;
      } catch (error) {
        console.error(error);
      }
    };

    const displayedIndex = (index) => {
      return index + 1;
    };

    onMounted(fetchRecipes);

    return {
      recipes,
      recipeId,
      recipeCount,
      addToCart,
      cartItems,
      displayedIndex,
      recipeSearchId,
      searchRecipe,
      recipeDetails,
    };
  },
};
</script>
