<template>
  <div>
    <img alt="Vue logo" src="./assets/logo.png" width="50px" style="margin-bottom: 40px">
    <div class="container">
      <div class="row">
        <div class="col-sm">
          <BlocksTable />
        </div>
        <div class="col-sm offset-md-1">
          <TransactionsTable />
        </div>
      </div>
      <NetworkInfoTable />
    </div>
  </div>
</template>

<script>
import BlocksTable from './components/BlocksTable.vue'
import TransactionsTable from './components/TransactionsTable.vue'
import NetworkInfoTable from './components/NetworkInfoTable.vue'

export default {
  name: 'App',
  components: {
    BlocksTable,
    TransactionsTable,
    NetworkInfoTable
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
