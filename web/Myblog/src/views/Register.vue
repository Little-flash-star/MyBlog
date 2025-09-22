<template>
  <div class="container mt-5" style="max-width: 400px;">
    <h2 class="mb-4">注册</h2>
    <form @submit.prevent="submitRegister">
      <div class="mb-3">
        <label for="username" class="form-label">用户名</label>
        <input v-model="username" type="text" id="username" class="form-control" required>
      </div>
      <div class="mb-3">
        <label for="password" class="form-label">密码</label>
        <input v-model="password" type="password" id="password" class="form-control" required>
      </div>
      <button type="submit" class="btn btn-success w-100">注册</button>
    </form>
    <p class="text-danger mt-2" v-if="error">{{ error }}</p>
    <p class="mt-2">已有账号？ <router-link to="/login">去登录</router-link></p>
  </div>
</template>

<script>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

export default {
  setup() {
    const username = ref('')
    const password = ref('')
    const error = ref('')
    const router = useRouter()

    const submitRegister = async () => {
      error.value = ''
      if (!username.value || !password.value) {
        error.value = '用户名和密码不能为空'
        return
      }

      try {
        await axios.post('http://localhost:8080/api/v1/users', {
          username: username.value,
          password: password.value
        })
        // 注册成功，跳转登录页
        router.push('/login')
      } catch (err) {
        error.value = err.response?.data?.message || '注册失败'
      }
    }

    return { username, password, error, submitRegister }
  }
}
</script>

<style scoped>
body {
  background-color: #f8f9fa;
}
</style>
