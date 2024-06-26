<template>
  <div>
    <h2>Test Results</h2>

    <template v-if="testResults === null">
      <p>Loading...</p>
    </template>

    <template v-else-if="testResults.length > 0">
      <ul>
        <li v-for="result in testResults" :key="result.TestID">
          <b>Username: {{ result.Username }}</b>: {{ result.AlcoholLevel }}, Is Drunk: {{ result.IsDrunk }}, Time: {{ result.TestTime }}
          <button @click="deleteTestResult(result.TestID)">Delete</button>
        </li>
      </ul>
    </template>

    <template v-else>
      <p>No test results found.</p>
    </template>

    <p v-if="error">{{ error }}</p>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'AdminTestResults',
  data() {
    return {
      testResults: null,
      error: null
    };
  },
  async created() {
    await this.fetchTestResults();
  },
  methods: {
    async fetchTestResults() {
      try {
        const response = await api.getTestResults();
        this.testResults = response.data;
      } catch (error) {
        console.error('Error fetching test results:', error);
        this.error = 'Failed to fetch test results. Please try again later.';
        this.testResults = [];
      }
    },
    async deleteTestResult(testID) {
      try {
        const response = await api.deleteTestResult(testID);
        console.log('Delete test result response:', response);

        this.testResults = this.testResults.filter(result => result.TestID !== testID);
      } catch (error) {
        console.error('Error deleting test result:', error);
        if (error.response) {
          console.error('Error details:', error.response.data);
          this.error = 'Failed to delete test result: ' + error.response.data.message;
        } else {
          console.error('No response from server.');
          this.error = 'Failed to delete test result. Please try again later.';
        }
      }
    }
  }
};
</script>

<style scoped>
/* Add your custom styles here */
</style>
