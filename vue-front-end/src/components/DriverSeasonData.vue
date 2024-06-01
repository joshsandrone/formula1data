<template>
  <div>
        <div class="w-full">
          <StatBar :stats="seasonRaceDataStatBar" :color="color"/>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center mb-4">
            <StatCircles class="" :stats="seasonRaceDataStatCircles" :primaryColor="color"/>
          </div>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center" >
           <LineGraph title="Race Results" :graphInputData="raceResultsGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
           <LineGraph title="Positions Gained Per Race" :graphInputData="racePositionsGainedGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
          </div>
        </div>


        <div class="w-full">
          <StatBar :stats="seasonQualyDataStatBar" :color="color"/>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center mb-4">
            <StatCircles class="" :stats="seasonQualyDataStatCircles" :primaryColor="color"/>
          </div>
          <!-- <div class="flex flex-wrap md:flex-nowrap justify-center items-center" >
           <LineGraph title="Race Results" :graphInputData="raceResultsGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
           <LineGraph title="Positions Gained Per Race" :graphInputData="racePositionsGainedGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
          </div> -->
        </div>
  </div>


</template>




<script>
import api from '@/services/api';
import StatCircles from './DataRepr/StatCircles.vue'
import LineGraph from './DataRepr/LineGraph.vue'
import StatBar from './DataRepr/StatBar.vue'

export default {
  props: [
    'selectedDriver',
    'selectedSeason',
    'color',
    'teamPoints'
  ],
  components: {
    StatCircles,
    LineGraph,
    StatBar
  },
  watch: {
    selectedDriver: 'fetchDriverSeasonStats',
    selectedSeason: 'fetchDriverSeasonStats'
  },
  mounted(){
    this.fetchDriverSeasonStats()
  },
  data() {
    return {
      seasonRaceDataStatBar: {},
      seasonRaceDataStatCircles: [
            {
              "id" : "avgRacePos",
              "desc" : "Avg Race Pos",
              "min" : 0,
              "max" : 20,
              "value": 0,
              "template" : "{value}",
              "overlay" : {}
            },
            {
              "id": "highestPos",
              "desc" : "Highest Result",
              "min" : 0,
              "max" : 20,
              "value": 0,
              "template" : "{value}",
              "overlay" : {}
            },
            {
              "id" : "percentageTeamPoints",
              "desc" : "% Team Points",
              "min" : 0,
              "max" : 100,
              "value": 0,
              "template" : "{value}%",
              "overlay" : {}
            },
            {
              "id" : "consistency",
              "desc" : "Consistency",
              "min" : 0,
              "max" : 100,
              "value": 0,
              "template" : "{value}%",
              "overlay" : {}
            }
      ],
      raceResultsGraphData: {
        dataset : null,
        labels : null,
        average : null,
        min : 0,
        max : 20
      },
      racePositionsGainedGraphData: {
        dataset : null,
        labels : null,
        average : null,
        min : null,
        max : null,
        zeroLine : true
      },
      seasonQualyDataStatBar: {},
      seasonQualyDataStatCircles: [
            {
              "id" : "avgQualyPos",
              "desc" : "Avg Qualy Pos",
              "min" : 0,
              "max" : 20,
              "value": 0,
              "template" : "{value}",
              "overlay" : {}
            },
            {
              "id": "highestPos",
              "desc" : "Highest Position",
              "min" : 0,
              "max" : 20,
              "value": 0,
              "template" : "{value}",
              "overlay" : {}
            },
            {
              "id" : "avgGapToPole",
              "desc" : "Avg Gap To Pole",
              "min" : 0,
              "max" : 1,
              "value": 0,
              "template" : "{value}s",
              "overlay" : {}
            },
            {
              "id" : "consistency",
              "desc" : "Consistency",
              "min" : 0,
              "max" : 100,
              "value": 0,
              "template" : "{value}%",
              "overlay" : {}
            }
      ],
    };
  },
  methods: {
    async fetchDriverSeasonStats() {
      try {
        let [firstname, surname] = this.selectedDriver.split(' ');
        const response = await api.getDriverSeasonStats(this.selectedSeason,firstname,surname);

        this.setSeasonRaceData(response.data)
        this.setSeasonQualyData(response.data)
      } catch (error) {
        console.error('Failed to fetch driver season stats:', error);
      }
    },

    setSeasonRaceData(seasonData){
      this.seasonRaceDataStatBar = {
        "Wins" : seasonData.races.wins,
        "Podiums" : seasonData.races.podiums,
        "Points" : seasonData.races.points + seasonData.SprintRaces.sprintPoints,
        "DNFs" : seasonData.races.dnfs
      }

      let driverTeamPointPercentage = ((seasonData.races.points + seasonData.SprintRaces.sprintPoints) * 100 / (this.teamPoints)).toFixed(1);
      driverTeamPointPercentage = isNaN(driverTeamPointPercentage) ? 0.0 : driverTeamPointPercentage

      this.seasonRaceDataStatCircles.find(r => r.id === "avgRacePos").value = seasonData.races.avgRacePos.toFixed(2);
      this.seasonRaceDataStatCircles.find(r => r.id === "highestPos").value = seasonData.races.highestRacePos;
      this.seasonRaceDataStatCircles.find(r => r.id === "percentageTeamPoints").value = driverTeamPointPercentage;
      this.seasonRaceDataStatCircles.find(r => r.id === "consistency").value = seasonData.races.raceConsistency.toFixed(1);

      let positionsGainedData = seasonData.races.results.map((race) => {
          const qualy = seasonData.qualifyings.results.find(r => r.location === race.location);
          if (!qualy) {return {value : 0};}

          if(race.Dnf){
            return {
              value : (qualy.position - race.position),
              color : "red"
            }
          }
          return {value : (qualy.position - race.position)}
        })

      this.racePositionsGainedGraphData = {
        dataset : positionsGainedData,
        labels : seasonData.races.results.map(r => r.location),
        average : null,
        min : Math.min(...positionsGainedData.map(r => r.value)) - 1,
        max : Math.max(...positionsGainedData.map(r => r.value)) + 1,
        zeroLine: true
      }

      this.raceResultsGraphData = {
        dataset : seasonData.races.results.map(r => {
          if(r.Dnf){
            return {
              value : r.position,
              color : "red"
            }
          }
          return {value : r.position}
        }),
        labels : seasonData.races.results.map(r => r.location),
        average : seasonData.races.avgRacePos,
        min : 0,
        max : 20
      }
    },

    setSeasonQualyData(seasonData){
      this.seasonQualyDataStatBar = {
        "Poles" : seasonData.qualifyings.poles,
        "FrontRows" : seasonData.qualifyings.frontRows,
      }


      this.seasonQualyDataStatCircles.find(r => r.id === "avgQualyPos").value = seasonData.qualifyings.avgQualyPos.toFixed(2);
      this.seasonQualyDataStatCircles.find(r => r.id === "highestPos").value = seasonData.qualifyings.highestQualyPos;
      this.seasonQualyDataStatCircles.find(r => r.id === "avgGapToPole").value = seasonData.qualifyings.avgGapToPole;
      this.seasonQualyDataStatCircles.find(r => r.id === "consistency").value = 66;


    },

    onSeasonChange(){
      this.fetchDriverSeasonStats()
    }
  }
};
</script>
