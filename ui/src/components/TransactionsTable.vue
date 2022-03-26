<template>
  <div>
    <h4>Last transactions</h4>
    <b-table striped hover :items="transactions" :fields="fields"></b-table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {transactions} from './mockData';
  import {apiURL, isMock} from '../config';

  export default {
    data() {
      return {
        fields: ['blockHeight', 'value', 'type'],
        transactions: isMock ? transactions : null,
      }
    },
    mounted () {
      axios
        .get(`${apiURL}/transactions/last`)
        .then(response => (this.transactions = response.data))
    }
  }
</script>