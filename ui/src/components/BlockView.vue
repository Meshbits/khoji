<template>
  <div>
    <h4>Block Info</h4>
    <b-spinner v-if="blockInfo === null" type="grow"></b-spinner>
    <div v-if="blockInfo">
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
              Merkle root
            </td>
            <td>
              {{ blockInfo.merkleRoot }}
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
      <div v-if="blockInfo.transactions && blockInfo.transactions.length" class="block-transactions">
        <h4>Transactions</h4>
        <b-table striped hover :items="blockInfo.transactions" :fields="transactionFields" class="transactions-table">
          <template #cell(txid)="data">
            <router-link :to="{ path: '/transaction/' + data.value }">{{ data.value }}</router-link>
          </template>
        </b-table>
      </div>
    </div>
  </div>
</template>

<script>
  import axios from 'axios';
  import {apiURL} from '../config';

  export default {
    computed: {
      height() {
        return this.$route.params.height;
      },
    },
    data() {
      return {
        transactionFields: ['txid', 'value'],
        blockInfo: null,
      }
    },
    watch: {
      '$route': 'fetchData'
    },
    methods: {
      transformBlockTransactions(txs) {
        for (let i = 0; i < txs.length; i++) {
          let vinSum = 0, voutSum = 0;

          for (let j = 0; j < txs[i].vin.length; j++) {
            if (txs[i].vin[j].value) {
              vinSum += txs[i].vin[j].value;
            }
          }

          for (let j = 0; j < txs[i].vout.length; j++) {
            if (txs[i].vout[j].value) {
              voutSum += txs[i].vout[j].value;
            }
          }

          txs[i].vinSum = vinSum;
          txs[i].value = voutSum;
        }

        return txs;
      },
      fetchData () {
        axios
          .get(`${apiURL}/block/${this.$route.params.height}`)
          .then(response => {
            this.blockInfo = response.data;

            if (this.blockInfo.transactions) {
              this.blockInfo.transactions = this.transformBlockTransactions(response.data.transactions);
            }
          });
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
  .block-transactions {
    margin-top: 50px;
  }
</style>