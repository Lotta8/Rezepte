import { createRouter, createWebHistory } from 'vue-router';
import RecipeDetail from './views/recipeDetail.vue'; // Aktualisiere den Importpfad

const routes = [
    // Andere Routen
    {
        path: '/recipe/:id',
        name: 'recipe-detail',
        component: RecipeDetail,
        props: true,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
