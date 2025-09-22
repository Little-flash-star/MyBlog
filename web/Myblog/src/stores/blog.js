// src/stores/blog.js
import { defineStore } from 'pinia';

export const useBlogStore = defineStore('blog', {
  state: () => ({
    title: 'My Blog',
    posts: [],
  }),
  actions: {
    addPost(post) {
      this.posts.push(post);
    },
  },
});
