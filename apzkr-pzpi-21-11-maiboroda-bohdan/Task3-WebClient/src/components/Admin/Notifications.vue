<template>
  <div>
    <h2>Notifications</h2>
    <template v-if="notifications === null">
      <p>Loading...</p>
    </template>
    <template v-else-if="notifications.length > 0">
      <ul>
        <li v-for="notification in notifications" :key="notification.NotificationID">
          {{ notification.Message }}
        </li>
      </ul>
    </template>
    <template v-else>
      <p>No notifications found.</p>
    </template>
    <p v-if="error">{{ error }}</p>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'AdminNotifications',
  data() {
    return {
      notifications: null,
      error: null
    };
  },
  async created() {
    try {
      const response = await api.getAllNotifications();
      this.notifications = response.data;
    } catch (error) {
      console.error('Error fetching notifications:', error);
      this.error = 'Failed to fetch notifications. Please try again later.';
      this.notifications = [];
    }
  }
};
</script>

<style scoped>
/* Ваши стили здесь */
</style>
