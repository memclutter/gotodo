<template>
  <div>
    <login-form/>
  </div>
</template>

<script>
import LoginForm from "../components/LoginForm";
export default {
  name: "LoginPage",
  components: { LoginForm },
  meta: {
    isAuth: false
  }
}
</script>
