<template>
  <div>
    <h1>Einkaufswagen</h1>
    <button @click="clearCart">Warenkorb leeren</button>
    <ul>
      <li v-for="recipe in cartItems.recipes" :key="recipe.id">
        <h2>{{ recipe.bezeichnung }}</h2>
        <p>Menge: {{ recipe.count }}</p>
        <button @click="removeFromCart(recipe.id)">Entfernen</button>
      </li>
    </ul>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  setup() {
    const cartItems = ref({ recipes: [] });

    const fetchCartItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/shopping-cart');
        cartItems.value = response.data;
      } catch (error) {
        console.error(error);
      }
    };

    const removeFromCart = async (id) => {
      try {
        await axios.delete(`http://localhost:8080/api/shopping-cart/${id}`);
        // Nach dem Entfernen das Rezept aus der Anzeige aktualisieren
        fetchCartItems();
      } catch (error) {
        console.error(error);
      }
    };

    const clearCart = async () => {
      try {
        await axios.delete('http://localhost:8080/api/shopping-cart');
        // Nach dem Leeren des Warenkorbs die Anzeige aktualisieren
        fetchCartItems();
      } catch (error) {
        console.error(error);
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
