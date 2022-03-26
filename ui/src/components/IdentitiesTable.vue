<template>
  <div class="identities-table">
    <h4>Last identities</h4>
    <b-table striped hover :items="identities" :fields="fields">
      <template #cell(name)="data">
        <router-link :to="{ path: '/identity/' + data.value }">{{ data.value }}</router-link>
      </template>
      <template #cell(blockheight)="data">
        <router-link :to="{ path: '/block/' + data.value }">{{ data.value }}</router-link>
      </template>
      <template #cell(identityaddress)="data">
        <router-link :to="{ path: '/address/' + data.value }">{{ data.value }}</router-link>
      </template>
      <template #cell(txid)="data">
        <router-link :to="{ path: '/transaction/' + data.value }">{{ data.value.substr(0, 4) }}...{{ data.value.substr(data.value.length - 4, 4) }}</router-link>
      </template>
    </b-table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {identities} from './mockData';
  import {apiURL, isMock} from '../config';

  export default {
    data() {
      return {
        fields: ['name', 'blockheight', 'identityaddress', 'txid'],
        identities: isMock ? identities : null,
      }
    },
    mounted () {
      axios
        .get(`${apiURL}/identities/last`)
        .then(response => (this.identities = response.data))
    }
  }
</script>

<style scoped>
  .identities-table {
    margin-top: 30px;
  }
</style>