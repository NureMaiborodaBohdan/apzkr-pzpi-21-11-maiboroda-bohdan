<template>
  <div>
    <h2>{{ $t('user.testResults') }}</h2>
    <ul>
      <li v-for="result in testResults" :key="result.id">
        {{ result.testName }}: {{ result.score }}
      </li>
    </ul>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'UserTestResults',
  props: {
    userID: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      testResults: []
    };
  },
  async created() {
    try {
      const response = await api.getUserTestResults(this.userID);
      this.testResults = response.data;
    } catch (error) {
      console.error('Помилка отримання даних:', error);
    }
  }
};
</script>


<style scoped>
/* Ваши стили */
</style>
