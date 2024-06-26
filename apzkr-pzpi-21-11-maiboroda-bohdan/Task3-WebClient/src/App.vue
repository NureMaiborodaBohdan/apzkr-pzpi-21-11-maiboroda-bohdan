<template>
  <div id="app">
    <nav>
      <router-link to="/">{{ $t('home') }}</router-link>
      <router-link to="/login">{{ $t('login') }}</router-link>
      <router-link to="/register">{{ $t('register') }}</router-link>
      <router-link to="/register-admin">{{ $t('registerAdmin') }}</router-link>
      <router-link to="/admin-dashboard">{{ $t('admin.dashboard') }}</router-link>

      <button @click="setLocale('en')">English</button>
      <button @click="setLocale('ua')">Українська</button>
    </nav>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  methods: {
    setLocale(locale) {
      this.$i18n.locale = locale;
    }
  }
};

</script>

<style>
nav {
  display: flex;
  gap: 10px;
}
</style>
