<script lang="ts" setup>
import tasksList from '@/apis/endpoints/tasks/list'
import {onMounted, ref} from "vue";

const tasks = ref([])
onMounted(async () => {
  tasks.value = await tasksList()
})
</script>
<template>
  Home
</template>
