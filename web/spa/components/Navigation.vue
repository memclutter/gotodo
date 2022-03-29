<template>
  <b-nav vertical>
    <template v-for="(item, index) in items.filter(item => item.isAuth ? $store.state.auth.user.id !== null : $store.state.auth.user.id === null)">
      <b-nav-item :to="item.to" :key="index">{{item.title}}</b-nav-item>
    </template>
  </b-nav>
</template>

<script>
export default {
  name: "Navigation",
  data: () => ({
    items: [
      {
        title: 'Login',
        to: {
          name: 'login'
        },
        isAuth: false
      },
      {
        title: 'Registration',
        to: {
          name: 'registration'
        },
        isAuth: false
      },
      {
        title: 'Tasks',
        to: {
          name: 'index'
        },
        isAuth: true
      },
      {
        title: 'Logout',
        to: {
          name: 'logout'
        },
        isAuth: true
      }
    ]
  })
}
</script>

<style scoped>

</style>
