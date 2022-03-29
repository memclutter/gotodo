<template>
  <div>
    <registration-form/>
  </div>
</template>

<script>
import RegistrationForm from "../components/RegistrationForm";
export default {
  name: "RegistrationPage",
  components: { RegistrationForm },
  meta: {
    isAuth: false
  }
}
</script>
