import { createRouter, createWebHistory } from 'vue-router';
import Layout from '@/layouts/Layout.vue';

const routes = [
  {
    path: '/',
    component: Layout,
    redirect: '',
    children: [
      {
        path: '',
        component: () => import('@/views/home/HomeContent.vue'),
        name: 'home',
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
