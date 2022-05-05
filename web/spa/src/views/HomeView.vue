<script lang="ts" setup>
import tasksList from '@/apis/endpoints/tasks/list'
import {onMounted} from "vue";
import TaskBoard from "@/components/TaskBoard.vue";
import {useTasksStore} from "@/stores/tasks";
import TaskUpdateDialog from "@/components/TaskUpdateDialog.vue";
import {Grid, List} from "@element-plus/icons-vue";

const tasksStore = useTasksStore()

onMounted(async () => {
  const {success, data} = await tasksList()
  if (success) {
    tasksStore.set(data)
  }
})
</script>
<template>
  <el-button-group>
    <el-button :type="tasksStore.listView === 'board' ? 'primary' : 'default'" @click="tasksStore.setListView('board')">
      <el-icon>
        <grid/>
      </el-icon>
    </el-button>
    <el-button :type="tasksStore.listView === 'list' ? 'primary' : 'default'" @click="tasksStore.setListView('list')">
      <el-icon>
        <list/>
      </el-icon>
    </el-button>
  </el-button-group>
  <template v-if="tasksStore.listView === 'board'">
    <task-board :items="tasksStore.items"/>
  </template>
  <template v-else-if="tasksStore.listView === 'list'">
    <task-list :items="tasksStore.items"/>
  </template>
  <task-create-dialog />
  <task-update-dialog />
</template>
