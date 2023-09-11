<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3 class="card-title">Alle Rezepte</h3>
          </div>
          <div class="card-body">
            <ul>
              <li v-for="(recipe, index) in recipes" :key="recipe.id">
                {{ displayedIndex(index) }} {{ recipe.bezeichnung }}
              </li>
            </ul>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3 class="card-title">Rezept suchen</h3>
          </div>
          <div class="card-body">
            <div class="input-container">
              <label for="recipeSearchId">Rezept-ID:</label>
              <input id="recipeSearchId" v-model="recipeSearchId" placeholder="Rezept ID" />
              <button @click="searchRecipe" class="btn btn-primary">Suchen</button>
              <button @click="clearRecipe" class="btn btn-danger">Rezept entfernen</button>
            </div>
            <div>
              <div v-if="recipeNotFound">
                <p>Das gesuchte Rezept wurde nicht gefunden.</p>
              </div>
              <div v-else-if="recipeDetails">
                <h3>
                  {{ recipeDetails.name }}
                  <button @click="toggleRecipeDetails" class="btn btn-primary">
                    {{ showRecipeDetails ? "Einklappen" : "Ausklappen" }}
                  </button>
                </h3>
                <ul v-show="showRecipeDetails">
                  <li v-for="ingredient in recipeDetails.ingredients" :key="ingredient.zutat.id">
                    {{ ingredient.menge }} {{ ingredient.einheit }} {{ ingredient.zutat.bezeichnung }}
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row mt-4">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3 class="card-title">Rezept zum Einkaufswagen hinzufügen</h3>
          </div>
          <div class="card-body">
            <div class="input-container">
              <label for="recipeId">Rezeptnummer:</label>
              <input id="recipeId" v-model="recipeId" placeholder="Rezept ID" />
              <label for="recipeCount">Anzahl Personen:</label>
              <input id="recipeCount" v-model="recipeCount" placeholder="Anzahl" />
              <button @click="addToCart" class="btn btn-success">Zum Einkaufswagen hinzufügen</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row mt-4">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3 class="card-title">Hinzugefügte Rezepte</h3>
          </div>
          <div class="card-body">
            <ul>
              <li v-for="cartItem in cartItems" :key="cartItem.id">
                Rezept {{ cartItem.bezeichnung }} {{ cartItem.count }}x hinzugefügt
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.input-container {
  display: flex;
  align-items: center;
  gap: 10px;
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
    const cartItems = ref([]);
    const recipeSearchId = ref('');
    const recipeDetails = ref(null);
    const recipeNotFound = ref(false);
    const showRecipeDetails = ref(true);

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
        recipeNotFound.value = false;
      } catch (error) {
        console.error(error);
        recipeNotFound.value = true;
      }
    };

    const toggleRecipeDetails = () => {
      showRecipeDetails.value = !showRecipeDetails.value;
    };

    const clearRecipe = () => {
      recipeDetails.value = null;
      recipeNotFound.value = false;
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
      recipeNotFound,
      showRecipeDetails,
      toggleRecipeDetails,
      clearRecipe,
    };
  },
};
</script>
