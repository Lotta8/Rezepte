import { createRouter, createWebHistory } from 'vue-router';
import RecipeDetail from '@/views/recipeDetail.vue'; // Importiere RecipeDetail

const routes = [
    // Andere Routen
    {
        path: '/recipe/:id',
        name: 'recipe-detail', // Dieser Name wird in router-link verwendet
        component: RecipeDetail,
        props: true,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
