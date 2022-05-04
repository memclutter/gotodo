import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
      meta: {
        isNavigation: true,
        title: 'Home',
        isAuth: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
      meta: {
        isNavigation: true,
        title: 'Login',
        isAuth: false
      }
    },
    {
      path: '/registration',
      name: 'registration',
      component: () => import('@/views/RegistrationView.vue'),
      meta: {
        isNavigation: true,
        title: 'Registration',
        isAuth: false
      }
    },
    {
      path: '/confirmation/:token',
      name: 'confirmation',
      component: () => import('@/views/ConfirmationView.vue'),
      meta: {
        isNavigation: false,
        title: 'Confirmation',
        isAuth: false
      }
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('@/views/SettingsView.vue'),
      meta: {
        isNavigation: true,
        title: 'Settings',
        isAuth: true
      }
    },
    {
      path: '/logout',
      name: 'logout',
      component: () => import('@/views/LogoutView.vue'),
      meta: {
        isNavigation: true,
        title: 'Logout',
        isAuth: true
      }
    }
  ]
})

export default router
