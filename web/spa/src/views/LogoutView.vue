<script setup lang="ts">
import {useAuthStore} from "@/stores/auth";
import {useRouter} from "vue-router";

const authStore = useAuthStore();
const router = useRouter()
authStore.unset()
router.replace({name: 'login'})
</script>
