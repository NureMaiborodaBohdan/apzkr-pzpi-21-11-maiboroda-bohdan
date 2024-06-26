<template>
  <div>
    <h2>{{ $t('login') }}</h2>
    <form @submit.prevent="handleLogin">
      <label for="username">{{ $t('Username') }}</label>
      <input type="text" v-model="username" required>
      <label for="password">{{ $t('Password') }}</label>
      <input type="password" v-model="password" required>
      <button type="submit">{{ $t('login') }}</button>
    </form>
  </div>
</template>

<script>
import {mapActions} from 'vuex';

export default {
  name: 'UserLogin',
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    ...mapActions(['login']),
    async handleLogin() {
      try {
        await this.login({username: this.username, password: this.password});
        this.$router.push('/');
      } catch (err) {
        console.error('Failed to login:', err);
      }
    }
  }
};
</script>

<style scoped>
/* Add your styles here */
</style>
