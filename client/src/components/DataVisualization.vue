<template>
  <div class="container">
    <!-- Dropdown menu for chart type -->
    <div class="dropdown-container">
      <div class="dropdown">
        <select v-model="selectedChart" class="select">
        <option value="bar">Bar Chart</option>
        <option value="line">Line Chart</option>
        <option value="pie">Pie Chart</option>
      </select>
      </div>
      <div class="dropdown">
          <select v-model="selectedCategory" @change="updateBackground" class="select">
          <option value="weather">Weather</option>
          <option value="internet">Internet Tower</option>
          <option value="birthRate">Birth Rate</option>
          <!-- Add more options as needed -->
        </select>
      </div>
      <div class="button">
        <button @click="updateChart" class="update-button">submit</button>
     </div>
    </div>

    <!-- Chart.js canvas -->
    <div class="chart-container">
      <canvas id="myChart"></canvas>
    </div>
  </div>
</template>

<script>
import Chart from 'chart.js/auto';
import axios from 'axios';

export default {
  data() {
    return {
      selectedChart: 'bar',
      selectedCategory: 'weather',
      chart: null,
      Data: {},
      backgroundImage: '',
    };
  },
  mounted() {
    this.updateChart();
    this.updateBackground();
  },
  methods: {
    updateBackground() {
      // Update background image based on the selected category
      switch (this.selectedCategory) {
        case 'weather':
          this.backgroundImage = 'https://i.ibb.co/7rgMbsg/skate-about-section-bg-xl-jpg-adapt-1920w.jpg';
          break;
        case 'internet':
          this.backgroundImage = 'https://i.ibb.co/wgqr5yZ/FREESTYLIN-FINAL-LANDSCAPE-1920-X1080-JPG-9928423d535d3d18f841.jpg';
          break;
        case 'birthRate':
          this.backgroundImage = 'https://i.ibb.co/wgqr5yZ/FREESTYLIN-FINAL-LANDSCAPE-1920-X1080-JPG-9928423d535d3d18f841.jpg';
          break;
        // Add more cases for additional categories
        default:
          this.backgroundImage = 'https://i.ibb.co/7rgMbsg/skate-about-section-bg-xl-jpg-adapt-1920w.jpg';
      }
      // Apply the background image to the body or any container element
      // Render chart with the fetched data
      document.querySelector('.container').style.backgroundImage = this.backgroundImage;

    
    },
    async updateChart() {
      // Fetch data based on the selected category
      await this.fetchData();

      // Render chart with the fetched data
      this.renderChart(this.Data.data,this.selectedChart);
    },
    async fetchData() {
      try {
        // Fetch data from API based on the selected category
        const response = await axios.get(`http://localhost:8081/${this.selectedCategory}`);

        // Assuming the API response contains data suitable for Chart.js
        this.Data = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    renderChart(data, chartType) {
    // Destroy previous chart if exists
    if (this.chart) {
        this.chart.destroy();
    }
    let labels = data.map((item) => `${item.date}`);
    if(this.selectedCategory=="weather"){
      labels = data.map((item) => `${item.date}`);
    }else if(this.selectedCategory=="internet"){
      labels = data.map((item) => `${item.location}`);

    }else{
      labels = data.map((item) => `${item.year}`);
    }

    // Extract labels and datasets dynamically from the provided data
    const datasets = [];

    // Iterate over each field in the first data object
    for (const field in data[0]) {
        if (field !== 'undefined'&&field !== 'date'&&field !== 'year'&&field !== 'location') { // Exclude undefined fields
            const dataset = {
                label: field, // Use the field name as the label
                data: data.map(item => item[field]), // Extract data for the field
                backgroundColor: this.generateRandomColor(),
                borderColor: this.generateRandomColor(),
                borderWidth: 1
            };
            datasets.push(dataset);
        }
    }

    // Render new chart
        const ctx = document.getElementById('myChart').getContext('2d');
        this.chart = new Chart(ctx, {
            type: chartType,
            data: {
                labels: labels,
                datasets: datasets
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });
    },
    generateRandomColor() {
        // Generate random color in hexadecimal format
        return '#' + Math.floor(Math.random()*16777215).toString(16);
    },


    
  }
}
</script>

<style scoped>
.container {
  padding: 20px;
  background-color: #333; /* Dark background */
  color: #fff; /* Light text */
  font-family: 'Arial', sans-serif; /* Change the font as needed */
  height: 100vh; /* Set container height to full viewport height */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.dropdown-container {
  display: flex;
  justify-content: space-between; /* Distribute items evenly along the main axis */
  width: 100%; /* Take up full width of the container */
  margin-bottom: 10px;
}

.dropdown {
  width: 20%; /* Set width to 48% to leave some space between dropdowns */
}

.select {
  width: 100%; /* Take up full width of the dropdown */
  padding: 8px;
  font-size: 16px;
  background-color: #555; /* Dark dropdown background */
  color: #fff; /* Light dropdown text */
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.select:focus {
  outline: none;
}

.chart-container {
  width: 80%; /* Adjust the width as needed */
  max-width: 600px; /* Set maximum width for the chart */
  margin-top: 20px; /* Adjust the top margin as needed */
  background-color: #fff;
}
.button {
  width: 100%; /* Take up full width of the container */
  margin-bottom: 10px;
}
</style>