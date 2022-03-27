<template>
  <div>
    <h4>Identity Info</h4>
    <table class="table b-table table-striped table-hover text-left">
      <tbody>
        <tr>
          <td>
            Name
          </td>
          <td>
            {{ identityDetails.name }}
          </td>
        </tr>
        <tr>
          <td>
            Height
          </td>
          <td>
            {{ identityDetails.blockheight }}
          </td>
        </tr>
        <tr>
          <td>
            Address
          </td>
          <td>
            {{ identityDetails.identityaddress }}
          </td>
        </tr>
        <tr>
          <td>
            Parent
          </td>
          <td>
            {{ identityDetails.parent }}
          </td>
        </tr>
        <tr>
          <td>
            Transaction ID
          </td>
          <td>
            <router-link :to="{ path: '/transaction/' + identityDetails.txid }">{{ identityDetails.txid }}</router-link>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {apiURL} from '../config';

  export default {
    computed: {
      name() {
        return this.$route.params.name;
      },
    },
    data() {
      return {
        identityDetails: {},
      }
    },
    watch: {
      '$route': 'fetchData'
    },
    methods: {
      fetchData () {
        axios
          .get(`${apiURL}/identity/${this.$route.params.name}`)
          .then(response => (this.identityDetails = response.data));
      }
    },
    mounted () {
      this.fetchData();
    },
  }
</script>

<style scoped>
  .text-left {
    text-align: left;
  }
</style>