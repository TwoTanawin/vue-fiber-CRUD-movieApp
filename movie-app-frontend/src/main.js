import { createApp } from 'vue'; // Import createApp for Vue 3
import App from './App.vue';
import axios from 'axios'; // Import axios

const app = createApp(App);

// Set a global default base URL for axios
app.config.globalProperties.$http = axios.create({
  baseURL: 'http://localhost:3000', // Adjust the base URL to your backend API
});

app.mount('#app'); // Mount the app
