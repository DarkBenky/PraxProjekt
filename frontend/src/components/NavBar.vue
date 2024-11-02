<template>
  <nav class="navbar">
    <div class="nav-left">
      <router-link to="/"><h1 class="nav-title">MyApp</h1></router-link>
      <router-link to="add_post" class="nav-link">Add Post</router-link>
    </div>
    <router-link to="/user" class="nav-link">
      <div class="nav-right" @click="clicked">
        <UserProfile :user="user" :expanded="expanded" />
      </div>
    </router-link>
  </nav>
</template>

<script>
import UserProfile from './UserProfile.vue';

export default {
  name: 'NavBar',

  data() {
    return {
      expanded: false
    }
  },

  components: {
    UserProfile
  },
  props: {

    user: {
      type: Object,
      required: true
    }
  },

}
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: #333;
  color: #fff;
}

.nav-left {
  display: flex;
  align-items: center;
}

.nav-title {
  margin-right: 2rem;
  font-size: 1.5rem;
  font-weight: bold;
}

.nav-links {
  list-style: none;
  display: flex;
  gap: 1.5rem;
}

.nav-links a {
  color: #fff;
  text-decoration: none;
  font-size: 1rem;
}

.nav-links a:hover {
  text-decoration: underline;
}

.nav-right {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}
</style>