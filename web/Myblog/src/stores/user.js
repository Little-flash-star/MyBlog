// src/stores/user.js
import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    username: '张三'
  }),
  actions: {
    logout() {
      this.username = '';
    }
  }
});
