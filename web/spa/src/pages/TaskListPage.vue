<template>
  <q-page class="row items-start justify-evenly q-mt-lg">
    <q-btn color="primary" :to="{name: 'task-create'}">
      Create
    </q-btn>
  </q-page>
</template>

<script setup lang="ts">
import {api} from "boot/axios";
import {onMounted} from "vue";

onMounted(async () => {
  const tasksResponse = await api.get('/tasks')
  console.log(tasksResponse)
})
</script>
