<template>
  <div>
    <h4>Last blocks</h4>
    <b-table striped hover :items="blocks" :fields="fields">
      <template #cell(height)="data">
        <router-link :to="{ path: '/block/' + data.value }">{{ data.value }}</router-link>
      </template>
    </b-table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {blocks} from './mockData';

  export default {
    data() {
      return {
        fields: ['height', 'miner', 'timestamp'],
        blocks: blocks,
      }
    },
    mounted () {
      axios
        .get('http://localhost:3334/api/blocks/last')
        .then(response => (this.blocks = response.data))
    }
  }
</script>