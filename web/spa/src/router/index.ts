import {createRouter, createWebHistory} from 'vue-router'

// @todo fix routes for ts
// @ts-ignore
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('@/views/HomeView.vue')
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/LoginView.vue')
        },
        {
            path: '/registration',
            name: 'registration',
            component: () => import('@/views/RegistrationView.vue')
        },
        {
            path: '/confirmation',
            name: 'confirmation',
            component: () => import('@/views/ConfirmationView.vue')
        }
    ]
})

export default router
