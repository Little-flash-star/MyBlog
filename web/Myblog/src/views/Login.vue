<template>
  <div class="login-wrapper">
    <!-- 动态背景 -->
    <div class="bg-animation">
      <div class="floating-shapes">
        <div class="shape shape-1"></div>
        <div class="shape shape-2"></div>
        <div class="shape shape-3"></div>
        <div class="shape shape-4"></div>
        <div class="shape shape-5"></div>
      </div>
    </div>

    <!-- 登录容器 -->
    <div class="login-container">
      <div class="login-card">
        <div class="card-header">
          <div class="logo-container">
            <div class="logo">
              <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                <circle cx="12" cy="12" r="10" fill="url(#gradient)" />
                <path d="M8 12l2 2 4-4" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                <defs>
                  <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="100%">
                    <stop offset="0%" style="stop-color:#667eea"/>
                    <stop offset="100%" style="stop-color:#764ba2"/>
                  </linearGradient>
                </defs>
              </svg>
            </div>
          </div>
          <h1 class="login-title">欢迎回来</h1>
          <p class="login-subtitle">请登录您的账户</p>
        </div>

        <form @submit.prevent="submitLogin" class="login-form">
          <div class="form-group">
            <label for="username" class="form-label">用户名</label>
            <div class="input-container" :class="{ 'focused': usernameFocused, 'error': usernameError, 'success': username && !usernameError }">
              <input 
                v-model="username" 
                type="text" 
                id="username" 
                class="form-input"
                @focus="usernameFocused = true"
                @blur="usernameFocused = false; validateUsername()"
                @input="validateUsername"
                :disabled="isLoading"
                required
              >
              <div class="input-icon">
                <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                  <circle cx="12" cy="7" r="4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </div>
            </div>
            <div v-if="usernameError" class="error-message">{{ usernameError }}</div>
          </div>

          <div class="form-group">
            <label for="password" class="form-label">密码</label>
            <div class="input-container" :class="{ 'focused': passwordFocused, 'error': passwordError, 'success': password && !passwordError }">
              <input 
                v-model="password" 
                :type="showPassword ? 'text' : 'password'" 
                id="password" 
                class="form-input"
                @focus="passwordFocused = true"
                @blur="passwordFocused = false; validatePassword()"
                @input="validatePassword"
                :disabled="isLoading"
                required
              >
              <div class="input-icon">
                <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <rect x="3" y="11" width="18" height="11" rx="2" ry="2" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M7 11V7a5 5 0 0 1 10 0v4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </div>
              <button 
                type="button" 
                class="password-toggle"
                @click="togglePasswordVisibility"
                :disabled="isLoading"
              >
                <svg v-if="showPassword" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                  <path d="M1 1l22 22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
                <svg v-else viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                  <circle cx="12" cy="12" r="3" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                </svg>
              </button>
            </div>
            <div v-if="passwordError" class="error-message">{{ passwordError }}</div>
          </div>

          <div class="form-options">
            <label class="checkbox-container">
              <input v-model="rememberMe" type="checkbox" :disabled="isLoading">
              <span class="checkmark"></span>
              记住我
            </label>
            <a href="#" class="forgot-password">忘记密码？</a>
          </div>

          <button 
            type="submit" 
            class="login-btn"
            :disabled="isLoading || !isFormValid"
            :class="{ 'loading': isLoading }"
          >
            <span v-if="isLoading" class="loading-spinner"></span>
            <span class="btn-text">{{ isLoading ? '登录中...' : '登录' }}</span>
          </button>

          <div v-if="error" class="alert error-alert">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2"/>
              <line x1="15" y1="9" x2="9" y2="15" stroke="currentColor" stroke-width="2"/>
              <line x1="9" y1="9" x2="15" y2="15" stroke="currentColor" stroke-width="2"/>
            </svg>
            {{ error }}
          </div>

          <div v-if="successMessage" class="alert success-alert">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M9 12l2 2 4-4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2"/>
            </svg>
            {{ successMessage }}
          </div>
        </form>

        <div class="signup-link">
          还没有账户？<a href="#">立即注册</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useUserStore } from '../../store'

export default {
  setup() {
    const username = ref('')
    const password = ref('')
    const error = ref('')
    const successMessage = ref('')
    const isLoading = ref(false)
    const showPassword = ref(false)
    const rememberMe = ref(false)
    const usernameError = ref('')
    const passwordError = ref('')
    const usernameFocused = ref(false)
    const passwordFocused = ref(false)
    
    const router = useRouter()
    const userStore = useUserStore()

    const isFormValid = computed(() => {
      return username.value.trim().length > 0 && 
              password.value.length > 0 && 
              !usernameError.value && 
              !passwordError.value
    })

    const togglePasswordVisibility = () => {
      showPassword.value = !showPassword.value
    }

    const validateUsername = () => {
      if (!username.value.trim()) {
        usernameError.value = '用户名不能为空'
      } else if (username.value.trim().length < 3) {
        usernameError.value = '用户名至少需要3个字符'
      } else {
        usernameError.value = ''
      }
    }

    const validatePassword = () => {
      if (!password.value) {
        passwordError.value = '密码不能为空'
      } else if (password.value.length < 6) {
        passwordError.value = '密码至少需要6个字符'
      } else {
        passwordError.value = ''
      }
    }

    const submitLogin = async () => {
      error.value = ''
      successMessage.value = ''
      
      validateUsername()
      validatePassword()
      
      if (!isFormValid.value) return

      isLoading.value = true

      try {
        const res = await axios.post('http://localhost:8080/api/v1/users/login', {
          username: username.value.trim(),
          password: password.value
        })

        console.log('后端返回的数据:', res.data)
        const { token, username: name } = res.data.data

        if (token) {
          successMessage.value = '登录成功！'
          userStore.login(token, name)
          
          setTimeout(() => {
            router.push('/home')
          }, 1000)
        } else {
          error.value = '登录失败，请检查用户名和密码'
        }
      } catch (err) {
        console.error('请求失败:', err)
        error.value = err.response?.data?.message || '登录失败，请检查用户名和密码'
      } finally {
        isLoading.value = false
      }
    }

    return {
      username,
      password,
      error,
      successMessage,
      isLoading,
      showPassword,
      rememberMe,
      usernameError,
      passwordError,
      usernameFocused,
      passwordFocused,
      isFormValid,
      togglePasswordVisibility,
      validateUsername,
      validatePassword,
      submitLogin
    }
  }
}
</script>

<style scoped>
.login-wrapper {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  position: relative;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  user-select: none;
}

/* 动态背景 */
.bg-animation {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.floating-shapes {
  position: relative;
  width: 100%;
  height: 100%;
}

.shape {
  position: absolute;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 50%;
  animation: float 6s ease-in-out infinite;
}

.shape-1 {
  width: 100px;
  height: 100px;
  top: 10%;
  left: 10%;
  animation-delay: 0s;
}

.shape-2 {
  width: 150px;
  height: 150px;
  top: 70%;
  right: 10%;
  animation-delay: 1s;
}

.shape-3 {
  width: 80px;
  height: 80px;
  top: 40%;
  left: 80%;
  animation-delay: 2s;
}

.shape-4 {
  width: 120px;
  height: 120px;
  bottom: 10%;
  left: 30%;
  animation-delay: 3s;
}

.shape-5 {
  width: 60px;
  height: 60px;
  top: 20%;
  right: 30%;
  animation-delay: 4s;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
  }
  50% {
    transform: translateY(-20px) rotate(180deg);
  }
}

/* 登录卡片 */
.login-container {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 420px;
  padding: 20px;
}

.login-card {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 40px;
  box-shadow: 0 25px 45px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  animation: slideIn 0.6s ease-out;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* 头部 */
.card-header {
  text-align: center;
  margin-bottom: 40px;
}

.logo-container {
  margin-bottom: 24px;
}

.logo {
  width: 64px;
  height: 64px;
  margin: 0 auto;
  animation: logoSpin 3s linear infinite;
}

@keyframes logoSpin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

.logo svg {
  width: 100%;
  height: 100%;
}

.login-title {
  font-size: 32px;
  font-weight: 700;
  color: #2d3748;
  margin: 0 0 8px 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.login-subtitle {
  color: #718096;
  font-size: 16px;
  margin: 0;
}

/* 表单样式 */
.login-form {
  margin-bottom: 32px;
}

.form-group {
  margin-bottom: 24px;
}

.form-label {
  display: block;
  margin-bottom: 8px;
  font-size: 14px;
  font-weight: 600;
  color: #4b5563;
}

.input-container {
  position: relative;
  transition: all 0.3s ease;
}

.form-input {
  width: 100%;
  height: 56px;
  padding: 16px 80px 16px 48px;
  border: 2px solid #e5e7eb;
  border-radius: 12px;
  font-size: 16px;
  background: #ffffff;
  transition: all 0.3s ease;
  outline: none;
}

.form-input:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.input-container.error .form-input {
  border-color: #ef4444;
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

.input-container.success .form-input {
  border-color: #10b981;
  box-shadow: 0 0 0 3px rgba(16, 185, 129, 0.1);
}

.input-icon {
  position: absolute;
  left: 16px;
  top: 50%;
  transform: translateY(-50%);
  width: 20px;
  height: 20px;
  color: #9ca3af;
  transition: color 0.3s ease;
}

.input-container.focused .input-icon {
  color: #667eea;
}

.password-toggle {
  position: absolute;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  width: 20px;
  height: 20px;
  background: none;
  border: none;
  color: #9ca3af;
  cursor: pointer;
  transition: color 0.3s ease;
}

.password-toggle:hover {
  color: #667eea;
}

.error-message {
  color: #ef4444;
  font-size: 14px;
  margin-top: 8px;
  display: flex;
  align-items: center;
  gap: 4px;
}

/* 表单选项 */
.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 32px;
  font-size: 14px;
}

.checkbox-container {
  display: flex;
  align-items: center;
  cursor: pointer;
  color: #4b5563;
  gap: 8px;
}

.checkbox-container input {
  display: none;
}

.checkmark {
  width: 18px;
  height: 18px;
  border: 2px solid #d1d5db;
  border-radius: 4px;
  position: relative;
  transition: all 0.3s ease;
}

.checkbox-container input:checked + .checkmark {
  background: #667eea;
  border-color: #667eea;
}

.checkbox-container input:checked + .checkmark:after {
  content: '';
  position: absolute;
  left: 5px;
  top: 2px;
  width: 4px;
  height: 8px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.forgot-password {
  color: #667eea;
  text-decoration: none;
  transition: color 0.3s ease;
}

.forgot-password:hover {
  color: #5a67d8;
}

/* 登录按钮 */
.login-btn {
  width: 100%;
  height: 56px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  border-radius: 12px;
  color: white;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-bottom: 24px;
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(102, 126, 234, 0.3);
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.loading-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top: 2px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* 警告框 */
.alert {
  padding: 12px 16px;
  border-radius: 8px;
  margin-bottom: 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  animation: alertSlide 0.3s ease-out;
}

@keyframes alertSlide {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.error-alert {
  background: #fef2f2;
  color: #dc2626;
  border: 1px solid #fecaca;
}

.success-alert {
  background: #f0fdf4;
  color: #16a34a;
  border: 1px solid #bbf7d0;
}

.alert svg {
  width: 16px;
  height: 16px;
  flex-shrink: 0;
}

/* 分隔线 */
.divider {
  text-align: center;
  margin: 32px 0;
  position: relative;
  color: #9ca3af;
  font-size: 14px;
}

.divider::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: 1px;
  background: #e5e7eb;
}

.divider span {
  background: rgba(255, 255, 255, 0.95);
  padding: 0 16px;
  position: relative;
}

/* 社交登录 */
.social-login {
  margin-bottom: 24px;
}

.social-btn {
  width: 100%;
  height: 48px;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  background: white;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.social-btn:hover {
  border-color: #d1d5db;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.social-btn svg {
  width: 20px;
  height: 20px;
}

/* 注册链接 */
.signup-link {
  text-align: center;
  color: #6b7280;
  font-size: 14px;
}

.signup-link a {
  color: #667eea;
  text-decoration: none;
  font-weight: 600;
}

.signup-link a:hover {
  color: #5a67d8;
}

/* 响应式设计 */
@media (max-width: 480px) {
  .login-container {
    padding: 16px;
  }
  
  .login-card {
    padding: 24px;
  }
  
  .login-title {
    font-size: 28px;
  }
}
</style>
