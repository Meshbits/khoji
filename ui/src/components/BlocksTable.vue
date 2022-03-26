<template>
  <div>
    <h4>Last blocks</h4>
    <b-table striped hover :items="blocks" :fields="fields" class="blocks-table">
      <template #cell(height)="data">
        <router-link :to="{ path: '/block/' + data.value }">{{ data.value }}</router-link>
      </template>
    </b-table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {blocks} from './mockData';
  import {apiURL, isMock} from '../config';

  export default {
    data() {
      return {
        fields: ['height', 'miner', 'timestamp'],
        blocks: isMock ? blocks : null,
      }
    },
    mounted () {
      axios
        .get(`${apiURL}/blocks/last`)
        .then(response => (this.blocks = response.data))
    }
  }
</script>

<style>
  .blocks-table tbody tr td:nth-child(2) {
    text-align: left;
    padding-left: 5%;
    padding-right: 5%;
  }
</style>
