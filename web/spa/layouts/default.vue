<template>
  <b-container class="pt-5">
    <b-row>
      <b-col>
        <navigation/>
      </b-col>
      <b-col>
        <b-card>
          <nuxt/>
        </b-card>
      </b-col>
    </b-row>
  </b-container>
</template>
<script>
import Navigation from "../components/Navigation";
export default {
  components: { Navigation }
}
</script>
