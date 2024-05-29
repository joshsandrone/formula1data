<script>
import HeaderBar from './components/Generic/HeaderBar.vue'
// import SeasonDash from './components/SeasonStatsDashboard.vue';


export default {
  components: {
    HeaderBar,
    // SeasonDash
  }
};
</script>




<template>
  <div id="app">
    <header class="mb-5">
      <HeaderBar />
      <router-view />
    </header>
    <main>
      <!-- <SeasonDash/> -->
    </main>
  </div>
</template>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
