<template>
    <div>
        <NavBar :user="getUserWithId(1)"></NavBar>
        <div class="feed-container">
            <div class="posts-section">
                <div v-for="post in posts" :key="post.idPost" class="post">
                    <!-- Post Header with User Info -->
                    <div class="post-header">
                        <UserProfile :user="getUserWithId(post.userID)" />
                    </div>

                    <!-- Post Content -->
                    <div class="post-content" @click="toggleComments(post.idPost)">
                        <p>{{ post.content_text }}</p>
                        <small class="post-date">{{ formatDate(post.created_at) }}</small>
                        <button class="toggle-comments">
                            {{ activePostId === post.idPost ? 'Hide Comments' : 'Show Comments' }}
                        </button>
                    </div>

                    <!-- Comments Section -->
                    <div v-if="activePostId === post.idPost" class="comments-section">
                        <div v-if="loadingComments" class="loading">
                            Loading comments...
                        </div>
                        <div v-else-if="commentsError" class="error">
                            {{ commentsError }}
                        </div>
                        <div v-else>
                            <ul v-if="comments.length > 0" class="comments-list">
                                <li v-for="comment in comments" :key="comment.idComment" class="comment">
                                    <!-- Comment Header with User Info -->
                                    <div class="comment-header">
                                        <UserProfile :user="getUserWithId(comment.idUser)" compact />
                                    </div>
                                    <div class="comment-content">
                                        <p>{{ comment.content_text }}</p>
                                        <small class="comment-date">{{ formatDate(comment.created_at) }}</small>
                                    </div>
                                </li>
                            </ul>
                            <p v-else class="no-comments">No comments yet</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import UserProfile from './UserProfile.vue'
import NavBar from './NavBar.vue';

export default {
    name: 'PostFeed',

    components: {
        UserProfile,
        NavBar
    },

    data() {
        return {
            posts: [],
            activePostId: null,
            comments: [],
            users: [],
            loadingComments: false,
            commentsError: null,
            baseUrl: "http://localhost:5050"
        }
    },

    created() {
        this.fetchPosts()
        this.fetchUsers()
    },

    methods: {
        getUserWithId(id) {
            return this.users.find(user => user.idUser === id)
        },

        async fetchPosts() {
            try {
                const response = await axios.get(`${this.baseUrl}/posts`)
                this.posts = response.data
            } catch (error) {
                console.error('Error fetching posts:', error)
            }
        },

        async fetchUsers() {
            try {
                const response = await axios.get(`${this.baseUrl}/users`)
                this.users = response.data
            } catch (error) {
                console.error('Error fetching users:', error)
            }
        },

        async toggleComments(postId) {
            if (this.activePostId === postId) {
                this.activePostId = null
                this.comments = []
            } else {
                this.activePostId = postId
                await this.fetchComments(postId)
            }
        },

        async fetchComments(postId) {
            this.loadingComments = true
            this.commentsError = null

            try {
                const response = await axios.get(`${this.baseUrl}/comments`, {
                    params: {
                        idPost: postId
                    }
                })
                this.comments = response.data
            } catch (error) {
                this.commentsError = 'Failed to load comments: ' + error.message
                console.error('Error fetching comments:', error)
            } finally {
                this.loadingComments = false
            }
        },

        formatDate(dateString) {
            try {
                const date = new Date(dateString)
                return date.toLocaleDateString() + ' ' + date.toLocaleTimeString()
            } catch (err) {
                return dateString
            }
        }
    }
}
</script>

<style scoped>
.feed-container {
    max-width: 800px;
    margin: 0 auto;
    padding: 1rem;
}

.post {
    margin-bottom: 1.5em;
    border: 1px solid #ddd;
    border-radius: 8px;
    overflow: hidden;
}

.post-header {
    padding: 0.5em 1em;
    background-color: #fff;
    border-bottom: 1px solid #eee;
}

.post-content {
    padding: 1em;
    cursor: pointer;
    background-color: #f9f9f9;
    transition: background-color 0.2s;
}

.post-content:hover {
    background-color: #f0f0f0;
}

.post-date {
    display: block;
    color: #666;
    margin-top: 0.5em;
}

.toggle-comments {
    margin-top: 0.5em;
    padding: 0.3em 0.8em;
    background-color: #e0e0e0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.comments-section {
    padding: 1em;
    background-color: #fff;
    border-top: 1px solid #eee;
}

.comment {
    padding: 0.8em;
    border-bottom: 1px solid #eee;
}

.comment-header {
    margin-bottom: 0.5em;
}

.comment-content {
    padding-left: 1em;
}

.loading,
.error,
.no-comments {
    padding: 1em;
    text-align: center;
}

.error {
    color: #d32f2f;
    background-color: #ffebee;
    border-radius: 4px;
}

.loading {
    color: #666;
    font-style: italic;
}

.no-comments {
    color: #666;
    font-style: italic;
}
</style>