<template>
  <div>
    <h4>Last transactions</h4>
    <b-spinner v-if="transactions === null" type="grow"></b-spinner>
    <b-table striped hover :items="transactions" :fields="fields">
      <template #cell(blockHeight)="data">
        <router-link :to="{ path: '/block/' + data.value }">{{ data.value }}</router-link>
      </template>
      <template #cell(hash)="data">
        <router-link :to="{ path: '/transaction/' + data.value }">{{ data.value.substr(0, 4) }}...{{ data.value.substr(data.value.length - 4, 4) }}</router-link>
      </template>
    </b-table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {transactions} from './mockData';
  import {apiURL, isMock} from '../config';
  import {transformTransactions} from './transaction-helpers';

  export default {
    data() {
      return {
        fields: [
          'blockHeight',
          {
            key: 'timestampHumanReadable',
            label: 'Time',
          },
          'value',
          'type',
          'hash'
        ],
        transactions: isMock ? transactions : null,
      }
    },
    mounted () {
      axios
        .get(`${apiURL}/transactions/last`)
        .then(response => (this.transactions = transformTransactions(response.data)))
    }
  }
</script>