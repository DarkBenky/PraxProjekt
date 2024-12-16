<template>
  <div class="edit-post">
    <!-- <NavBar :user="getUserWithId($store.state.userId)"></NavBar> -->
    <div v-if="loading">Loading...</div>
    <div v-else-if="error">{{ error }}</div>
    <div v-else class="edit-form">
      <h2>Edit Post</h2>
      <textarea
        v-model="postContent"
        rows="5"
        placeholder="Edit your post..."
      ></textarea>
      <div class="buttons">
        <button @click="savePost">Save Changes</button>
        <button @click="cancel">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
// import NavBar from "./NavBar.vue";

export default {
  name: "EditPost",
  components: {
    // NavBar,
  },

  data() {
    return {
      postContent: "",
      loading: true,
      error: null,
      users: [],
      baseUrl: "http://localhost:5050",
    };
  },

  async created() {
    if (this.$store.state.userId == -1) {
      this.$router.push({ path: "/" });
    }
    try {
      // Get post ID from route params
      const postId = this.$route.params.id;
      console.log(postId);
      // Fetch post data
      const response = await axios.get(`${this.baseUrl}/post`, {
        params: {
          id: postId,
        },
      });
      this.postContent = response.data.content_text;
      await this.fetchUsers();
    } catch (error) {
      console.error("Error fetching post data:", error);
      this.error = "Failed to load post";
    } finally {
      this.loading = false;
    }
  },

  methods: {
    async fetchUsers() {
      try {
        const response = await axios.get(`${this.baseUrl}/users`);
        this.users = response.data;
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    },

    getUserWithId(id) {
      return this.users.find((user) => user.idUser === id);
    },

    async savePost() {
      try {
        const response = await axios.put(`${this.baseUrl}/editPost`, {
          postID: String(this.$route.params.id),
          contentText: this.postContent,
        });

        if (response.status === 200) {
          this.$router.push("/");
        }
      } catch (error) {
        console.error("Error updating post:", error);
        this.error = "Failed to update post";
      }
    },

    cancel() {
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.edit-post {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.edit-form {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

textarea {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ddd;
  border-radius: 4px;
  resize: vertical;
}

.buttons {
  display: flex;
  gap: 10px;
  justify-content: flex-end;
  margin-top: 20px;
}

button {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:first-child {
  background-color: #4caf50;
  color: white;
}

button:last-child {
  background-color: #f44336;
  color: white;
}
</style>
