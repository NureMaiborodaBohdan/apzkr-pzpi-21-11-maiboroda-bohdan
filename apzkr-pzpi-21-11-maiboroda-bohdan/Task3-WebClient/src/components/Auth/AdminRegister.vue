<template>
  <div>
    <h2>{{ $t('registerAdmin') }}</h2>
    <form @submit.prevent="handleRegisterAdmin" class="form-container">
      <div class="form-group">
        <label for="username">{{ $t('username') }}</label>
        <input type="text" v-model="username" required>
      </div>

      <div class="form-group">
        <label for="email">{{ $t('email') }}</label>
        <input type="email" v-model="email" required>
      </div>

      <div class="form-group">
        <label for="password">{{ $t('password') }}</label>
        <input type="password" v-model="password" required>
      </div>

      <div class="form-group">
        <label for="name">{{ $t('name') }}</label>
        <input type="text" v-model="name" required>
      </div>

      <div class="form-group">
        <label for="surname">{{ $t('surname') }}</label>
        <input type="text" v-model="surname" required>
      </div>

      <div class="form-group">
        <label for="patronymic">{{ $t('patronymic') }}</label>
        <input type="text" v-model="patronymic" required>
      </div>

      <div class="form-group">
        <label for="companyID">{{ $t('companyID') }}</label>
        <select v-model="companyID">
          <option v-for="company in companies" :key="company.CompanyID" :value="company.CompanyID">
            {{ company.Name }}
          </option>
        </select>
      </div>

      <div class="form-group">
        <label for="sex">{{ $t('sex') }}</label>
        <select v-model="sex">
          <option value="Male">{{ $t('male') }}</option>
          <option value="Female">{{ $t('female') }}</option>
        </select>
      </div>

      <button type="submit">{{ $t('registerAdmin') }}</button>
    </form>
  </div>
</template>


<script>
import { mapActions } from 'vuex';
import api from '../../api';

export default {
  name: 'RegisterAdmin',
  data() {
    return {
      username: '',
      email: '',
      password: '',
      role: 'Admin',
      name: '',
      surname: '',
      patronymic: '',
      companyID: null,
      sex: '',
      companies: []
    };
  },
  async created() {
    await this.fetchCompanies();
  },
  methods: {
    ...mapActions(['registerAdmin']),
    async fetchCompanies() {
      try {
        const response = await api.getCompanies();
        this.companies = response.data;
      } catch (error) {
        console.error('Error fetching companies:', error);
      }
    },
    async handleRegisterAdmin() {
      try {
        await this.registerAdmin({
          username: this.username,
          email: this.email,
          password: this.password,
          role: this.role,
          name: this.name,
          surname: this.surname,
          patronymic: this.patronymic,
          companyID: this.companyID,
          sex: this.sex
        });
        this.$router.push('/');
      } catch (err) {
        console.error('Failed to register admin:', err);
      }
    }
  }
};
</script>


<style scoped>
.form-container {
  display: flex;
  flex-direction: column;
  max-width: 400px;
  margin: auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 4px;
  background-color: #f9f9f9;
}

.form-group {
  margin-bottom: 10px;
}

label {
  font-weight: bold;
}

input[type="text"],
input[type="email"],
input[type="password"],
select {
  width: 100%;
  padding: 8px;
  margin-top: 4px;
  margin-bottom: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

button {
  width: 100%;
  background-color: #4CAF50;
  color: white;
  padding: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

button:hover {
  background-color: #45a049;
}

button:disabled {
  background-color: #cccccc;
  cursor: not-allowed;
}
</style>
