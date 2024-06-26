<template>
  <div>
    <h2>{{ $t('user.accessControls') }}</h2>
    <ul>
      <li v-for="control in accessControls" :key="control.id">
        {{ control.name }}: {{ control.status }}
      </li>
    </ul>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'UserAccessControls',
  data() {
    return {
      accessControls: [],
      errorMessage: ''
    };
  },
  async created() {
    try {
      const response = await api.getUserAccessControls();
      this.accessControls = response.data;
    } catch (error) {
      this.errorMessage = 'Failed to fetch user access controls. Please try again later.';
      console.error('Error fetching data:', error);
    }
  }
};
</script>


<style scoped>
/* Add your styles here */
</style>
