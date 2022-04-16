<script lang="ts" setup>
import tasksList from '@/apis/endpoints/tasks/list'
import {onMounted, ref} from "vue";

const tasks = ref([])
onMounted(async () => {
  const {success, data} = await tasksList()
  if (success) {
    tasks.value = data;
  }
})
</script>
<template>
  Home
</template>
