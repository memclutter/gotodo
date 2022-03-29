<template>
  <div>
    Index
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage',
  meta: {
    isAuth: true
  },
  async mounted() {
    await this.$store.dispatch('tasksList')
  }
})
</script>
