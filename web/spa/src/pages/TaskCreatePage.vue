<template>
  <q-page>
    <q-form @submit="onSubmit" class="q-gutter-md">
      <q-input type="text" label="Title" />
      <q-input type="textarea" label="Body" />
      <div>
        <q-btn type="submit" color="primary">
          Submit
        </q-btn>
        <q-btn type="button" :to="{name: 'task-list'}" color="secondary">
          Cancel
        </q-btn>
      </div>
    </q-form>
  </q-page>
</template>

<script lang="ts" setup>
import {ref} from "vue";
import {api} from "boot/axios";

const title = ref('')
const body = ref('')

const onSubmit = async () => {
  const createResponse = await api.post('/tasks', {title: title.value, body: body.value})
  console.log(createResponse)
}
</script>
