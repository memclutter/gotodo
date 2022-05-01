<script setup lang="ts">
import {Edit, Delete} from "@element-plus/icons-vue";

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
      <el-row justify="space-between" align="center">
        <el-col :span="18">
          {{title}}
        </el-col>
        <el-col :span="6" style="text-align: right">
          <el-button>
            <el-icon>
              <edit />
            </el-icon>
          </el-button>
          <el-button>
            <el-icon>
              <delete/>
            </el-icon>
          </el-button>
        </el-col>
      </el-row>
    </template>
    <div>
      {{description}}
    </div>
  </el-card>
</template>
