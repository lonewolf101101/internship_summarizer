import { createVuetify } from 'vuetify';
import 'vuetify/styles'; // Ensure you are using css-loader

const vuetify = createVuetify(); // Replaces new Vuetify()

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(vuetify);
});