<template>
  <div>
    <h2>Companies Management</h2>
    <ul>
      <li v-for="company in companies" :key="company.CompanyID">
        <b>{{ company.Name }}</b>: {{ company.Description }}, Legal Limit: {{ company.LegalLimit }}
        <br>
        Location: {{ getLocationInfo(company.LocationID) }}
        <br>
        <button @click="deleteCompany(company.CompanyID)">Delete</button>
        <button @click="prepareUpdate(company)">Update</button>
      </li>
    </ul>

    <form v-if="isUpdating" @submit.prevent="updateCompany">
      <h3>Update Company</h3>
      <div>
        <label for="updateName">Company Name:</label>
        <input type="text" id="updateName" v-model="updateForm.Name" placeholder="Company Name" required>
      </div>

      <div>
        <label for="updateDescription">Description:</label>
        <input type="text" id="updateDescription" v-model="updateForm.Description" placeholder="Description">
      </div>

      <div>
        <label for="updateLegalLimit">Legal Limit:</label>
        <input type="number" step="any" id="updateLegalLimit" v-model="updateForm.LegalLimit" placeholder="Legal Limit">
      </div>

      <div>
        <label for="updateLocationSelect">Location:</label>
        <select id="updateLocationSelect" v-model="updateForm.LocationID">
          <option value="">Select Location</option>
          <option v-for="location in locations" :key="location.LocationID" :value="location.LocationID">
            {{ getLocationInfo(location.LocationID) }}
          </option>
        </select>
      </div>

      <button type="submit">Update</button>
    </form>

    <form v-else @submit.prevent="createCompany">
      <div>
        <label for="companyName">Company Name:</label>
        <input type="text" id="companyName" v-model="newCompany.Name" placeholder="Company Name" required>
      </div>

      <div>
        <label for="companyDescription">Description:</label>
        <input type="text" id="companyDescription" v-model="newCompany.Description" placeholder="Description">
      </div>

      <div>
        <label for="legalLimit">Legal Limit:</label>
        <input type="number" step="any" id="legalLimit" v-model="newCompany.LegalLimit" placeholder="Legal Limit">
      </div>

      <div>
        <label for="locationSelect">Location:</label>
        <select id="locationSelect" v-model="newCompany.LocationID">
          <option value="">Select Location</option>
          <option v-for="location in locations" :key="location.LocationID" :value="location.LocationID">
            {{ getLocationInfo(location.LocationID) }}
          </option>
        </select>
      </div>

      <button type="submit">Create Company</button>
    </form>

    <p v-if="errorMessage" style="color: red;">{{ errorMessage }}</p>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'CompaniesManagement',
  data() {
    return {
      companies: [],
      locations: [],
      newCompany: {
        Name: '',
        Description: '',
        LegalLimit: null,
        LocationID: ''
      },
      updateForm: {
        CompanyID: '',
        Name: '',
        Description: '',
        LegalLimit: null,
        LocationID: ''
      },
      isUpdating: false,
      errorMessage: ''
    };
  },
  async created() {
    try {
      await this.fetchCompanies();
      await this.fetchLocations();
    } catch (error) {
      this.errorMessage = 'Error fetching data. Please try again later.';
      console.error('Error fetching initial data:', error);
    }
  },
  methods: {
    async fetchCompanies() {
      try {
        const response = await api.getCompanies();
        this.companies = response.data;
      } catch (error) {
        this.errorMessage = 'Error fetching companies. Please try again later.';
        console.error('Error fetching companies:', error);
      }
    },
    async fetchLocations() {
      try {
        const response = await api.getLocations();
        this.locations = response.data;
      } catch (error) {
        this.errorMessage = 'Error fetching locations. Please try again later.';
        console.error('Error fetching locations:', error);
      }
    },
    async createCompany() {
      try {
        if (this.newCompany.LegalLimit !== null && this.newCompany.LegalLimit.toString().trim() !== '') {
          this.newCompany.LegalLimit = parseFloat(this.newCompany.LegalLimit);
        } else {
          this.newCompany.LegalLimit = null;
        }

        const response = await api.createCompany(this.newCompany);
        console.log('Create company response:', response);

        this.newCompany = { Name: '', Description: '', LegalLimit: null, LocationID: '' };
        this.errorMessage = '';

        await this.fetchCompanies();
      } catch (error) {
        this.errorMessage = 'Error creating company. Please try again.';
        console.error('Error creating company:', error);
        console.error('Error details:', error.response ? error.response.data : 'No response from server.');
      }
    },

    async deleteCompany(companyID) {
      try {
        const response = await api.deleteCompany(companyID);
        console.log('Delete company response:', response);

        await this.fetchCompanies();
      } catch (error) {
        this.errorMessage = 'Error deleting company. Please try again.';
        console.error('Error deleting company:', error);
      }
    },
    async prepareUpdate(company) {
      this.updateForm.CompanyID = company.CompanyID;
      this.updateForm.Name = company.Name;
      this.updateForm.Description = company.Description;
      this.updateForm.LegalLimit = company.LegalLimit;
      this.updateForm.LocationID = company.LocationID;

      this.isUpdating = true;
    },
    async updateCompany() {
      try {
        const payload = {
          Name: this.updateForm.Name.trim(),
          Description: this.updateForm.Description.trim(),
          LegalLimit: parseFloat(this.updateForm.LegalLimit),
          LocationID: this.updateForm.LocationID
        };

        const companyId = this.updateForm.CompanyID;
        const response = await api.updateCompany(companyId, payload);
        console.log('Update company response:', response);

        this.updateForm = {
          CompanyID: '',
          Name: '',
          Description: '',
          LegalLimit: null,
          LocationID: ''
        };

        this.isUpdating = false;
        this.errorMessage = '';

        await this.fetchCompanies();
      } catch (error) {
        this.errorMessage = 'Error updating company. Please try again.';
        console.error('Error updating company:', error);
        console.error('Error details:', error.response ? error.response.data : 'No response from server.');
      }
    },
    getLocationInfo(locationID) {
      const location = this.locations.find(loc => loc.LocationID === locationID);
      if (location) {
        return `${location.Country}, ${location.City} - ${location.Address}, ${location.PostCode}`;
      }
      return 'Unknown Location';
    }
  }
};
</script>

<style scoped>
/* Add your custom styles here */
</style>
