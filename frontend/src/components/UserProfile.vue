<template>
    <div :class="['user-profile', { 'compact': compact }]">
      <div v-if="loading" class="loading">
        Loading user...
      </div>
      <div v-else-if="error" class="error">
        {{ error }}
      </div>
      <div v-else class="profile-content">
        <div class="avatar">
          {{ initials }}
        </div>
        <div class="user-info">
          <div class="display-name">{{ user.displayName }}</div>
          <div v-if="!compact" class="username">@{{ user.username }}</div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    name: 'UserProfile',
    
    props: {
      userId: {
        type: Number,
        required: true
      },
      compact: {
        type: Boolean,
        default: false
      }
    },
    
    data() {
      return {
        url: "http://localhost:5050",
        user: null,
        loading: true,
        error: null
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
      userId: {
        immediate: true,
        handler: 'fetchUser'
      }
    },
    
    methods: {
      async fetchUser() {
        this.loading = true
        this.error = null
        
        try {
          const response = await axios.get(`${this.url}/users`, {
            params: {
              id: this.userId
            }
          })
          this.user = response.data
        } catch (error) {
          this.error = 'Failed to load user'
          console.error('Error fetching user:', error)
        } finally {
          this.loading = false
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
  
  .loading, .error {
    font-size: 0.9em;
    color: #666;
  }
  
  .error {
    color: #d32f2f;
  }
  </style>