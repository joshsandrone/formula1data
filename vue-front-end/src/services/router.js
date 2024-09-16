// src/router.js
import { createRouter, createWebHistory } from 'vue-router';
import  SeasonStatsDashboard from '@/components/SeasonStatsDashboard.vue';
import  RaceStatsDashboard from '@/components/RaceStatsDashboard.vue';

const routes = [
  {
    path: '/seasonData/:season?/:page?',
    name: 'SeasonStats',
    component: SeasonStatsDashboard,
  },
  {
    path: '/raceData/:season?/:race?',
    name: 'RaceStats',
    component: RaceStatsDashboard,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
