<script lang="ts" setup>
import {computed} from "vue";
import type {RouteRecordRaw} from "vue-router";
import {useRouter} from "vue-router";
import {useAuthStore} from "@/stores/auth";

const router = useRouter()
const authStore = useAuthStore()

const menuItems = computed(() => {
  return router.options.routes
    .filter((route: RouteRecordRaw) => route.meta?.isNavigation && route.meta?.isAuth === authStore.isAuth)
    .map((route: RouteRecordRaw) => ({title: route.meta?.title || route.name, to: route.path}))
})
</script>

<template>
  <el-menu
    mode="horizontal"
    router
  >
    <el-menu-item>
      GoTODO
    </el-menu-item>
    <el-menu-item v-for="(menuItem, index) in menuItems" :key="index" :index="menuItem.to">
      {{ menuItem.title }}
    </el-menu-item>
  </el-menu>
</template>
