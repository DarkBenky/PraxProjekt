<template>
  <div>
    <NavBar :user="user"></NavBar>
    <div class="add-post">
      <h2>Add a Post</h2>
      <form @submit.prevent="submitPost">
        <label for="contentText">Content:</label>
        <textarea id="contentText" v-model="contentText" required></textarea>
        
        <button type="submit">Add Post</button>
        
        <p v-if="message" :class="{ success: success, error: !success }">{{ message }}</p>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import NavBar from './NavBar.vue'

export default {
  name: 'AddPost',
  
  components: {
    NavBar
  },
  
  data() {
    return {
      user: {},
      contentText: '',
      message: '',
      success: false
    }
  },

  async created() {
    if (this.$store.state.userId == -1) {
      this.$router.push({path: '/'})
    }

  },
  
  methods: {
    async submitPost() {
      try {
        // Convert userId to a string
        const response = await axios.post('http://localhost:5050/addPost', {
          userID: String(this.$store.state.userId),  // Convert userId to string
          contentText: this.contentText
        })
        
        this.message = response.data.message
        this.success = true
        this.clearForm()

        // transefer to the post feed page
        this.$router.push({path: '/'})

      } catch (error) {
        console.error('Error adding post:', error)
        this.message = error.response?.data?.error || 'Failed to add post'
        this.success = false
      }
    },
    
    clearForm() {
      this.contentText = ''
    }
  }
}
</script>


<style scoped>
.add-post {
  max-width: 500px;
  margin: auto;
  padding: 1em;
  background-color: #f9f9f9;
  border-radius: 8px;
}

label {
  display: block;
  font-weight: bold;
  margin-top: 1em;
}

input,
textarea {
  width: 100%;
  padding: 0.5em;
  margin-top: 0.5em;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  margin-top: 1em;
  padding: 0.5em 1em;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}

.success {
  color: green;
}

.error {
  color: red;
}
</style>
