<script setup lang="ts">
import {Delete, Edit} from "@element-plus/icons-vue";
import {tasksDelete} from "@/apis/endpoints/tasks"
import {useTasksStore} from "@/stores/tasks";

const props = defineProps({
  id: Number,
  title: String,
  description: String,
  status: String
})

const tasksStore = useTasksStore()

const del = async () => {
  const {success} = await tasksDelete(<number>props.id)
  if (success) {
    tasksStore.delete(<number>props.id)
  }
}

const update = async () => {
  tasksStore.setUpdateTask(<number>props.id)
  tasksStore.setUpdateDialog(true)
}
</script>

<template>
  <el-card>
    <template #header>
      <el-row justify="space-between" align="middle">
        <el-col :span="18">
          {{ title }}
        </el-col>
        <el-col :span="6" style="text-align: right">
          <el-button @click="update">
            <el-icon>
              <edit/>
            </el-icon>
          </el-button>
          <el-popconfirm
            title="Are you sure to delete this task?"
            @confirm="del"
          >
            <template #reference>
              <el-button>
                <el-icon>
                  <delete/>
                </el-icon>
              </el-button>
            </template>
          </el-popconfirm>
        </el-col>
      </el-row>
    </template>
    <div>
      {{ description }}
    </div>
  </el-card>
</template>
