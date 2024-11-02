<template>
    <div :class="['user-profile', { 'compact': compact, 'expanded': expanded }]">
        <div class="profile-content">
            <div class="avatar">
                {{ initials }}
            </div>
            <div class="user-info">
                <div v-if="!compact" class="username">
                    <p v-if="user">{{ user.displayName }}</p>
                    <div v-if="expanded">
                        <p>Username: {{ user.username }}</p>
                        <p>Email: {{ user.email }}</p>
                        <!-- <p>Display Name: {{ user.displayName }}</p> -->
                        <h1>Posts</h1>
                        <div>
                            <div v-for="post in usersPosts" :key="post.idPost" class="post">
                                <p>{{ post.content_text }}</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <!-- Close Button for Full-Screen Mode -->
        <button v-if="expanded" class="close-button" @click="$emit('close')">Close</button>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'UserProfile',

    props: {
        user: {
            type: Object,
            required: true
        },
        compact: {
            type: Boolean,
            default: false
        },
        expanded: {
            type: Boolean,
            default: false
        }
    },

    data() {
        return {
            url: "http://localhost:5050",
            usersPosts: [],
        }
    },

    computed: {
        initials() {
            if (!this.user || !this.user.displayName) return '?'
            return this.user.displayName
                .split(' ')
                .map(word => word[0])
                .join('')
                .toUpperCase()
                .slice(0, 2)
        }
    },

    watch: {
        expanded: {
            immediate: true,
            handler(newVal) {
                if (newVal) {
                    this.getUserPosts()
                }
            }
        }
    },

    methods: {
        async getUserPosts() {
            if (!this.expanded) return  // Prevent fetching if not expanded
            try {
                const response = await axios.get(`${this.url}/posts/user`, {
                    params: {
                        id: this.user.id
                    }
                })
                this.usersPosts = response.data
            } catch (error) {
                console.error('Error fetching user posts:', error)
                this.usersPosts = []
            }
        }
    }
}
</script>

<style scoped>
.user-profile {
    display: flex;
    align-items: center;
    padding: 0.5em;
    transition: all 0.3s ease;
}

.profile-content {
    display: flex;
    align-items: center;
    gap: 0.5em;
}

.avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: #2196f3;
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: bold;
    font-size: 1em;
}

.compact .avatar {
    width: 30px;
    height: 30px;
    font-size: 0.8em;
}

.user-info {
    display: flex;
    flex-direction: column;
}

.display-name {
    font-weight: bold;
    color: #333;
}

.username {
    color: #666;
    font-size: 0.9em;
}

.compact .user-info {
    font-size: 0.9em;
}

.loading,
.error {
    font-size: 0.9em;
    color: #666;
}

.error {
    color: #d32f2f;
}

/* Full-screen mode */
.expanded {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.8);
    color: white;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.expanded .profile-content {
    flex-direction: column;
    align-items: center;
}

.expanded .avatar {
    width: 100px;
    height: 100px;
    font-size: 2em;
}

.expanded .display-name {
    font-size: 1.5em;
}

.expanded .username {
    font-size: 1.2em;
    color: #aaa;
}

/* Close button for full-screen mode */
.close-button {
    position: absolute;
    top: 2rem;
    right: 3rem;
    background: none;
    border: none;
    color: #fff;
    font-size: 1.5rem;
    cursor: pointer;
}
</style>
