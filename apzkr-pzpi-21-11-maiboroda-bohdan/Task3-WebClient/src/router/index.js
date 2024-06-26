import { createRouter, createWebHistory } from 'vue-router';
import store from '../store';
import HomePage from '../views/HomePage.vue';
import UserDashboard from '../views/UserDashboard.vue';
import AdminDashboard from '../views/AdminDashboard.vue';
import UserLogin from '../components/Auth/UserLogin.vue';
import UserRegister from '../components/Auth/UserRegister.vue';
import AdminRegister from '../components/Auth/AdminRegister.vue';
import CompaniesManagement from '../components/Admin/CompaniesManagement.vue';
import LocationsManagement from '../components/Admin/LocationsManagement.vue';
import UsersManagement from '../components/Admin/UsersManagement.vue';
import AdminTestResults from '../components/Admin/TestResults.vue';
import AdminNotifications from '../components/Admin/Notifications.vue';
import DataManagement from '../components/Admin/DataManagement.vue';
import UserTestResults from '../components/User/TestResults.vue';
import UserNotifications from '../components/User/Notifications.vue';
import UserAccessControl from '../components/User/AccessControl.vue';

const routes = [
    {
        path: '/',
        name: 'HomePage',
        component: HomePage,
        meta: { requiresGuest: true }
    },
    {
        path: '/login',
        name: 'UserLogin',
        component: UserLogin,
        meta: { requiresGuest: true }
    },
    {
        path: '/register',
        name: 'UserRegister',
        component: UserRegister,
        meta: { requiresGuest: true }
    },
    {
        path: '/register-admin',
        name: 'AdminRegister',
        component: AdminRegister,
        meta: { requiresGuest: true }
    },
    {
        path: '/user-dashboard',
        name: 'UserDashboard',
        component: UserDashboard,
        meta: { requiresAuth: true, requiresRole: 'User' },
        children: [
            {
                path: 'test-results',
                name: 'UserTestResults',
                component: UserTestResults,
                props: true
            },
            {
                path: 'notifications',
                name: 'UserNotifications',
                component: UserNotifications
            },
            {
                path: 'access-controls',
                name: 'UserAccessControls',
                component: UserAccessControl
            }
        ]
    },
    {
        path: '/admin-dashboard',
        name: 'AdminDashboard',
        component: AdminDashboard,
        meta: { requiresAuth: true, requiresRole: 'Admin' },
        children: [
            {
                path: 'companies',
                name: 'CompaniesManagement',
                component: CompaniesManagement
            },
            {
                path: 'locations',
                name: 'LocationsManagement',
                component: LocationsManagement
            },
            {
                path: 'users',
                name: 'UsersManagement',
                component: UsersManagement
            },
            {
                path: 'test-results',
                name: 'AdminTestResults',
                component: AdminTestResults
            },
            {
                path: 'notifications',
                name: 'AdminNotifications',
                component: AdminNotifications
            },
            {
                path: 'data-management',
                name: 'DataManagement',
                component: DataManagement
            }
        ]
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

router.beforeEach(async (to, from, next) => {
    const isLoggedIn = store.getters.isLoggedIn;

    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (!isLoggedIn) {
            next('/login');
        } else {
            try {
                const token = localStorage.getItem('token');
                const payload = parseJwt(token);
                const userRole = payload.role;

                if (to.meta.requiresRole && to.meta.requiresRole !== userRole) {
                    if (userRole === 'Admin' && to.path !== '/admin-dashboard') {
                        next('/admin-dashboard');
                    } else if (userRole === 'User' && to.path !== '/user-dashboard') {
                        next('/user-dashboard');
                    } else {
                        next();
                    }
                } else {
                    next();
                }
            } catch (error) {
                console.error('Error decoding token:', error);
                next('/login'); // Redirect to login page on error
            }
        }
    } else if (to.matched.some(record => record.meta.requiresGuest)) {
        if (isLoggedIn) {
            next();
        } else {
            next();
        }
    } else {
        next();
    }
});

function parseJwt(token) {
    try {
        return JSON.parse(atob(token.split('.')[1]));
    } catch (e) {
        return null;
    }
}

export default router;
