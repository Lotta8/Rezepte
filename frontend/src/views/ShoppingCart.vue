<template>
  <div>
    <h1>Einkaufswagen</h1>
    <button @click="clearCart">Warenkorb leeren</button>
    <ul>
      <li v-for="recipe in cartItems" :key="recipe.id">
        <h2>{{ recipe.bezeichnung }}</h2>
        <ul>
          <li v-for="ingredient in recipe.zutaten" :key="ingredient.id">
            <p>{{ ingredient.menge }} {{ ingredient.einheit }} {{ ingredient.zutat.bezeichnung }}</p>
          </li>
        </ul>
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
    const cartItems = ref([]); // Ändere die Initialisierung auf ein Array

    const fetchCartItems = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/shopping-cart');
        cartItems.value = response.data.recipes; // Verwende die Rezeptdaten aus der API-Antwort
      } catch (error) {
        console.error('Fehler beim Abrufen des Warenkorbs:', error);
      }
    };

    const removeFromCart = async (id) => {
      try {
        await axios.delete(`http://localhost:8080/api/shopping-cart/${id}`);
        // Nach dem Entfernen das Rezept aus der Anzeige aktualisieren
        fetchCartItems();
      } catch (error) {
        console.error('Fehler beim Entfernen des Rezepts:', error);
      }
    };

    const clearCart = async () => {
      try {
        await axios.delete('http://localhost:8080/api/shopping-cart');
        // Nach dem Leeren des Warenkorbs die Anzeige aktualisieren
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
