import { createRouter, createWebHistory } from 'vue-router';
import Rezepte from "./views/Rezepte.vue";
import RecipeDetail from './views/recipeDetail.vue';
import ShoppingCart from './views/ShoppingCart.vue';

const routes = [
    {
        path: '/',
        redirect: '/rezeptliste',
    },
    {
        path: '/rezepte',
        name: 'rezepte',
        component: Rezepte,
    },
    {
        path: '/rezeptliste',
        name: 'rezeptListe',
        component: RecipeDetail,
    },

    {
        path: '/einkaufswagen',
        name: 'einkaufswagen',
        component: ShoppingCart,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
