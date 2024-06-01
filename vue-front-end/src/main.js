import { createApp } from 'vue';
import App from './App.vue';
import router from './services/router';



// Tailwind
import './assets/tailwind.css'; // Tailwind CSS

// Use components from PrimeVue
import PrimeVue from 'primevue/config';
import 'primevue/resources/themes/aura-light-green/theme.css';
import Knob from 'primevue/knob';
import Card from 'primevue/card';
import Dropdown from 'primevue/dropdown';
import Button from 'primevue/button';
import Chart from 'primevue/chart';
import OverlayPanel from 'primevue/overlaypanel';
import SelectButton from 'primevue/selectbutton';



const app = createApp(App);
app.use(router)
app.use(PrimeVue);
app.component('PrimeKnob', Knob);
app.component('PrimeCard', Card);
app.component('PrimeDropdown', Dropdown);
app.component('PrimeButton', Button)
app.component('PrimeChart', Chart);
app.component('PrimeOverlay', OverlayPanel);
app.component('PrimeSelectButton', SelectButton);

app.mount('#app');
