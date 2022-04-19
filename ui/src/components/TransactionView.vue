<template>
  <div>
    <h4>Transaction Details</h4>
    <b-spinner v-if="transactionDetails === null" type="grow"></b-spinner>
    <table v-if="transactionDetails" class="table b-table table-striped table-hover text-left">
      <tbody>
        <tr>
          <td>
            Height
          </td>
          <td>
            <router-link :to="{path: '/block/' + transactionDetails.blockHeight}">{{transactionDetails.blockHeight}}</router-link>
          </td>
        </tr>
        <tr>
          <td>
            Block hash
          </td>
          <td>
            <router-link :to="{path: '/block/' + transactionDetails.blockHeight}">{{transactionDetails.blockHash}}</router-link>
          </td>
        </tr>
        <tr>
          <td>
            Time
          </td>
          <td>
            {{ transactionDetails.timestampHumanReadable }}
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
          <td colspan="2">
            <table class="transaction-movement-table">
              <thead>
                <th>Inputs</th>
                <th>Outputs</th>
              </thead>
              <tbody>
                <tr>
                  <td v-if="transactionDetails.movement !== null && transactionDetails.movement.input !== null && transactionDetails.movement.input.length > 0">
                    <div v-for="[address, value] in transactionDetails.movement.input" :key="address" style="padding-bottom: 10px">
                      <router-link :to="{ path: '/address/' + address }">{{address}}</router-link> <span style="padding-left: 10px">{{value}}</span>
                    </div>
                  </td>
                  <td v-if="transactionDetails.movement.input && !transactionDetails.movement.input.length">Coinbase</td>
                  <td v-if="transactionDetails.movement !== null && transactionDetails.movement.output !== null && transactionDetails.movement.output.length > 0">
                    <div v-for="[address, value] in transactionDetails.movement.output" :key="address" style="padding-bottom: 10px">
                      <router-link :to="{ path: '/address/' + address }">{{address}}</router-link> <span style="padding-left: 10px">{{value}}</span>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {apiURL} from '../config';
  import {transformTransactions} from './transaction-helpers';

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
          .get(`${apiURL}/transaction/${this.$route.params.hash}`)
          .then(response => (this.transactionDetails = transformTransactions([response.data])[0]));
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
  .transaction-movement-table {
    border-collapse: separate;
    width: 100%;
  }
  .transaction-movement-table td {
    width: 50%;
  }
  .transaction-movement-table td:last-child {
    white-space: nowrap;
  }
</style>