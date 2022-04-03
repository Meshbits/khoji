<template>
  <div>
    <h4 class="title">Address Details</h4>
    {{ address }}
    <h5 class="balance">Balance: {{ balance }} VRSCTEST</h5>
    <b-spinner v-if="balance === null || transactions === null" type="grow"></b-spinner>
    <div v-if="transactions">
      <h5>Transactions</h5>
      <table v-for="transactionDetails in transactions" :key="transactionDetails.hash" class="transacction-details-small margin-top table b-table table-striped table-hover text-left">
        <tbody>
          <tr>
            <td>
              Height
            </td>
            <td>
              <router-link :to="{ path: '/block/' + transactionDetails.blockHeight }">{{ transactionDetails.blockHeight }}</router-link>
            </td>
          </tr>
          <tr>
            <td>
              Block hash
            </td>
            <td>
              <router-link :to="{ path: '/block/' + transactionDetails.blockHeight }">{{ transactionDetails.blockHash }}</router-link>
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
          <!--tr>
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
          </tr-->
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import {fromSats} from './math-helpers';
  import {apiURL} from '../config';

  export default {
    computed: {
      address() {
        return this.$route.params.address;
      },
    },
    data() {
      return {
        balance: null,
        transactions: null,
      }
    },
    watch: {
      '$route': 'fetchData'
    },
    methods: {
      fetchData () {
        axios
          .get(`${apiURL}/balance/${this.$route.params.address}`)
          .then(response => (this.balance = fromSats(response.data.Balance)));
        axios
          .get(`${apiURL}/transactions/${this.$route.params.address}`)
          .then(response => (this.transactions = response.data));
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
  .margin-top {
    margin-bottom: 50px !important;
  }
  .transacction-details-small {
    max-width: 700px;
    margin: 0 auto;
  }
  .balance {
    margin-top: 20px;
    margin-bottom: 30px;
  }
  .title {
    padding-top: 10px;
    padding-bottom: 10px;
  }
  pre {
    word-break: break-all;
  }
</style>