<template>
    <div class="login-container">
      <h2>Login</h2>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">Username or Email</label>
          <input
            type="text"
            id="username"
            v-model="username"
            required
            placeholder="Enter your username or email"
          />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            required
            placeholder="Enter your password"
          />
        </div>
        <button type="submit" :disabled="loading">
          {{ loading ? 'Logging in...' : 'Login' }}
        </button>
        <p v-if="error" class="error">{{ error }}</p>
      </form>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    name: 'LoginPage',
    data() {
      return {
        username: '',
        password: '',
        loading: false,
        error: null,
        baseUrl: 'http://localhost:5050'
      }
    },
    methods: {
      async handleLogin() {
        this.loading = true
        this.error = null
        try {
          const response = await axios.post(this.baseUrl+'/login', {
            username: this.username,
            password: this.password,
          })
          console.log(response.data)
          this.$store.commit('setUserId', response.data.idUser)
          this.$store.commit('setCurrentUser', response.data)

          // print stored data
          console.log(this.$store.state.userId)
          console.log(this.$store.state.currentUser)

          this.$router.push('/')
        } catch (err) {
          this.error = err.response?.data?.error || 'Login failed. Please try again.'
        } finally {
          this.loading = false
        }
      },
    },
  }
  </script>
  
  <style scoped>
  .login-container {
    max-width: 400px;
    margin: 2rem auto;
    padding: 2rem;
    background-color: #f9f9f9;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  }
  
  h2 {
    text-align: center;
    margin-bottom: 1.5rem;
    color: #333;
  }
  
  .form-group {
    margin-bottom: 1rem;
  }
  
  label {
    display: block;
    margin-bottom: 0.5rem;
    color: #555;
  }
  
  input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 4px;
  }
  
  button {
    width: 100%;
    padding: 0.7rem;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1rem;
  }
  
  button:disabled {
    background-color: #a5d6a7;
    cursor: not-allowed;
  }
  
  .error {
    margin-top: 1rem;
    color: #d32f2f;
    text-align: center;
  }
  </style>