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
    </div>
</template>

<script>
import axios from 'axios'
import NavBar from './NavBar.vue';

export default {
    name: 'UserProfile',

    components: {
        NavBar
    },
    data() {
        return {
            user: null,
            loading: true,
            error: null
        }
    },

    async created() {
        try {
            // Fetch user data from the API
            const response = await axios.get(`http://localhost:5050/user`, {
                params: { id: this.$store.state.userId }
            })
            this.user = response.data
        } catch (error) {
            console.error('Error fetching user data:', error)
            this.error = 'Failed to load user data'
        } finally {
            this.loading = false
        }
    }
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
