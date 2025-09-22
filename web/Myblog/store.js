// store.js
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    username: ''
  }),
  actions: {
    login(token, username) {
      this.token = token
      this.username = username
      localStorage.setItem('token', token)
      console.log('用户 token 已更新:', this.token)
    },
    logout() {
      this.token = ''
      this.username = ''
      localStorage.removeItem('token')
    }
  }
})
