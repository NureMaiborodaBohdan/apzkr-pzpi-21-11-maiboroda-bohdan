import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://localhost:8000',
    headers: {
        'Content-Type': 'application/json',
    },
    withCredentials: true,
});

apiClient.interceptors.request.use(config => {
    const token = localStorage.getItem('token');
    console.log('Intercepting request, token:', token);
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    console.log('Config headers:', config.headers);
    return config;
}, error => {
    return Promise.reject(error);
});

apiClient.interceptors.response.use(response => {
    console.log('Response headers:', response.headers);
    return response;
}, error => {
    if (error.response) {
        console.error('Response error:', error.response);
    } else {
        console.error('Error without response:', error);
    }
    return Promise.reject(error);
});
export default {
    // Auth endpoints
    register(user) {
        return apiClient.post('/auth/register', user);
    },
    registerAdmin(admin) {
        return apiClient.post('/auth/registerAdmin', admin);
    },
    login(credentials) {
        return apiClient.post('/auth/login', credentials);
    },
    logout(credentials){
        return apiClient.post('/auth/logout', credentials);
    },
    // User endpoints
    getUserTestResults(userID) {
        return apiClient.get(`/api/user/${userID}/testresults`);
    },
    createUserTestResult(userID, testResult) {
        return apiClient.post(`/api/user/${userID}/testresults`, testResult);
    },
    getUserNotifications(userID) {
        return apiClient.get(`/api/user/${userID}/notifications`);
    },
    getUserAccessControls(userID) {
        return apiClient.get(`/api/user/${userID}/accesscontrols`);
    },

    // Admin endpoints
    getCompanies() {
        return apiClient.get('/api/admin/companies');
    },
    getCompanyByID(companyID) {
        return apiClient.get(`/api/admin/companies/${companyID}`);
    },
    createCompany(company) {
        return apiClient.post('/api/admin/companies', company);
    },
    updateCompany(companyID, company) {
        return apiClient.put(`/api/admin/companies/${companyID}`, company);
    },
    deleteCompany(companyID) {
        return apiClient.delete(`/api/admin/companies/${companyID}`);
    },

    getLocations() {
        return apiClient.get('/api/admin/locations');
    },
    getLocationByID(locationID) {
        return apiClient.get(`/api/admin/locations/${locationID}`);
    },
    createLocation(location) {
        return apiClient.post('/api/admin/locations', location);
    },
    updateLocation(locationID, location) {
        return apiClient.put(`/api/admin/locations/${locationID}`, location);
    },
    deleteLocation(locationID) {
        return apiClient.delete(`/api/admin/locations/${locationID}`);
    },

    getUsers() {
        return apiClient.get('/api/admin/users/');
    },
    createUser(user) {
        return apiClient.post('/api/admin/users/', user);
    },
    getUserByID(userID) {
        return apiClient.get(`/api/admin/users/${userID}`);
    },
    updateUser(userID, user) {
        return apiClient.put(`/api/admin/users/${userID}`, user);
    },
    deleteUser(userID) {
        return apiClient.delete(`/api/admin/users/${userID}`);
    },

    getTestResults() {
        return apiClient.get('/api/admin/testresults');
    },
    getTestResultByID(testID) {
        return apiClient.get(`/api/admin/testresults/${testID}`);
    },
    deleteTestResult(testID) {
        return apiClient.delete(`/api/admin/testresults/${testID}`);
    },

    getAllNotifications() {
        return apiClient.get('/api/admin/notifications');
    },
    getNotificationByID(notificationID) {
        return apiClient.get(`/api/admin/notifications/${notificationID}`);
    },
    createNotification(notification) {
        return apiClient.post('/api/admin/notifications', notification);
    },
    deleteNotification(notificationID) {
        return apiClient.delete(`/api/admin/notifications/${notificationID}`);
    },

    // Data management endpoints
    backupData() {
        return apiClient.post('/api/admin/data/backup');
    },
    restoreData() {
        return apiClient.post('/api/admin/data/restore');
    },
    exportData() {
        return apiClient.get('/api/admin/data/export');
    },
    importData(data) {
        return apiClient.post('/api/admin/data/import', data);
    },
};
