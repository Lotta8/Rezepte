<template>
  <div class="recipe-container">
    <h1 class="text-center">Rezepte</h1>

    <div class="recipe-list">
      <div v-for="recipe in recipes" :key="recipe.id" class="recipe-item">
        {{ recipe.bezeichnung }}
      </div>
    </div>
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

<style>
.recipe-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 50px;
  background-color: #f4f4f4;
  border-radius: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.recipe-list {
  width: 50%;
  margin-top: 20px;
}

.recipe-item {
  background-color: #fff;
  border: 1px solid #ddd;
  padding: 25px;
  margin-bottom: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
}

.recipe-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}
</style>