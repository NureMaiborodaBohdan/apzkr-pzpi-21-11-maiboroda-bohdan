import { createStore } from 'vuex';
import api from '../api';

export default createStore({
    state: {
        user: null,
        token: localStorage.getItem('token') || '',
        status: '',
    },
    mutations: {
        auth_request(state) {
            state.status = 'loading';
        },
        auth_success(state, { token, user }) {
            state.status = 'success';
            state.token = token;
            state.user = user;
        },
        auth_error(state, err) {
            state.status = 'error';
            console.error('Authentication error:', err);
        },
        logout(state) {
            state.status = '';
            state.token = '';
            state.user = null;
        },
    },
    actions: {
        async login({ commit }, user) {
            commit('auth_request');
            let headers = {};
            const token = localStorage.getItem('token');
            if (token) {
                headers['Authorization'] = `Bearer ${token}`;
            }
            try {
                const resp = await api.login(user, headers);
                if (resp && resp.data && resp.data.token) {
                    const { token, user } = resp.data;
                    localStorage.setItem('token', token);
                    commit('auth_success', { token, user });
                    return resp;
                } else {
                    throw new Error('Invalid response structure');
                }
            } catch (err) {
                commit('auth_error', err);
                localStorage.removeItem('token');
                throw err;
            }
        },
        async register({ commit }, user) {
            commit('auth_request');
            try {
                const resp = await api.register(user);
                if (resp && resp.data && resp.data.token) {
                    const { token, user } = resp.data;
                    localStorage.setItem('token', token);
                    commit('auth_success', { token, user });
                    return resp;
                } else {
                    throw new Error('Invalid response structure');
                }
            } catch (err) {
                commit('auth_error', err);
                localStorage.removeItem('token');
                throw err;
            }
        },
        logout({ commit }) {
            commit('logout');
            localStorage.removeItem('token');
            if (api.defaults && api.defaults.headers.common['Authorization']) {
                delete api.defaults.headers.common['Authorization'];
            } else {
                console.error('Authorization header not found in api.defaults');
            }
        },
        async createCompany({ commit }, company) {
            commit('create_company_request');
            try {
                const response = await api.createCompany(company);
                const createdCompany = response.data;
                commit('create_company_success', createdCompany);
                return response;
            } catch (error) {
                commit('create_company_error', error);
                throw error;
            }
        },
    },
    getters: {
        isLoggedIn: state => !!state.token,
        authStatus: state => state.status,
        currentUser: state => state.user,
        isAdmin: state => {
            if (!state.user) {
                return false;
            }
            return state.user.role && state.user.role.toLowerCase() === 'admin';
        },
        userID: state => {
            return state.user ? state.user.id : localStorage.getItem('userID');
        }
    },
});
