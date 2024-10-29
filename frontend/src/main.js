// main.js - Your entry point
import { createApp } from 'vue'
import App from './App.vue'
import userProfileSinglePage from './components/userProfileSinglePage.vue'
import { createWebHistory, createRouter } from 'vue-router'
import PostFeed from './components/PostFeed.vue'

const routes = [
    { path: '/', component: PostFeed },
    { path: '/user', component: userProfileSinglePage }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const app = createApp(App)
app.use(router)
app.mount('#app')