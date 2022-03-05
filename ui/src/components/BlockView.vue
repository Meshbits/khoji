<template>
  <div>
    <h4>Block Info</h4>
    <table class="table b-table table-striped table-hover text-left">
      <tbody>
        <tr>
          <td>
            Height
          </td>
          <td>
            {{ blockInfo.height }}
          </td>
        </tr>
        <tr>
          <td>
            Hash
          </td>
          <td>
            {{ blockInfo.hash }}
          </td>
        </tr>
        <tr>
          <td>
            Previous block
          </td>
          <td>
            {{ blockInfo.previousBlock }}
          </td>
        </tr>
        <tr>
          <td>
            Next block
          </td>
          <td>
            {{ blockInfo.nextBlock }}
          </td>
        </tr>
        <tr>
          <td>
            Chainwork
          </td>
          <td>
            {{ blockInfo.chainWork }}
          </td>
        </tr>
        <tr>
          <td>
            Difficulty
          </td>
          <td>
            {{ blockInfo.difficulty }}
          </td>
        </tr>
      </tbody>
    </table>
    <div>
      <router-link :to="{ path: '/block/' + (Number(this.height) - 1) }" class="padding-right">Prev block</router-link>
      <router-link :to="{ path: '/block/' + (Number(this.height) + 1) }">Next block</router-link>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  //import {networkInfo} from './mockData';

  export default {
    computed: {
      height() {
        return this.$route.params.height;
      },
    },
    data() {
      return {
        blockInfo: {},
      }
    },
    watch: {
      '$route': 'fetchData'
    },
    methods: {
      fetchData () {
        axios
          .get('http://localhost:3334/api/block/' + this.$route.params.height)
          .then(response => (this.blockInfo = response.data));
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
  .padding-right {
    padding-right: 50px;
  }
</style>