<template>
    <div>
        <NavBar :user="user"></NavBar>
        <div class="user-profile">
            <div v-if="loading" class="loading">Loading...</div>
            <div v-else-if="error" class="error">{{ error }}</div>
            <div v-else class="user-info">
                <h2>User Profile</h2>
                <p><strong>ID:</strong> {{ user.idUser }}</p>
                <p><strong>Username:</strong> {{ user.username }}</p>
                <p><strong>Display Name:</strong> {{ user.displayName }}</p>
                <p><strong>Email:</strong> {{ user.email }}</p>
            </div>
        </div>

        <div class="user-posts">
            <h2>Posts</h2>
            <div v-if="loading" class="loading">Loading...</div>
            <div v-else-if="error" class="error">{{ error }}</div>
            <div v-else>
                <PostView @post-deleted="removePost" v-for="post in userPosts" :key="post.idPost" :post="post" :user="user" :users="users "></PostView>
            </div>
        </div>

    </div>
</template>

<script>
import axios from 'axios'
import NavBar from './NavBar.vue';
import PostView from './Post.vue';

export default {
    name: 'UserProfile',

    components: {
        NavBar,
        PostView
    },
    data() {
        return {
            user: {},
            loading: true,
            error: null,
            userPosts: [],
            users: [],
            baseUrl: "http://localhost:5050"
        }
    },

    methods: {
        removePost(postId) {
            // Remove the post from the local posts array
            this.userPosts = this.userPosts.filter(post => post.idPost !== postId);
        }
    },

    async created() {
        try {
            // Fetch user data from the API
            const response = await axios.get(`${this.baseUrl}/user`, {
                params: { id: this.$store.state.userId }
            })
            this.user = response.data
        } catch (error) {
            console.error('Error fetching user data:', error)
            this.error = 'Failed to load user data'
        } finally {
            this.loading = false
        }

        try {
            // Fetch user posts from the API
            const response = await axios.get(`${this.baseUrl}/posts/user`, {
                params: { id: this.$store.state.userId }
            })
            this.userPosts = response.data
            console.log(this.userPosts)
        } catch (error) {
            console.error('Error fetching user posts:', error)
            this.userPosts = []
        }

        try {
                const response = await axios.get(`${this.baseUrl}/users`);
                this.users = response.data;
                console.log(this.users)
            } catch (error) {
                console.error('Error fetching users:', error);
            }
    },
}
</script>

<style scoped>
.user-profile {
    max-width: 800px;
    margin: 0 auto;
    padding: 1rem;
    background-color: #f9f9f9;
    border-radius: 8px;
    border: 1px solid #ddd;
}

.user-profile-nav {
    margin-bottom: 1em;
    padding: 0.5em 1em;
    background-color: #fff;
    border-bottom: 1px solid #eee;
    border-radius: 8px 8px 0 0;
}

.user-info {
    padding: 1em;
    background-color: #f9f9f9;
}

.user-info p {
    margin: 0.5em 0;
    color: #333;
}

h2 {
    margin: 0;
    color: #333;
}

.loading {
    padding: 1em;
    text-align: center;
    color: #666;
    font-style: italic;
}

.error {
    padding: 1em;
    text-align: center;
    color: #d32f2f;
    background-color: #ffebee;
    border-radius: 4px;
}
</style>
