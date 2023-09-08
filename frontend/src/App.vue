<template>
  <div class="container">
    <h1 class="display-1">Omas Rezepte</h1>
    <router-link to="/rezeptliste">Rezeptliste</router-link>
<br>
    <router-link :to="{ name: 'recipe-detail', params: { id: 1 } }">Rezeptdetails</router-link>
    <router-view></router-view>
    <ul class="list-group mt-3">
      <li
          class="list-group-item d-flex justify-content-between align-items-center"
          v-for="recipe in recipes"
          :key="recipe.id"
      >
        <span>{{ recipe.bezeichnung }}</span>
      </li>
    </ul>
    <form @submit.prevent="addRecipe">
      <div class="input-group mb-3">
        <input
            type="text"
            class="form-control"
            placeholder="Enter a new recipe"
            aria-label="Enter a new recipe"
            aria-describedby="button-add"
            v-model="newRecipeName"
        />
        <button class="btn btn-outline-primary" type="submit" id="button-add">
          Add Recipe
        </button>
      </div>
    </form>
    <ul class="list-group mt-3">
      <li
          class="d-flex justify-content-between list-group-item list-group-item-action align-items-center"
          v-for="recipe in recipes"
          :key="recipe.id"
      >
        <span>
          <span :class="{ 'completed': recipe.done }">{{ recipe.bezeichnung }}</span>
        </span>
        <span>
          <button class="btn" @click="deleteRecipe(recipe.id)">🗑</button>
        </span>
      </li>
    </ul>
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const recipes = ref([]);
    const newRecipeName = ref('');

    const fetchRecipes = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/recipe/all');
        recipes.value = response.data;
      } catch (error) {
        console.error(error);
      }
    };

    const addRecipe = async () => {
      if (newRecipeName.value.trim() !== '') {
        const newRecipe = {
          bezeichnung: newRecipeName.value.trim(),
          done: false
        };
        try {
          const response = await axios.post('http://localhost:8080/api/recipe', newRecipe);
          const createdRecipe = response.data;
          recipes.value.push(createdRecipe);
          newRecipeName.value = '';
        } catch (error) {
          console.error(error);
        }
      }
    };

    const deleteRecipe = async (id) => {
      try {
        await axios.delete(`http://localhost:8080/api/recipe/${id}`);
        recipes.value = recipes.value.filter(recipe => recipe.id !== id);
      } catch (error) {
        console.error(error);
      }
    };

    onMounted(fetchRecipes);

    return {
      recipes,
      newRecipeName,
      addRecipe,
      deleteRecipe
    };
  }
};
</script>

<style>
.completed {
  text-decoration: line-through;
}
</style>
