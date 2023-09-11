<template>
  <div class="cart-container">
    <h1 class="text-center">Einkaufswagen</h1>
    <button @click="clearCart" class="styled-button">Warenkorb leeren</button>

    <div class="cart-list">
      <div v-for="recipe in cartItems" :key="recipe.id" class="cart-item">
        <h2>{{ recipe.bezeichnung }}</h2>
        <ul>
          <li v-for="ingredient in recipe.zutaten" :key="ingredient.id">
            <p>{{ ingredient.menge }} {{ ingredient.einheit }} {{ ingredient.zutat.bezeichnung }}</p>
          </li>
        </ul>
        <p>Menge: {{ recipe.count }}</p>
        <button @click="removeFromCart(recipe.id)" class="styled-button">Entfernen</button>
      </div>
    </div>
  </div>
</template>

<style>
.cart-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 20px;
  background-color: #f4f4f4;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-top: 20px;
}

.cart-list {
  width: 100%;
}

.cart-item {
  background-color: #fff;
  border: 1px solid #ddd;
  padding: 20px;
  margin-bottom: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
}

.cart-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  cursor: pointer;
}

.styled-button {
  padding: 5px 10px;
  background-color: #007BFF;
  color: white;
  border: none;
  cursor: pointer;
  margin-top: 10px;
}
</style>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const cartItems = ref([]);

    const fetchCartItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/shopping-cart');
        cartItems.value = response.data.recipes;
      } catch (error) {
        console.error('Fehler beim Abrufen des Warenkorbs:', error);
      }
    };

    const removeFromCart = async (id) => {
      try {
        await axios.delete(`http://localhost:8080/api/shopping-cart/${id}`);
        fetchCartItems();
      } catch (error) {
        console.error('Fehler beim Entfernen des Rezepts:', error);
      }
    };

    const clearCart = async () => {
      try {
        await axios.delete('http://localhost:8080/api/shopping-cart');
        fetchCartItems();
      } catch (error) {
        console.error('Fehler beim Leeren des Warenkorbs:', error);
      }
    };

    onMounted(() => {
      fetchCartItems();
    });

    return {
      cartItems,
      fetchCartItems,
      removeFromCart,
      clearCart,
    };
  },
};
</script>
