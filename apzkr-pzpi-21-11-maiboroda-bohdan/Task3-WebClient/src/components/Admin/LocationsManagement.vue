<template>
  <div>
    <h2>Locations Management</h2>
    <ul>
      <li v-for="location in locations" :key="location.LocationID">
        <b>{{ location.Country }}, {{ location.City }}</b>: {{ location.Address }}, {{ location.PostCode }}
        <button @click="deleteLocation(location.LocationID)">Delete</button>
        <button @click="prepareUpdate(location)">Update</button>
      </li>
    </ul>

    <!-- Form for creating a new location -->
    <form v-if="!isUpdatingLocation" @submit.prevent="createLocation">
      <h3>Create Location</h3>

      <div>
        <label for="country">Country:</label>
        <input type="text" id="country" v-model="newLocation.Country" placeholder="Country" required>
      </div>

      <div>
        <label for="city">City:</label>
        <input type="text" id="city" v-model="newLocation.City" placeholder="City" required>
      </div>

      <div>
        <label for="address">Address:</label>
        <input type="text" id="address" v-model="newLocation.Address" placeholder="Address" required>
      </div>

      <div>
        <label for="postCode">Post Code:</label>
        <input type="number" id="postCode" v-model="newLocation.PostCode" placeholder="Post Code" required>
      </div>

      <button type="submit">Create Location</button>
    </form>

    <!-- Form for updating an existing location -->
    <form v-if="isUpdatingLocation" @submit.prevent="updateLocation">
      <h3>Update Location</h3>

      <div>
        <label for="updateCountry">Country:</label>
        <input type="text" id="updateCountry" v-model="updateForm.Country" placeholder="Country" required>
      </div>

      <div>
        <label for="updateCity">City:</label>
        <input type="text" id="updateCity" v-model="updateForm.City" placeholder="City" required>
      </div>

      <div>
        <label for="updateAddress">Address:</label>
        <input type="text" id="updateAddress" v-model="updateForm.Address" placeholder="Address" required>
      </div>

      <div>
        <label for="updatePostCode">Post Code:</label>
        <input type="number" id="updatePostCode" v-model="updateForm.PostCode" placeholder="Post Code" required>
      </div>

      <button type="submit">Update Location</button>
    </form>
  </div>
</template>

<script>
import api from '../../api';

export default {
  name: 'LocationsManagement',
  data() {
    return {
      locations: [],
      newLocation: {
        Country: '',
        City: '',
        Address: '',
        PostCode: null
      },
      updateForm: {
        LocationID: '',
        Country: '',
        City: '',
        Address: '',
        PostCode: null
      },
      isUpdatingLocation: false
    };
  },
  async created() {
    try {
      await this.fetchLocations();
    } catch (error) {
      console.error('Error fetching locations:', error);
    }
  },
  methods: {
    async fetchLocations() {
      try {
        const response = await api.getLocations();
        this.locations = response.data;
      } catch (error) {
        console.error('Error fetching locations:', error);
      }
    },
    async createLocation() {
      try {
        const response = await api.createLocation({
          Country: this.newLocation.Country,
          City: this.newLocation.City,
          Address: this.newLocation.Address,
          PostCode: this.newLocation.PostCode
        });

        console.log('Create location response:', response);

        this.newLocation = {
          Country: '',
          City: '',
          Address: '',
          PostCode: null
        };

        await this.fetchLocations();
      } catch (error) {
        console.error('Error creating location:', error);
        if (error.response) {
          console.error('Error details:', error.response.data);
        } else {
          console.error('No response from server.');
        }
      }
    },
    async updateLocation() {
      try {
        const response = await api.updateLocation(this.updateForm.LocationID, {
          Country: this.updateForm.Country,
          City: this.updateForm.City,
          Address: this.updateForm.Address,
          PostCode: this.updateForm.PostCode
        });

        console.log('Update location response:', response);

        this.updateForm = {
          LocationID: '',
          Country: '',
          City: '',
          Address: '',
          PostCode: null
        };

        this.isUpdatingLocation = false;

        await this.fetchLocations();
      } catch (error) {
        console.error('Error updating location:', error);
        if (error.response) {
          console.error('Error details:', error.response.data);
        } else {
          console.error('No response from server.');
        }
      }
    },
    async deleteLocation(locationID) {
      try {
        const response = await api.deleteLocation(locationID);

        console.log('Delete location response:', response);

        await this.fetchLocations();
      } catch (error) {
        console.error('Error deleting location:', error);
      }
    },
    prepareUpdate(location) {
      this.updateForm = {
        LocationID: location.LocationID,
        Country: location.Country,
        City: location.City,
        Address: location.Address,
        PostCode: location.PostCode
      };

      this.isUpdatingLocation = true;
    }
  }
};
</script>

<style scoped>
/* Add your custom styles here */
</style>
