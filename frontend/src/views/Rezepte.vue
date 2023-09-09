<template>
  <div>
    <h1>Rezepte</h1>
    <ul>
      <li v-for="recipe in recipes" :key="recipe.id">{{ recipe.bezeichnung }}</li>
    </ul>
  </div>
</template>
<script>
import { onMounted, ref } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const recipes = ref([]);

    const fetchRecipes = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/recipe/all');
        recipes.value = response.data;
      } catch (error) {
        console.error(error);
      }
    };

    onMounted(fetchRecipes);

    return {
      recipes
    };
  }
};
</script>