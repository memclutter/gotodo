<script setup lang="ts">
import {ref} from "vue";
import authLogin from '@/apis/endpoints/auth/login'

const email = ref('')
const password = ref('')

function submit() {
  const {data, isFinished} = authLogin({email: email.value, password: password.value})
  console.log(data)
}
</script>

<template>
  <form @submit.stop.prevent="submit">
    <label for="email">Email</label>
    <input type="email" id="email" v-model="email">
    <label for="password">Password</label>
    <input type="password" id="password" v-model="password">
    <input type="submit" value="Login"/>
    <pre>
      Email: {{email}}
      Password: {{password}}
    </pre>
  </form>
</template>