<template>
  <div>
    <h1>{{ $t('user.dashboard') }}</h1>
    <nav>
      <router-link :to="{ name: 'UserTestResults' }">{{ $t('user.testResults') }}</router-link>
      <router-link :to="{ name: 'UserNotifications' }">{{ $t('user.notifications') }}</router-link>
      <router-link :to="{ name: 'UserAccessControls' }">{{ $t('user.accessControls') }}</router-link>
    </nav>
    <button @click="handleLogout">{{ $t('logout') }}</button>
    <router-view></router-view>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'UserDashboard',
  computed: {
    ...mapGetters(['isLoggedIn', 'currentUser']),
  },
  methods: {
    ...mapActions(['logout']),
    handleLogout() {
      this.logout().then(() => {
        this.$router.push({ name: 'UserLogin' });
      });
    },
  },
};
</script>

<style scoped>
/* Add your styles here */
nav {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

button {
  margin-top: 20px;
}
</style>
