<script setup lang="ts">
import TaskCard from "@/components/TaskCard.vue";

const props = defineProps({
  items: Array
})
</script>

<template>
  <task-card v-bind="item" :key="index" v-for="(item, index) in items"/>
</template>
