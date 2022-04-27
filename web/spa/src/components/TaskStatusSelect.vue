<script setup lang="ts">
import {statuses} from "@/models/tasks";

const props = defineProps({
  placeholder: String,
  value: String
})
</script>

<template>
  <el-select :model-value="value" @input="$emit('input', $event)" :placeholder="placeholder">
    <el-option
      v-for="status in statuses"
      :key="status.value"
      :label="status.title"
      :value="status.value"
    />
  </el-select>
</template>
