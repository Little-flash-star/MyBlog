<template>
  <div class="blog-wrapper dark-mode">
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

    <!-- 主容器 -->
    <div class="blog-container">
      <!-- 顶部导航栏 -->
      <header class="blog-header">
        <div class="blog-title-group">
          <div class="logo">
            <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="12" cy="12" r="10" fill="url(#gradient)" />
              <path d="M8 12l2 2 4-4" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
              <defs>
                <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="100%">
                  <stop offset="0%" style="stop-color:var(--gradient-start)"/>
                  <stop offset="100%" style="stop-color:var(--gradient-end)"/>
                </linearGradient>
              </defs>
            </svg>
          </div>
          <div class="title-text-group">
            <h1 class="blog-title">小航的博客</h1>
            <p class="blog-subtitle">分享你的想法</p>
          </div>
        </div>
        <div class="user-actions">
          <button class="action-btn create-post-btn" @click="createPost">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <line x1="12" y1="8" x2="12" y2="16"></line>
              <line x1="8" y1="12" x2="16" y2="12"></line>
            </svg>
            <span class="btn-text">创建新文章</span>
          </button>
          <div class="user-info">
            <img src="https://placehold.co/40x40/6366f1/ffffff?text=U" alt="用户头像" class="user-avatar">
            <span class="user-name">欢迎, {{ userStore.username }}</span>
            <button class="action-btn logout-btn" @click="logout">退出</button>
          </div>
        </div>
      </header>

      <!-- 主要内容区域 -->
      <div class="main-content-grid">
        <!-- 左侧侧边栏 -->
        <aside class="sidebar left-sidebar">
          <div class="sidebar-card">
            <h3 class="sidebar-title">热门文章</h3>
            <ul class="popular-posts">
              <li v-for="post in popularPosts" :key="post.id" class="popular-post-item">
                <a href="#">{{ post.title }}</a>
                <span class="views">{{ post.views }} 次浏览</span>
              </li>
            </ul>
          </div>
        </aside>

        <!-- 文章列表 -->
        <main class="post-list">
          <div v-if="blogPosts.length === 0" class="empty-state">
            <h2>还没有文章</h2>
            <p>点击上方按钮发布你的第一篇文章。</p>
          </div>
          <div v-else class="posts-grid">
            <div v-for="post in blogPosts" :key="post.id" class="post-card">
              <h2 class="post-title">{{ post.title }}</h2>
              <p class="post-meta">
                作者: {{ post.author }} | 发布于: {{ post.date }}
              </p>
              <p class="post-excerpt">{{ post.excerpt }}</p>
              <div class="card-footer">
                <button class="read-more-btn">
                  阅读全文
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="5" y1="12" x2="19" y2="12"></line>
                    <polyline points="12 5 19 12 12 19"></polyline>
                  </svg>
                </button>
                <div class="post-views">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                    <circle cx="12" cy="12" r="3"></circle>
                  </svg>
                  <span>{{ post.views }}</span>
                </div>
              </div>
            </div>
          </div>
        </main>

        <!-- 右侧侧边栏 -->
        <aside class="sidebar right-sidebar">
          <div class="sidebar-card">
            <h3 class="sidebar-title">文章分类</h3>
            <ul class="categories">
              <li v-for="category in categories" :key="category.id" class="category-item">
                <a href="#">#{{ category.name }}</a>
                <span class="count">{{ category.count }}</span>
              </li>
            </ul>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
// 假设这些来自你的 Vue Store
import { useUserStore } from '../stores/user';
import { useRouter } from 'vue-router';


export default {
  setup() {
    const API_BASE_URL = 'http://localhost:8080';

    const userStore = useUserStore();
    const router = useRouter();

    // 博客文章数据
    const blogPosts = ref([]);
    const popularPosts = ref([]);
    const categories = ref([]);
    
    // 从后端API获取文章列表
    const fetchPosts = async () => {
      try {
        const response = await fetch(`${API_BASE_URL}/articles`);
        if (!response.ok) {
          throw new Error('网络请求失败');
        }
        const data = await response.json();
        
        // --- 数据处理逻辑开始 ---
        const articles = data.articles || [];
        
        // 映射后端数据到前端需要的格式
        const processedPosts = articles.map(article => ({
          id: article.id,
          title: article.title,
          excerpt: article.summary,
          author: '小航', // 后端未提供作者名，使用固定值
          date: new Date(article.created_at).toLocaleDateString(), // 格式化日期
          views: article.views,
          category: article.category ? article.category.Name : '未分类'
        }));

        blogPosts.value = processedPosts;

        // 根据处理后的数据生成热门文章
        popularPosts.value = [...processedPosts]
          .sort((a, b) => b.views - a.views)
          .slice(0, 3);
        
        // 根据处理后的数据生成文章分类
        const categoryCounts = {};
        processedPosts.forEach(post => {
          if (post.category) {
            categoryCounts[post.category] = (categoryCounts[post.category] || 0) + 1;
          }
        });

        categories.value = Object.keys(categoryCounts).map(name => ({
          id: name,
          name: name,
          count: categoryCounts[name]
        }));
        // --- 数据处理逻辑结束 ---

      } catch (error) {
        console.error('获取文章失败:', error);
      }
    };

    onMounted(() => {
      fetchPosts();
    });

    const logout = () => {
      // 假设 userStore 有一个 logout 方法
      // userStore.logout();
      console.log('用户退出登录');
      // router.push('/login');
    };

    // 假设文章标题和内容是从某个表单或编辑器中获取的
    const createPost = async () => {
      // 检查用户是否已登录，获取认证token
      // const token = userStore.token;
      const token = 'your_jwt_token_here'; // 占位符，实际应从store获取

      if (!token) {
        console.error('请先登录以创建文章');
        // router.push('/login');
        return;
      }

      const newArticle = {
        title: '这是一个新的API测试文章',
        content: '这是通过POST请求创建的一篇文章。',
        author: userStore.username,
        // 其他字段...
      };
      
      try {
        const response = await fetch(`${API_BASE_URL}/articles`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          },
          body: JSON.stringify(newArticle)
        });

        if (!response.ok) {
          throw new Error('创建文章失败');
        }

        const createdArticle = await response.json();
        console.log('文章创建成功:', createdArticle);
        
        // 创建成功后，重新获取文章列表以更新页面
        await fetchPosts();

      } catch (error) {
        console.error('创建文章失败:', error);
      }
    };

    return {
      userStore,
      blogPosts,
      popularPosts,
      categories,
      logout,
      createPost,
    };
  },
};
</script>

<style scoped>
/* 定义CSS变量 - 白天模式 (默认) */
:root {
  --bg-start: #f0f2f5;
  --bg-end: #e0e4eb;
  --card-bg: rgba(255, 255, 255, 0.9);
  --text-primary: #333;
  --text-secondary: #667085;
  --text-link: #667eea;
  --text-link-hover: #5a64d1;
  --border-color: rgba(0, 0, 0, 0.1);
  --shadow-color: rgba(0, 0, 0, 0.1);
  --divider-color: #e2e8f0;
  --gradient-start: #667eea;
  --gradient-end: #764ba2;
  --cta-bg: #667eea;
  --cta-hover: #5a64d1;
}

/* 暗黑模式覆盖 */
.dark-mode {
  --bg-start: #232526;
  --bg-end: #414345;
  --card-bg: rgba(45, 53, 62, 0.95);
  --text-primary: #f8f9fa;
  --text-secondary: #a0aec0;
  --text-link: #9f7aea;
  --text-link-hover: #805ad5;
  --border-color: rgba(255, 255, 255, 0.1);
  --shadow-color: rgba(0, 0, 0, 0.3);
  --divider-color: #4a5568;
  --gradient-start: #484c98;
  --gradient-end: #7b4397;
  --cta-bg: #9f7aea;
  --cta-hover: #805ad5;
}

.blog-wrapper {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  background: linear-gradient(135deg, var(--bg-start) 0%, var(--bg-end) 100%);
  position: relative;
  overflow: hidden;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  user-select: none;
  transition: background 0.6s ease;
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
  background: var(--border-color);
  backdrop-filter: blur(15px);
  border-radius: 50%;
  animation: float 8s ease-in-out infinite;
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
    transform: translateY(-15px) rotate(180deg);
  }
}

/* 博客主容器 */
.blog-container {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 1200px;
  padding: 20px;
}

/* 顶部导航栏 */
.blog-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: var(--card-bg);
  backdrop-filter: blur(25px);
  border-radius: 24px;
  padding: 20px 40px;
  margin-bottom: 32px;
  box-shadow: 0 4px 12px var(--shadow-color);
  border: 1px solid var(--border-color);
  animation: slideIn 0.6s ease-out;
  transition: all 0.6s ease;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.blog-title-group {
  display: flex;
  align-items: center;
  gap: 16px;
}

.title-text-group {
  display: flex;
  flex-direction: column;
}

.logo {
  width: 48px;
  height: 48px;
}

.logo svg {
  width: 100%;
  height: 100%;
}

.blog-title {
  font-size: 28px;
  font-weight: 700;
  margin: 0;
  background: linear-gradient(135deg, var(--gradient-start) 0%, var(--gradient-end) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1.3;
}

.blog-subtitle {
  font-size: 16px;
  color: var(--text-secondary);
  margin: 0;
}

.user-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.action-btn {
  padding: 10px 20px;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  gap: 8px;
  background: var(--card-bg);
  border: 1px solid var(--divider-color);
  color: var(--text-primary);
  font-size: 14px;
}

.create-post-btn {
  background: var(--cta-bg);
  color: white;
  border: none;
}

.create-post-btn:hover {
  background: var(--cta-hover);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.25);
}

.logout-btn {
  background: #ef4444;
  color: white;
  border: none;
}

.logout-btn:hover {
  background: #dc2626;
  transform: translateY(-1px);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-name {
  font-size: 15px;
  color: var(--text-primary);
  font-weight: 500;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 2px solid var(--text-link);
  box-shadow: 0 3px 8px var(--shadow-color);
}

/* 主要内容网格布局 */
.main-content-grid {
  display: grid;
  grid-template-columns: 280px 1fr 280px;
  gap: 24px;
}

.post-list {
  width: 100%;
}

.posts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
}

.sidebar {
  padding: 20px 0;
}

.sidebar-card {
  background: var(--card-bg);
  backdrop-filter: blur(25px);
  border-radius: 24px;
  padding: 30px;
  box-shadow: 0 8px 24px var(--shadow-color); /* 增强阴影以更好地隔离 */
  border: 1px solid var(--border-color);
  transition: all 0.6s ease;
}

.sidebar-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-primary);
  margin-bottom: 24px;
  position: relative;
  padding-bottom: 8px;
}

.sidebar-title::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 60px;
  height: 2px;
  background: var(--text-link);
  border-radius: 2px;
  opacity: 0.8;
}

.popular-posts, .categories {
  list-style: none;
  padding: 0;
  margin: 0;
}

.popular-post-item {
  margin-bottom: 16px;
}

.popular-post-item a {
  font-size: 15px;
  color: var(--text-primary);
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
  display: block;
  line-height: 1.4;
}

.popular-post-item a:hover {
  color: var(--text-link);
}

.popular-post-item .views {
  font-size: 13px;
  color: var(--text-secondary);
  margin-top: 4px;
  display: block;
}

.category-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.category-item a {
  font-size: 15px;
  color: var(--text-link);
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.category-item a:hover {
  color: var(--text-link-hover);
}

.category-item .count {
  font-size: 12px;
  color: var(--text-secondary);
  background: var(--border-color);
  padding: 3px 8px;
  border-radius: 10px;
}

/* 文章列表和卡片 */
.post-card {
  background: var(--card-bg);
  backdrop-filter: blur(25px);
  border-radius: 24px;
  padding: 30px;
  box-shadow: 0 8px 24px var(--shadow-color); /* 增强阴影以更好地隔离 */
  border: 1px solid var(--border-color);
  transition: transform 0.3s ease, box-shadow 0.3s ease, background 0.6s ease;
}

.post-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 12px 30px var(--shadow-color); /* 增强悬停阴影 */
}

.post-title {
  font-size: 22px;
  font-weight: 600;
  margin-bottom: 8px;
  background: linear-gradient(135deg, var(--gradient-start) 0%, var(--gradient-end) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  line-height: 1.3;
}

.post-meta {
  font-size: 13px;
  color: var(--text-secondary);
  margin-bottom: 16px;
}

.post-excerpt {
  font-size: 15px;
  color: var(--text-primary);
  line-height: 1.6;
  margin-bottom: 24px;
  opacity: 0.9;
}

.card-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 16px;
  border-top: 1px solid var(--divider-color);
}

.read-more-btn {
  background: none;
  border: none;
  color: var(--text-link);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  transition: color 0.3s ease;
}

.read-more-btn:hover {
  color: var(--text-link-hover);
}

.post-views {
  display: flex;
  align-items: center;
  gap: 6px;
  color: var(--text-secondary);
  font-size: 13px;
}

.empty-state {
  text-align: center;
  padding: 80px 20px;
  background: var(--card-bg);
  backdrop-filter: blur(25px);
  border-radius: 24px;
  color: var(--text-primary);
  box-shadow: 0 4px 12px var(--shadow-color);
  border: 1px solid var(--border-color);
  grid-column: 1 / -1;
  transition: all 0.6s ease;
}

.empty-state h2 {
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 12px;
}

.empty-state p {
  font-size: 16px;
  color: var(--text-secondary);
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .main-content-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }
  
  .left-sidebar, .right-sidebar {
    padding: 0;
  }
}

@media (max-width: 768px) {
  .blog-header {
    flex-direction: column;
    align-items: flex-start;
    padding: 20px;
    gap: 16px;
  }
  
  .blog-title-group {
    margin-bottom: 0;
  }
  
  .user-actions {
    width: 100%;
    flex-direction: column;
    align-items: stretch;
    gap: 12px;
  }
  
  .action-btn {
    width: 100%;
    justify-content: center;
  }
  
  .user-info {
    flex-direction: row;
    justify-content: space-between;
    width: 100%;
    gap: 8px;
  }

  .user-name {
    flex-grow: 1;
    text-align: center;
  }

  .left-sidebar, .right-sidebar {
    display: none;
  }
}
</style>
