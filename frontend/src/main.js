// main.js
import { createApp } from 'vue'
import App from './App.vue'
import userProfileSinglePage from './components/userProfileSinglePage.vue'
import { createWebHistory, createRouter } from 'vue-router'
import PostFeed from './components/PostFeed.vue'
import { createStore } from 'vuex'
// import AddPost from './components/AddPost.vue'
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:5050';
axios.defaults.headers.common['Content-Type'] = 'application/json';
axios.defaults.withCredentials = true;

const routes = [
    { path: '/', component: PostFeed },
    { path: '/user', component: userProfileSinglePage },
    // { path: '/post', component: AddPost }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

const store = createStore({
    state() {
        return {
            userId: 1,
            currentUser: null
        }
    },
    mutations: {
        setUserId(state, userId) {
            state.userId = userId
        },
        setCurrentUser(state, user) {
            state.currentUser = user
        }
    },
    getters: {
        getCurrentUser: (state) => state.currentUser
    }
})

const app = createApp(App)
app.use(router)
app.use(store)

// enable CORS
app.config.globalProperties.$http = axios


// Provide the store to all components
app.config.globalProperties.$store = store

app.mount('#app')