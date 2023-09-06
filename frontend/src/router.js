import { createRouter, createWebHistory } from 'vue-router';
import RecipeDetail from './views/recipeDetail.vue';


const routes = [
    // Andere Routen
    {
        path: '/recipe/:id',
        name: 'recipe-detail',
        component: RecipeDetail,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
