<template>
  <div>
    <img alt="Vue logo" src="./assets/logo.png" width="50px" style="margin-bottom: 40px">
    <div class="container">
      <NetworkInfoTable />
    </div>
  </div>
</template>

<script>
import NetworkInfoTable from './components/NetworkInfoTable.vue'

export default {
  name: 'App',
  components: {
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
