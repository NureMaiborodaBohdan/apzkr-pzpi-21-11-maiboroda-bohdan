<template>
  <div>
    <h2>{{ $t('admin.users') }}</h2>
    <ul>
      <li v-for="user in users" :key="user.UserID">
        <b>{{ user.Username }}</b>: {{ user.Email }}, {{ user.Role }}, {{ user.Name }}, {{ user.Surname }}, {{ user.Patronymic }},
        <span v-if="user.Company">
          {{ user.Company.CompanyName }}
        </span>
        <span v-else>
          No Company
        </span>,
        {{ user.Sex }}
        <button @click="deleteUser(user.UserID)">Delete</button>
        <button @click="prepareUpdate(user)">Update</button>
      </li>
    </ul>

    <form v-if="isUpdating" @submit.prevent="updateUser">
      <h3>Update User</h3>
      <div>
        <label for="updateUsername">Username:</label>
        <input type="text" id="updateUsername" v-model="updateForm.Username" placeholder="Username" required>
      </div>

      <div>
        <label for="updateEmail">Email:</label>
        <input type="email" id="updateEmail" v-model="updateForm.Email" placeholder="Email" required>
      </div>

      <div>
        <label for="updateRole">Role:</label>
        <input type="text" id="updateRole" v-model="updateForm.Role" placeholder="Role">
      </div>

      <div>
        <label for="updateName">Name:</label>
        <input type="text" id="updateName" v-model="updateForm.Name" placeholder="Name">
      </div>

      <div>
        <label for="updateSurname">Surname:</label>
        <input type="text" id="updateSurname" v-model="updateForm.Surname" placeholder="Surname">
      </div>

      <div>
        <label for="updatePatronymic">Patronymic:</label>
        <input type="text" id="updatePatronymic" v-model="updateForm.Patronymic" placeholder="Patronymic">
      </div>

      <div>
        <label for="updateCompanyID">Company:</label>
        <select id="updateCompanyID" v-model="updateForm.CompanyID">
          <option value="">No Company</option>
          <option v-for="company in companies" :key="company.CompanyID" :value="company.CompanyID">
            {{ company.Name }}
          </option>
        </select>
      </div>

      <div>
        <label for="updateSex">Sex:</label>
        <input type="text" id="updateSex" v-model="updateForm.Sex" placeholder="Sex">
      </div>

      <button type="submit">Update</button>
    </form>

    <form v-else @submit.prevent="createUser">
      <div>
        <label for="newUsername">Username:</label>
        <input type="text" id="newUsername" v-model="newUser.Username" placeholder="Username" required>
      </div>

      <div>
        <label for="newEmail">Email:</label>
        <input type="email" id="newEmail" v-model="newUser.Email" placeholder="Email" required>
      </div>

      <div>
        <label for="newRole">Role:</label>
        <input type="text" id="newRole" v-model="newUser.Role" placeholder="Role">
      </div>

      <div>
        <label for="newName">Name:</label>
        <input type="text" id="newName" v-model="newUser.Name" placeholder="Name">
      </div>

      <div>
        <label for="newSurname">Surname:</label>
        <input type="text" id="newSurname" v-model="newUser.Surname" placeholder="Surname">
      </div>

      <div>
        <label for="newPatronymic">Patronymic:</label>
        <input type="text" id="newPatronymic" v-model="newUser.Patronymic" placeholder="Patronymic">
      </div>

      <div>
        <label for="newCompanyID">Company:</label>
        <select id="newCompanyID" v-model="newUser.CompanyID">
          <option value="">No Company</option>
          <option v-for="company in companies" :key="company.CompanyID" :value="company.CompanyID">
            {{ company.Name }}
          </option>
        </select>
      </div>

      <div>
        <label for="newSex">Sex:</label>
        <input type="text" id="newSex" v-model="newUser.Sex" placeholder="Sex">
      </div>

      <button type="submit">Create User</button>
    </form>

    <p v-if="errorMessage" style="color: red;">{{ errorMessage }}</p>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'UsersManagement',
  data() {
    return {
      users: [],
      companies: [],
      newUser: {
        Username: '',
        Email: '',
        Role: '',
        Name: '',
        Surname: '',
        Patronymic: '',
        CompanyID: '',
        Sex: ''
      },
      updateForm: {
        UserID: '',
        Username: '',
        Email: '',
        Role: '',
        Name: '',
        Surname: '',
        Patronymic: '',
        CompanyID: '',
        Sex: ''
      },
      isUpdating: false,
      errorMessage: ''
    };
  },
  async created() {
    try {
      await this.fetchUsers();
      await this.fetchCompanies();
    } catch (error) {
      this.errorMessage = 'Error fetching data. Please try again later.';
      console.error('Error fetching data:', error);
    }
  },
  methods: {
    async fetchUsers() {
      try {
        const response = await api.getUsers();
        this.users = response.data;
      } catch (error) {
        this.errorMessage = 'Error fetching users. Please try again later.';
        console.error('Error fetching users:', error);
      }
    },
    async fetchCompanies() {
      try {
        const response = await api.getCompanies();
        this.companies = response.data;
      } catch (error) {
        this.errorMessage = 'Error fetching companies. Please try again later.';
        console.error('Error fetching companies:', error);
      }
    },
    async createUser() {
      try {
        const payload = {
          Username: this.newUser.Username.trim(),
          Email: this.newUser.Email.trim(),
          Role: this.newUser.Role ? this.newUser.Role.trim() : '',
          Name: this.newUser.Name ? this.newUser.Name.trim() : '',
          Surname: this.newUser.Surname ? this.newUser.Surname.trim() : '',
          Patronymic: this.newUser.Patronymic ? this.newUser.Patronymic.trim() : '',
          CompanyID: parseInt(this.newUser.CompanyID) || null,
          Sex: this.newUser.Sex ? this.newUser.Sex.trim() : ''
        };

        if (!payload.Username || !payload.Email) {
          console.error('Username and Email are required fields.');
          return;
        }

        const response = await api.createUser(payload);
        console.log('Create user response:', response);

        this.newUser = {
          Username: '',
          Email: '',
          Role: '',
          Name: '',
          Surname: '',
          Patronymic: '',
          CompanyID: '',
          Sex: ''
        };

        this.errorMessage = '';
        await this.fetchUsers();
      } catch (error) {
        this.errorMessage = 'Error creating user. Please try again.';
        console.error('Error creating user:', error);
        if (error.response) {
          console.error('Error details:', error.response.data);
        } else {
          console.error('No response from server.');
        }
      }
    },
    async deleteUser(userID) {
      try {
        const response = await api.deleteUser(userID);
        console.log('Delete user response:', response);

        await this.fetchUsers();
      } catch (error) {
        this.errorMessage = 'Error deleting user. Please try again.';
        console.error('Error deleting user:', error);
      }
    },
    async prepareUpdate(user) {
      this.updateForm.UserID = user.UserID;
      this.updateForm.Username = user.Username;
      this.updateForm.Email = user.Email;
      this.updateForm.Role = user.Role;
      this.updateForm.Name = user.Name;
      this.updateForm.Surname = user.Surname;
      this.updateForm.Patronymic = user.Patronymic;
      this.updateForm.CompanyID = user.CompanyID;
      this.updateForm.Sex = user.Sex;

      this.isUpdating = true;
    },
    async updateUser() {
      try {
        const payload = {
          Username: this.updateForm.Username.trim(),
          Email: this.updateForm.Email.trim(),
          Role: this.updateForm.Role ? this.updateForm.Role.trim() : '',
          Name: this.updateForm.Name ? this.updateForm.Name.trim() : '',
          Surname: this.updateForm.Surname ? this.updateForm.Surname.trim() : '',
          Patronymic: this.updateForm.Patronymic ? this.updateForm.Patronymic.trim() : '',
          CompanyID: parseInt(this.updateForm.CompanyID) || null,
          Sex: this.updateForm.Sex ? this.updateForm.Sex.trim() : ''
        };

        const userID = this.updateForm.UserID;
        const response = await api.updateUser(userID, payload);
        console.log('Update user response:', response);

        this.updateForm = {
          UserID: '',
          Username: '',
          Email: '',
          Role: '',
          Name: '',
          Surname: '',
          Patronymic: '',
          CompanyID: '',
          Sex: ''
        };

        this.isUpdating = false;
        this.errorMessage = '';
        await this.fetchUsers();
      } catch (error) {
        this.errorMessage = 'Error updating user. Please try again.';
        console.error('Error updating user:', error);
        if (error.response) {
          console.error('Error details:', error.response.data);
        } else {
          console.error('No response from server.');
        }
      }
    }
  }
};
</script>

<style scoped>
/* Add your custom styles here */
</style>
