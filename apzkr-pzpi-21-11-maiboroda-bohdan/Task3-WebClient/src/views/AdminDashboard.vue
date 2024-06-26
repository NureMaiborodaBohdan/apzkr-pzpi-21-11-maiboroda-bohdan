<template>
  <div>
    <h1>{{ $t('admin.dashboard') }}</h1>
    <nav>
      <router-link :to="{ name: 'CompaniesManagement' }">{{ $t('admin.companies') }}</router-link>
      <router-link :to="{ name: 'LocationsManagement' }">{{ $t('admin.locations') }}</router-link>
      <router-link :to="{ name: 'UsersManagement' }">{{ $t('admin.users') }}</router-link>
      <router-link :to="{ name: 'AdminTestResults' }">{{ $t('admin.testResults') }}</router-link>
      <router-link :to="{ name: 'AdminNotifications' }">{{ $t('admin.notifications') }}</router-link>
      <router-link :to="{ name: 'DataManagement' }">{{ $t('admin.dataManagement') }}</router-link>
    </nav>
    <button @click="handleLogout">{{ $t('logout') }}</button>
    <router-view></router-view>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'AdminDashboard',
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
