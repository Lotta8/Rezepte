// router.js
import { createRouter, createWebHistory } from 'vue-router';
import RecipeDetail from '@/views/recipeDetail.vue';

const routes = [
    // Andere Routen, falls vorhanden
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
