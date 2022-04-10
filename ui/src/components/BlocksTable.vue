<template>
  <div>
    <h4>Last blocks</h4>
    <b-spinner v-if="blocks === null" type="grow"></b-spinner>
    <b-table striped hover :items="blocks" :fields="fields" class="blocks-table">
      <template #cell(height)="data">
        <router-link :to="{ path: '/block/' + data.value }">{{ data.value }}</router-link>
      </template>
      <template #cell(miner)="data">
        <router-link :to="{ path: '/address/' + data.value }">{{ data.value }}</router-link>
      </template>
    </b-table>
  </div>
</template>

<script>
  import axios from 'axios';
  import {blocks} from './mockData';
  import {apiURL, isMock} from '../config';
  import {secondsToString} from './time-helpers';

  export default {
    data() {
      return {
        fields: [
          'height',
          'miner',
          {
            key: 'timestampHumanReadable',
            label: 'Time',
          }
        ],
        blocks: isMock ? blocks : null,
      }
    },
    methods: {
      transformBlocks(blocks) {
        for (let i = 0; i < blocks.length; i++) {
          blocks[i].timestampHumanReadable = secondsToString(blocks[i].timestamp);
        }

        return blocks;
      },
      fetchData () {
        axios
        .get(`${apiURL}/blocks/last`)
        .then(response => (this.blocks = this.transformBlocks(response.data)));
      }
    },
    mounted () {
      this.fetchData();
    }
  }
</script>

<style scoped>
  .blocks-table tbody tr td:nth-child(2) {
    text-align: left;
    padding-left: 5%;
    padding-right: 5%;
  }
</style>
