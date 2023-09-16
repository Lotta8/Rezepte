<template>
  <div class="cart-container">
    <h1 class="page-title">Einkaufswagen</h1>

    <!-- Box für Rezeptnamen -->
    <div class="recipe-names">
      <h2 class="recipe-title">Hinzugefügte Rezepte:</h2>
      <ul class="added-recipes">
        <li v-for="recipe in cartItems.recipes" :key="recipe.id">
          {{ recipe.bezeichnung }}
          <button @click="removeFromCart(recipe.id)" class="remove-button">Entfernen</button>
        </li>
      </ul>
    </div>

    <!-- Box für Zutaten -->
    <div class="cart-item">
      <h2 class="recipe-title">Warenkorb</h2>
      <ul class="ingredient-list">
        <li v-for="ingredient in cartItems.ingredients" :key="ingredient.bezeichnung" class="ingredient-item">
          <p class="ingredient-details">
            {{ ingredient.menge }} {{ ingredient.einheit }} {{ ingredient.bezeichnung }}
          </p>
        </li>
      </ul>
    </div>

    <!-- PDF Button -->
    <button @click="downloadPDF" class="download-button">PDF herunterladen</button>

    <button @click="clearCart" class="clear-button">Warenkorb leeren</button>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import html2pdf from 'html2pdf.js';

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

    const downloadPDF = async () => {
      const cartHtml = document.querySelector('.cart-container');

      const pdfOptions = {
        margin: 10,
        filename: 'warenkorb.pdf',
        image: { type: 'jpeg', quality: 0.98 },
        html2canvas: { scale: 2 },
        jsPDF: { unit: 'mm', format: 'a4', orientation: 'portrait' },
      };

      try {
        const pdfFile = await html2pdf().from(cartHtml).set(pdfOptions).outputPdf();
        const blob = new Blob([pdfFile], { type: 'application/pdf' });
        const url = URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = 'warenkorb.pdf';
        a.click();
      } catch (error) {
        console.error('Fehler beim Erstellen des PDFs:', error);
      }
    };

    onMounted(() => {
      fetchCartItems();
    });

    return {
      cartItems,
      removeFromCart,
      clearCart,
      downloadPDF,
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

.added-recipes {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 50px;
  background-color: #f4f4f4;
  border-radius: 80px;
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
.download-button {
  padding: 5px 10px;
  background-color: #28a745;
  color: white;
  border: none;
  cursor: pointer;
  margin-bottom: 20px;
}
</style>
