<template>
  <div class="container">
    <h1 class="display-1">Rezeptdetails</h1>
    <div v-if="recipe" class="card">
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
      <p>Loading...</p>
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

    const fetchRecipe = async (id) => {
      try {
        const response = await axios.get(`http://localhost:8080/api/recipe/${id}`);
        recipe.value = response.data;
      } catch (error) {
        console.error(error);
      }
    };

    onMounted(() => {
      if (props.recipeId) {
        fetchRecipe(props.recipeId);
      }
    });

    return {
      recipe,
    };
  },
};
</script>

<style scoped>
/* Hier können Sie Ihre CSS-Stile hinzufügen */
</style>
