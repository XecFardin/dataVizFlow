<template>
  <div class="container">
    <div class="dropdown-container">
      <div class="dropdown">
        <select v-model="selectedChart" @change="updateChart" class="select">
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
        </select>
      </div>
    </div>

    <button @click="updateChart" class="update-button">Update Category</button>

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
      Data: {},
      backgroundImage: '',
    };
  },
  mounted() {
    this.updateChart();
    this.updateBackground();
  },
  methods: {
    async updateChart() {
      await this.fetchData();

      this.renderChart(this.Data.data,this.selectedChart);
    },
    async fetchData() {
      try {
        const response = await axios.get(`http://localhost:8081/${this.selectedCategory}`);

        this.Data = response.data;
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    renderChart(data, chartType) {
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

    const datasets = [];

    for (const field in data[0]) {
        if (field !== 'undefined'&&field !== 'date'&&field !== 'year'&&field !== 'location') { 
            const dataset = {
                label: field, 
                data: data.map(item => item[field]),
                backgroundColor: this.generateRandomColor(),
                borderColor: this.generateRandomColor(),
                borderWidth: 1
            };
            datasets.push(dataset);
        }
    }

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
        return '#' + Math.floor(Math.random()*16777215).toString(16);
    },


    updateBackground() {
      switch (this.selectedCategory) {
        case 'weather':
          this.backgroundImage = 'url(https://i.ibb.co/qxZSHvb/anime-style-clouds-23-2151071681.jpg)';
          break;
        case 'internet':
          this.backgroundImage = 'url(https://i.ibb.co/XSs4tJh/531062.jpg)';
          break;
        case 'birthRate':
          this.backgroundImage = 'url(https://i.ibb.co/rMxNN20/supermassive-black-holes-birth-to-stars-wallpaper-preview.jpg)';
          break;
        default:
          this.backgroundImage = 'url(https://i.ibb.co/7rgMbsg/skate-about-section-bg-xl-jpg-adapt-1920w.jpg)';
      }
      document.querySelector('.container').style.backgroundImage = this.backgroundImage;
    }
    
  }
}
</script>

<style scoped>
.container {
  padding: 20px;
  background-color: #333; 
  color: #fff;
  font-family: 'Arial', sans-serif; 
  height: 100vh; 
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.dropdown-container {
  display: flex;
  justify-content: space-between;
  width: 100%; 
  margin-bottom: 10px;
}

.dropdown {
  width: 20%; 
}

.select {
  width: 100%; 
  padding: 8px;
  font-size: 16px;
  background-color: #555; 
  color: #fff; 
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.select:focus {
  outline: none;
}

.chart-container {
  width: 80%;
  max-width: 600px;
  margin-top: 20px; 
  background-color: #fff;
}

.update-button {
  margin-top: 10px;
  padding: 8px 16px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.update-button:hover {
  background-color: #0056b3;
}
</style>