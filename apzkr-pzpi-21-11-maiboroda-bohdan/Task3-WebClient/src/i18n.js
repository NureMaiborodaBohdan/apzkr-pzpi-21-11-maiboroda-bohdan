import { createI18n } from 'vue-i18n';
import { createApp } from 'vue';
import App from './App.vue';
import en from './i18n/en.json';
import ua from './i18n/ua.json';

const messages = {
    en,
    ua
};

const i18n = createI18n({
    locale: 'en',
    messages
});

const app = createApp(App);
app.use(i18n);
app.mount('#app');

export default i18n;
