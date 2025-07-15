<script setup>
import { ref } from 'vue'
import { login } from '@/api'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const router = useRouter()

const handleLogin = async () => {
  try {
    const data = await login(username.value, password.value)
    console.log('登录成功:', data)
    router.push('/home')
  } catch (error) {
    console.error('登录出错:', error)
  }
}
</script>

<template>
  <div class="login-container">
    <h2>用户登录</h2>
    <input
      v-model="username"
      type="text"
      placeholder="用户名"
      class="login-input"
    />
    <input
      v-model="password"
      type="password"
      placeholder="密码"
      class="login-input"
    />
    <button @click="handleLogin" class="login-button">登录</button>
  </div>
</template>

<style scoped>
.login-container {
  max-width: 300px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  text-align: center;
}

.login-input {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  box-sizing: border-box;
}

.login-button {
  width: 100%;
  padding: 10px;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.login-button:hover {
  background-color: #3aa876;
}
</style>