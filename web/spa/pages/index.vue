<template>
  <div>
    Index
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage',
  async mounted() {
    await this.$store.dispatch('tasksList')
  }
})
</script>
