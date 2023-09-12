<template>
  <div class="cart-container">
    <h1 class="page-title">Einkaufswagen</h1>
    <button @click="clearCart" class="clear-button">Warenkorb leeren</button>

    <div class="cart-list">
      <div v-for="recipe in cartItems.recipes" :key="recipe.id" class="cart-item">
        <h2 class="recipe-title">{{ recipe.bezeichnung }}</h2>
        <p class="recipe-count">Anzahl Personen: {{ recipe.count }}</p>
        <h3>Zutaten:</h3>
        <ul class="ingredient-list">
          <li v-for="ingredient in cartItems.ingredients" :key="ingredient.bezeichnung" class="ingredient-item">
            <p class="ingredient-details">{{ ingredient.menge }} {{ ingredient.einheit }} {{ ingredient.bezeichnung }}</p>
          </li>
        </ul>
        <button @click="removeFromCart(recipe.id)" class="remove-button">Entfernen</button>
      </div>
    </div>
  </div>
</template>


<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const cartItems = ref({
      ingredients: [],
      recipes: [],
    });

    const fetchCartItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/shopping-cart');
        console.log('API Response:', response.data);
        cartItems.value = response.data;
      } catch (error) {
        console.error('Fehler beim Abrufen des Warenkorbs:', error);
      }
    };

    const removeFromCart = async (id) => {
      try {
        await axios.delete(`http://localhost:8080/api/shopping-cart/${id}`);
        await fetchCartItems();
      } catch (error) {
        console.error('Fehler beim Entfernen des Rezepts:', error);
      }
    };

    const clearCart = async () => {
      try {
        await axios.delete('http://localhost:8080/api/shopping-cart');
        await fetchCartItems();
      } catch (error) {
        console.error('Fehler beim Leeren des Warenkorbs:', error);
      }
    };

    onMounted(() => {
      fetchCartItems();
    });

    return {
      cartItems,
      removeFromCart,
      clearCart,
    };
  },
};
</script>

<style scoped>
.cart-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  background-color: #f4f4f4;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-top: 20px;
}

.page-title {
  font-size: 24px;
  text-align: center;
  margin-bottom: 10px;
}

.clear-button {
  padding: 5px 10px;
  background-color: #007BFF;
  color: white;
  border: none;
  cursor: pointer;
  margin-bottom: 20px;
}

.cart-item {
  background-color: #fff;
  border: 1px solid #ddd;
  padding: 20px;
  margin-bottom: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
  width: 100%;
}

.recipe-title {
  font-size: 20px;
}

.recipe-count {
  font-size: 16px;
  margin-top: 5px;
}

.ingredient-list {
  list-style: none;
  padding: 0;
}

.ingredient-item {
  margin-top: 10px;
}

.ingredient-details {
  font-size: 14px;
}

.remove-button {
  padding: 5px 10px;
  background-color: #FF0000;
  color: white;
  border: none;
  cursor: pointer;
  margin-top: 10px;
}
</style>
