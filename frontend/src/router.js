import { createRouter, createWebHistory } from 'vue-router';
import App from './App.vue';
import RecipeDetail from './recipeDetail.vue'; // Importieren Sie die recipeDetail-Komponente

const routes = [
    {
        path: '/',
        name: 'Home',
        component: App,
    },
    {
        path: '/recipe/:id', // Definieren Sie einen dynamischen Pfadparameter ":id"
        name: 'RecipeDetail',
        component: RecipeDetail, // Verknüpfen Sie die routeDetail-Komponente
    },
    // Weitere Routen hier hinzufügen, falls benötigt
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
