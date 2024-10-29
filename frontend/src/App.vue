<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',

  components: {
  },

  data() {
    return {
      url: "http://localhost:5050",
      Posts: [],
      activePostId: null,
      comments: [],
      loadingComments: false,
      commentsError: null,
      Users: [],
    }
  },

  // created() {
  //   this.GetAllPosts()
  //   this.GetAllUsers()
  // },

  methods: {

    getUserWithId(id) {
      return this.Users.find(user => user.idUser === id)
    },

    // async GetAllPosts() {
    //   axios.get(this.url + "/posts")
    //     .then(response => {
    //       this.Posts = response.data
    //     })
    //     .catch(error => {
    //       console.error('Error fetching posts:', error)
    //     })
    // },

    // async GetAllUsers() {
    //   axios.get(this.url + "/users")
    //     .then(response => {
    //       this.Users = response.data
    //     })
    // },

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
        const response = await axios.get(`${this.url}/comments`, {
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