<script setup lang="ts">
const props = defineProps({
  id: Number,
  title: String,
  description: String,
  status: String
})
</script>

<template>
  <el-card>
    <template #header>
      {{title}}
    </template>
    <div>
      {{description}}
    </div>
  </el-card>
</template>
