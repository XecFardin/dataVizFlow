<template>
  <div id="app">
    <ChartComponent />
  </div>
</template>

<script>
import ChartComponent from './components/DataVisualization.vue'; // Import the ChartComponent

export default {
  components: {
    ChartComponent
  }
}
</script>

<style>
/* Add any global styles here */
body {
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
}
</style>
