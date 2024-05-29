// src/router.js
import { createRouter, createWebHistory } from 'vue-router';
import  SeasonStatsDashboard from '@/components/SeasonStatsDashboard.vue'; // Import your component

const routes = [
  {
    path: '/:season?/:driver?', // Optional parameters for season and driver
    name: 'Season',
    component: SeasonStatsDashboard,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
