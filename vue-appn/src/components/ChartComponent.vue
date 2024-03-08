<template>
    <div>
      <h1>Virat Kohli Stats</h1>
      <!-- The v-if is used to conditionally render a block -->
      <Bar id="my-chart-id" v-if="loaded" :options="chartOptions" :data="chartData" :width="600" />
    </div>
  </template>

  <script>
  import { Bar } from 'vue-chartjs'
  import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'

  ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

  export default {
    name: 'BarChart',
    components: { Bar },
    data: () => ({
      // Prevents chart to mount before the API data arrives
      loaded: false,
      chartData: {
        labels: [],
        datasets: [
          {
            label: 'Matches Played',
            data: [],
            backgroundColor: 'rgba(54, 162, 235, 0.2)'
          }
        ]
      },
      chartOptions: {
        responsive: true
      }
    }),
    async mounted() {
      const apiUrl = 'http://localhost:8080/datasets'

      // Make an HTTP request to fetch the data from the API endpoint
      await fetch(apiUrl)
        .then((response) => response.json())
        .then((data) => {
          // Extract data from the API response and update the chartData
          this.chartData.labels = data.map((item) => item.format) // Use "format" field as labels
          this.chartData.datasets[0].data = data.map((item) => item.matches) // Use "matches" field as data

          // Allow the chart to display the data from the API endpoint
          this.loaded = true
        })
        .catch((error) => {
          console.error('Error fetching data:', error)
        })
    }
  }
  </script>
