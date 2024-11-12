<template>
    <div>
        <NavBar :user="getUserWithId($store.state.userId)"></NavBar>
        <div class="feed-container">
            <div class="posts-section">
                <PostView v-for="post in posts" :key="post.idPost" :post="post" :user="getUserWithId(post.userID)"
                    :users="users" @post-deleted="removePost" />
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import NavBar from './NavBar.vue';
import PostView from './Post.vue';


export default {
    name: 'PostFeed',

    components: {
        NavBar,
        PostView
    },

    data() {
        return {
            posts: [],
            users: [],
            baseUrl: "http://localhost:5050"
        };
    },

    created() {
        this.fetchPosts();
        this.fetchUsers();
    },

    methods: {
        removePost(postId) {
            // Remove the post from the local posts array
            this.posts = this.posts.filter(post => post.idPost !== postId);
        },
        getUserWithId(id) {
            if (this.users.length === 0) return null;
            return this.users.find(user => user.idUser === id);
        },

        async fetchPosts() {
            try {
                const response = await axios.get(`${this.baseUrl}/posts`);
                this.posts = response.data;
            } catch (error) {
                console.error('Error fetching posts:', error);
            }
        },

        async fetchUsers() {
            try {
                const response = await axios.get(`${this.baseUrl}/users`);
                this.users = response.data;
            } catch (error) {
                console.error('Error fetching users:', error);
            }
        }
    }
};
</script>

<style scoped>
.feed-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 1rem;
}
</style>
