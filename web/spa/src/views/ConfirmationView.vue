<script lang="ts" setup>
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import authConfirmation from '@/apis/endpoints/auth/confirmation'
import {ElMessage} from "element-plus";

const route = useRoute()
const router = useRouter()
const token: String = route.params.token.toString();

const confirmationLoading = ref(false)

// @TODO not found
onMounted(async () => {
  confirmationLoading.value = true
  const {success} = await authConfirmation({token: token})
  if (success) {
    ElMessage({
      type: 'success',
      message: 'Confirmation success, please login'
    })

    confirmationLoading.value = false;
    await router.push({
      name: 'login'
    })
  } else {
    confirmationLoading.value = false;
    await router.push({
      name: 'login'
    })
  }
})
</script>
<template>
  <div v-loading.fullscreen.lock="confirmationLoading"/>
</template>
