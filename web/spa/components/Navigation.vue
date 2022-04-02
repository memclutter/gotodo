<template>
  <b-nav vertical>
    <template v-for="(item, index) in items">
      <b-nav-item :active="$route.name === item.to.name" :to="item.to" :key="index">{{ item.title }}</b-nav-item>
    </template>
  </b-nav>
</template>

<script>
const items = [
  { title: 'Login', to: { name: 'login' }, isAuth: false },
  { title: 'Registration', to: { name: 'registration' }, isAuth: false },
  { title: 'Tasks', to: { name: 'index' }, isAuth: true },
  { title: 'Logout', to: { name: 'logout' }, isAuth: true }
]

export default {
  name: "Navigation",
  computed: {
    items() {
      const isAuth = this.$store.state.auth.session.user.id !== null;
      return items.filter(item => item.isAuth === isAuth);
    }
  }
}
</script>

<style>
.active {
  color: black;
}
</style>
