<template>
  <div id="app">
    <nav class="navbar is-dark" role="navigation">
      <div class="navbar-brand">
        <div class="navbar-item title is-4">PartyPlayer</div>
      </div>
    </nav>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
  @import '../node_modules/bulma/css/bulma.css';
/*#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}*/
</style>
