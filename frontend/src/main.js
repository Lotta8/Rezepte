import { createApp } from 'vue';
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";
import App from './App.vue';
import router from './router'; // Importieren Sie den Router

const app = createApp(App);

app.use(router); // Verwenden Sie den Router in Ihrer App

app.mount('#app');
