<template>
  <div>
    <h2>Data Management</h2>
    <button @click="backupData">Backup Data</button>
    <button @click="restoreData">Restore Data</button>
    <button @click="exportData">Export Data</button>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'DataManagement',
  methods: {
    async backupData() {
      try {
        await api.backupData();
        alert('Backup successful');
      } catch (error) {
        console.error('Error backing up data:', error);
      }
    },
    async restoreData() {
      try {
        await api.restoreData();
        alert('Restore successful');
      } catch (error) {
        console.error('Error restoring data:', error);
      }
    },
    async exportData() {
      try {
        await api.exportData();
        alert('Export successful');
      } catch (error) {
        console.error('Error exporting data:', error);
      }
    },
  }
};
</script>

<style scoped>
button {
  margin-right: 10px;
}
</style>
