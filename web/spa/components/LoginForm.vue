<template>
  <b-form @submit.stop.prevent="submit">
    <b-form-group label="Email">
      <b-form-input
        type="email"
        v-model="v$.model.email.$model"
        :state="validateState('email')"
        aria-describedby="email-error"
      />
      <b-form-invalid-feedback id="email-error" :state="validateState('email')">
        {{ validateMessages('email').join(', ') }}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="Password">
      <b-form-input
        type="password"
        v-model="v$.model.password.$model"
        :state="validateState('password')"
      />
      <b-form-invalid-feedback id="password-error" :state="validateState('password')">
        {{ validateMessages('password').join(', ') }}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-button type="submit" variant="primary">Login</b-button>
  </b-form>
</template>

<script>
import useVuelidate from "@vuelidate/core";
import { email, minLength, required } from "@vuelidate/validators";
import { ref } from "@nuxtjs/composition-api";

export default {
  name: "LoginForm",
  setup() {
    const externalResults = ref({});
    return {
      externalResults,
      v$: useVuelidate({ $externalResults: externalResults })
    }
  },
  data() {
    return {
      model: {
        email: null,
        password: null
      }
    }
  },
  validations() {
    return {
      model: {
        email: {required, email},
        password: {required, minLength: minLength(4)}
      }
    }
  },
  methods: {
    validateState(name) {
      const { $dirty, $error } = this.v$.model[name];
      return $dirty ? !$error : null;
    },
    validateMessages(name) {
      return this.v$.model[name].$errors.map(err => err.$message)
    },
    async submit() {
      const isValid = await this.v$.$validate()
      // if (!isValid) {
      //   return;
      // }

      const errors = await this.$store.dispatch('auth/login', this.model)
      if (errors) {
        this.externalResults.value = {model: {email: 'test'}}
        // Object.assign(this.externalResults, {model: {email: ['aaaas']}})
      }
    }
  }
}
</script>
