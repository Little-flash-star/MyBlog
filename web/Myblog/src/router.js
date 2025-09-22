import { createRouter, createWebHistory } from 'vue-router'
import Login from './views/Login.vue'
import Register from './views/register.vue'
import Home from './views/Home.vue'

// 定义路由
const routes = [
  { path: '/', redirect: '/login' },  // 默认重定向到登录页
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { path: '/home', component: Home, meta: { requiresAuth: true } }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const { useUserStore } = await import('../store')  // 动态引入 Pinia store
  const userStore = useUserStore()  // 获取 Pinia store
  
  // 如果路由需要登录且用户没有 token，则跳转登录页
  if (to.meta.requiresAuth && !userStore.token) {
    next('/login')  // 重定向到登录页
  } else {
    next()  // 继续
  }
})

export default router
