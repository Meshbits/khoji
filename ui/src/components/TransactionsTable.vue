<template>
  <div>
    <h4>Last transactions</h4>
    <b-table striped hover :items="transactions" :fields="fields"></b-table>
  </div>
</template>

<script>
  //import axios from 'axios';
  import {transactions} from './mockData';

  export default {
    data() {
      return {
        fields: ['blockHeight', 'value', 'type'],
        transactions: transactions,
      }
    },
    // takes too long to fetch :(
    /*mounted () {
      axios
        .get('http://localhost:3334/api/transactions/last')
        .then(response => (this.transactions = response.data))
    }*/
  }
</script>