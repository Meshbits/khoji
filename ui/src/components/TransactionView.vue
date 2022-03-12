<template>
  <div>
    <h4>Transaction Details</h4>
    <table class="table b-table table-striped table-hover text-left">
      <tbody>
        <tr>
          <td>
            Height
          </td>
          <td>
            {{ transactionDetails.blockHeight }}
          </td>
        </tr>
        <tr>
          <td>
            Block hash
          </td>
          <td>
            {{ transactionDetails.blockHash }}
          </td>
        </tr>
        <tr>
          <td>
            Time
          </td>
          <td>
            {{ transactionDetails.timestamp }}
          </td>
        </tr>
        <tr>
          <td>
            Type
          </td>
          <td>
            {{ transactionDetails.type }}
          </td>
        </tr>
        <tr>
          <td>
            Version
          </td>
          <td>
            {{ transactionDetails.version }}
          </td>
        </tr>
        <tr>
          <td>
            Inputs
          </td>
          <td>
            <pre>{{ JSON.stringify(transactionDetails.vin, null, 2) }}</pre>
          </td>
        </tr>
        <tr>
          <td>
            Outputs
          </td>
          <td>
            <pre>{{ JSON.stringify(transactionDetails.vout, null, 2) }}</pre>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  import axios from 'axios';
  //import {networkInfo} from './mockData';

  export default {
    computed: {
      hash() {
        return this.$route.params.hash;
      },
    },
    data() {
      return {
        transactionDetails: {},
      }
    },
    watch: {
      '$route': 'fetchData'
    },
    methods: {
      fetchData () {
        axios
          .get('http://localhost:3334/api/transaction/' + this.$route.params.hash)
          .then(response => (this.transactionDetails = response.data));
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