<template>
  <div class="container">
    <h1 class="display-1">Rezeptdetails</h1>
    <div v-if="error" class="alert alert-danger">
      {{ error }}
    </div>
    <div v-else-if="recipe" class="card">
      <div class="card-body">
        <h5 class="card-title">{{ recipe.bezeichnung }}</h5>
        <h6 class="card-subtitle mb-2 text-muted">ID: {{ recipe.id }}</h6>
        <p class="card-text">Zutaten:</p>
        <ul>
          <li v-for="ingredient in recipe.zutaten" :key="ingredient.zutat.id">
            {{ ingredient.zutat.bezeichnung }} - {{ ingredient.menge }} {{ ingredient.einheit }}
          </li>
        </ul>
      </div>
    </div>
    <div v-else>
      <div class="alert alert-info">
        Lade...
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  props: ['recipeId'],
  setup(props) {
    const recipe = ref(null);
    const error = ref(null);

    const fetchRecipe = async (id) => {
      try {
        const response = await axios.get(`http://localhost:8080/api/recipe/${id}`);
        recipe.value = response.data;
        error.value = null; // Fehler zurücksetzen, falls erfolgreich
      } catch (err) {
        console.error(err);
        error.value = "Fehler beim Abrufen der Rezeptdetails.";
      }
    };

    onMounted(() => {
      if (props.recipeId) {
        fetchRecipe(props.recipeId);
      }
    });

    return {
      recipe,
      error,
    };
  },
};
</script>

<style scoped>
/* Hier können Sie Ihre CSS-Stile hinzufügen */
</style>
