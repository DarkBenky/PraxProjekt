<template>
  <nav class="navbar">
    <div class="nav-left">
      <router-link to="/" class="nav-title">
        MyApp
      </router-link>
      <router-link to="/add_post" class="nav-link">
        Add Post
      </router-link>
    </div>
    <div class="nav-right">
      <router-link
        v-if="this.$store.state.userId != -1"
        to="/user"
        class="nav-link"
        @click="clicked"
      >
        <UserProfile :user="user" :expanded="expanded" />
      </router-link>
      <router-link v-else to="/login" class="nav-link">
        Login
      </router-link>
      <div class="nav-link" v-if="this.$store.state.userId != -1">
        <h3  @click="logout">Logout</h3>
      </div>
    </div>
  </nav>
</template>

<script>
import UserProfile from "./UserProfile.vue";

export default {
  name: "NavBar",

  data() {
    return {
      expanded: false,
    };
  },

  components: {
    UserProfile,
  },
  props: {
    user: {
      type: Object,
      required: true,
    },
  },

  methods: {
    clicked() {
      this.expanded = !this.expanded;
    },
    logout() {
      this.$store.commit("logout");
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: #222; /* Dark background for contrast */
  color: #fff;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2); /* Subtle shadow for depth */
}

.nav-left,
.nav-right {
  display: flex;
  align-items: center;
  gap: 1rem; /* Spacing between items */
}

.nav-title {
  font-size: 1.8rem;
  font-weight: bold;
  color: #fff;
  text-decoration: none;
  transition: color 0.3s ease;
}

.nav-title:hover {
  color: #ff9800; /* Highlight effect on hover */
}

.nav-link {
  font-size: 1rem;
  color: #fff;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.nav-link:hover {
  background-color: #444; /* Subtle hover background */
  color: #ff9800; /* Highlight effect on hover */
}

.nav-right {
  display: flex;
  align-items: center;
}

.nav-right .nav-link {
  display: flex;
  align-items: center;
}

.nav-right .UserProfile {
  display: flex;
  align-items: center;
}

@media (max-width: 768px) {
  .navbar {
    flex-direction: column;
    align-items: flex-start;
  }

  .nav-left,
  .nav-right {
    width: 100%;
    justify-content: space-between;
  }

  .nav-link {
    padding: 0.5rem 0.8rem;
    font-size: 0.9rem;
  }
}
</style>
