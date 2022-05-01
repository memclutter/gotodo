<script setup lang="ts">
import {Edit} from "@element-plus/icons-vue";

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
      <el-button>
        <el-icon :size="12" :color="'black'">
          <edit />
        </el-icon>
      </el-button>
    </template>
    <div>
      {{description}}
    </div>
  </el-card>
</template>
