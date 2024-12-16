<template>
    <div>
        <NavBar :user="user"></NavBar>
        <div class="user-profile">
            <div v-if="loading" class="loading">Loading...</div>
            <div v-else-if="error" class="error">{{ error }}</div>
            <div v-else>
                <div class="user-info" v-if="!isEditing">
                    <h2>User Profile</h2>
                    <p><strong>ID:</strong> {{ user.idUser }}</p>
                    <p><strong>Username:</strong> {{ user.username }}</p>
                    <p><strong>Display Name:</strong> {{ user.displayName }}</p>
                    <p><strong>Email:</strong> {{ user.email }}</p>
                    <button @click="toggleEdit">Edit Profile</button>
                </div>
                
                <div class="edit-user-info" v-else>
                    <h2>Edit Profile</h2>
                    <form @submit.prevent="updateProfile">
                        <div class="form-group">
                            <label for="username">Username:</label>
                            <input
                                type="text"
                                id="username"
                                v-model="editableUser.username"
                                required
                            />
                        </div>
                        <div class="form-group">
                            <label for="displayName">Display Name:</label>
                            <input
                                type="text"
                                id="displayName"
                                v-model="editableUser.displayName"
                                required
                            />
                        </div>
                        <div class="form-group">
                            <label for="email">Email:</label>
                            <input
                                type="email"
                                id="email"
                                v-model="editableUser.email"
                                required
                            />
                        </div>
                        <button type="submit" :disabled="updating">
                            {{ updating ? 'Updating...' : 'Save Changes' }}
                        </button>
                        <button type="button" @click="toggleEdit" :disabled="updating">Cancel</button>
                        <p v-if="updateError" class="error">{{ updateError }}</p>
                    </form>
                </div>
            </div>
        </div>

        <div class="user-posts">
            <h2>Posts</h2>
            <div v-if="loadingPosts" class="loading">Loading...</div>
            <div v-else-if="postsError" class="error">{{ postsError }}</div>
            <div v-else>
                <PostView @post-deleted="removePost" v-for="post in userPosts" :key="post.idPost" :post="post" :user="user" :users="users"></PostView>
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
            editableUser: {},
            loading: true,
            error: null,
            userPosts: [],
            loadingPosts: true,
            postsError: null,
            users: [],
            baseUrl: "http://localhost:5050",
            isEditing: false,
            updating: false,
            updateError: null,
        }
    },


    methods: {
        toggleEdit() {
            this.isEditing = !this.isEditing
            if (this.isEditing) {
                // Create a copy of the user data for editing
                this.editableUser = { ...this.user }
            }
        },
        async updateProfile() {
            this.updating = true
            this.updateError = null
            try {
                const response = await axios.put(`${this.baseUrl}/userEdit`, {
                    id: this.user.idUser,
                    username: this.editableUser.username,
                    displayName: this.editableUser.displayName,
                    email: this.editableUser.email,
                })
                // Update local user data and Vuex store
                this.user = response.data
                this.$store.commit('setCurrentUser', response.data)
                this.isEditing = false
            } catch (error) {
                console.error('Error updating profile:', error)
                this.updateError = error.response?.data?.error || 'Failed to update profile. Please try again.'
            } finally {
                this.updating = false
            }
        },
        removePost(postId) {
            this.userPosts = this.userPosts.filter(post => post.idPost !== postId);
        }
    },
    async created() {
        if (this.$store.state.userId == -1) {
            this.$router.push({ path: '/' });
        }
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
            this.postsError = 'Failed to load user posts'
        } finally {
            this.loadingPosts = false
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

.user-info, .edit-user-info {
    padding: 1em;
    background-color: #f9f9f9;
}

.user-info p {
    margin: 0.5em 0;
    color: #333;
}

.edit-user-info .form-group {
    margin-bottom: 1rem;
}

.edit-user-info label {
    display: block;
    margin-bottom: 0.5rem;
    color: #555;
}

.edit-user-info input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 4px;
}

.edit-user-info button {
    margin-top: 1rem;
    padding: 0.5rem 1rem;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1rem;
    margin-right: 0.5rem;
}

.edit-user-info button:disabled {
    background-color: #a5d6a7;
    cursor: not-allowed;
}

.edit-user-info .error {
    margin-top: 1rem;
    color: #d32f2f;
    text-align: center;
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

button {
    padding: 0.5rem 1rem;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

button:hover {
    background-color: #45a049;
}
</style>