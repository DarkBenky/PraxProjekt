<template>
  <div class="add-post">
    <h2>Add a Post</h2>
    <form @submit.prevent="submitPost">
      <label for="contentText">Content:</label>
      <textarea id="contentText" v-model="contentText" required></textarea>

      <button type="submit">Add Post</button>

      <!-- Display success or error messages -->
      <p v-if="message" :class="{ success: success, error: !success }">{{ message }}</p>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'AddPost',
  data() {
    return {
      contentText: '',     // Content of the post
      message: '',         // Success or error message
      success: false       // Indicates if the operation was successful
    }
  },
  methods: {
    async submitPost() {
      try {
        const response = await axios.post('http://localhost:5050/addPost', {
          userID: this.$store.state.userId,
          content_text: this.contentText
        }, {
          headers: { 'Content-Type': 'multipart/form-data' }
        })

        this.message = response.data.message
        this.success = true
        this.clearForm()

        //   forward to home page
        // this.$router.push('/')
      } catch (error) {
        console.error('Error adding post:', error)
        this.message = 'Failed to add post'
        this.success = false
      }
    },
    clearForm() {
      this.userID = ''
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
