<template>
  <div>
    <PrimeCard class="bg-gray-50 mx-2">
        <template #content>
              <h4 class="font-medium">{{ title }}</h4>
             <PrimeChart type="line" :data="graphData" :options="graphOptions" class="h-30rem" style="min-height:250px;"/>
        </template>
    </PrimeCard>
  </div>
</template>


<script>
export default {
  props: [
    'graphInputData',
    'primaryColor',
    'title'
  ],
  data() {
    return {
        graphData: null,
        graphOptions: null,
        graphPointColors : []
    };
  },
//   mounted() {
//     this.graphData = this.setGraphData();
//     this.graphOptions = this.setGraphOptions();
//   },
//   created() {

//   },
  watch: {
    graphInputData: 'setGraphData',
  },
  methods: {
    setGraphData(){

      // extract the and color for each data point
      this.graphPointColors = this.graphInputData.dataset.map(r => {
        if (r.color){
          return r.color;
        }
        return this.primaryColor;
      })

        let graphDatasets = [
            {
                    label: this.title,
                    fill: false,
                    borderColor: this.primaryColor,
                    pointBackgroundColor: this.graphPointColors,
                    yAxisID: 'y',
                    tension: 0.1,
                    data: this.graphInputData.dataset.map(r => r.value)
            }
        ]

        // Add zero line (y=0) if set in inputData
        if(this.graphInputData.zeroLine){
            let zeroPointDataSer = [];
            for (let i = 0; i < this.graphInputData.dataset.length; i++){
              zeroPointDataSer.push(0);
            }
            graphDatasets.push(
              {
                    label: 'Zero Line',
                    fill: false,
                    borderColor: '#c8c8c8',
                    pointRadius: 0,
                    yAxisID: 'y',
                    tension: 0.1,
                    data: zeroPointDataSer
              })          
        }

        // Add an average line if set in the inputData
        if (this.graphInputData.average){
            let averagePosDataSet = [];
            for (let i = 0; i < this.graphInputData.dataset.length; i++){
              averagePosDataSet.push(this.graphInputData.average);
            }
            graphDatasets.push(
              {
                    label: 'Average',
                    fill: false,
                    borderColor: '#6b7280',
                    borderDash: [10,5],
                    pointRadius: 0,
                    yAxisID: 'y',
                    tension: 0.1,
                    data: averagePosDataSet
              })
        }

        this.graphData = {
            labels: this.graphInputData.labels,
            datasets: graphDatasets
        };
        this.setGraphOptions();

    },
    setGraphOptions(){
      this.graphOptions = {
        responsive: true,
        maintainAspectRatio : false,
        plugins: {
          legend: {
            display : true,
            position: 'bottom'
          }
        },
        scales: {
          y: {
            min: this.graphInputData.min % 2 == 0 ? this.graphInputData.min : this.graphInputData.min - 1 ,
            max: this.graphInputData.max % 2 == 0 ? this.graphInputData.max : this.graphInputData.max + 1,
            ticks: {
              stepSize: 2
            }
          }
        }
      }
    }
  }
};
</script>
