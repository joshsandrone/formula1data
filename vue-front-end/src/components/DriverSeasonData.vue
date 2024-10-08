<template>
  <div>
        <div class="w-full" id="race-data">
          <StatBar :stats="seasonRaceDataStatBar" :color="color"/>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center mb-4">
            <StatCircles class="" :stats="seasonRaceDataStatCircles" :primaryColor="color"/>
          </div>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center" >
           <LineGraph title="Race Results" :graphInputData="raceResultsGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
           <LineGraph title="Positions Gained Per Race" :graphInputData="racePositionsGainedGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
          </div>
        </div>


        <div class="w-full hidden" id="qualy-data">
          <StatBar :stats="seasonQualyDataStatBar" :color="color"/>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center mb-4">
            <StatCircles class="" :stats="seasonQualyDataStatCircles" :primaryColor="color"/>
          </div>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center" >
           <LineGraph title="Qualy Results" :graphInputData="qualyResultsGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
           <LineGraph title="Gap To Pole" :graphInputData="qualyGapToPoleGraphData" :primaryColor="color" class="w-full md:w-1/2 mb-4"/>
          </div>
        </div>
  </div>


</template>




<script>
import StatCircles from './DataRepr/StatCircles.vue'
import LineGraph from './DataRepr/LineGraph.vue'
import StatBar from './DataRepr/StatBar.vue'
import fetchDriverSeasonData from '@/mixins/fetchDriverSeasonData.vue';

export default {
  mixins: [fetchDriverSeasonData],
  props: [
    'selectedDriver',
    'selectedSeason',
    'color',
    'teamPoints',
    'statsType'
  ],
  components: {
    StatCircles,
    LineGraph,
    StatBar
  },
  watch: {
    selectedDriver: 'driverChange',
    selectedSeason: 'driverChange',
    statsType: 'toggleStatView'
  },
  mounted(){
    this.driverChange()
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
              "max" : 3.0,
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
      qualyResultsGraphData: {
        dataset : null,
        labels : null,
        average : null,
        min : 0,
        max : 20
      },
      qualyGapToPoleGraphData: {
        dataset : null,
        labels : null,
        average : null,
        min : null,
        max : null,
      },
    };
  },
  methods: {
    async driverChange(){
      let seasonData = await this.fetchDriverSeasonData(this.selectedDriver, this.selectedSeason, this.seasonData);
      this.setSeasonRaceData(seasonData);
      this.setSeasonQualyData(seasonData);
    },

    setSeasonRaceData(seasonData){
      this.seasonRaceDataStatBar = {
        "Races" : seasonData.races.RacesEntered,
        "Wins" : seasonData.races.wins,
        "Podiums" : seasonData.races.podiums,
        "Points" : seasonData.races.points + seasonData.SprintRaces.sprintPoints,
        "DNFs" : seasonData.races.dnfs
      }

      let driverTeamPointPercentage = (seasonData.extra.totalPoints * 100 / (this.teamPoints)).toFixed(1);
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
      this.seasonQualyDataStatCircles.find(r => r.id === "avgGapToPole").value = seasonData.qualifyings.avgGapToPole.toFixed(3);
      this.seasonQualyDataStatCircles.find(r => r.id === "consistency").value = 66;

      this.qualyResultsGraphData = {
        dataset : seasonData.qualifyings.results.map(r => {
          return {value : r.position}
        }),
        labels : seasonData.qualifyings.results.map(r => r.location),
        average : seasonData.qualifyings.avgRacePos,
        min : 0,
        max : 20
      }


      console.log(Math.max(...seasonData.qualifyings.results.map(r => r.gapToPole)) + 0.1)

      this.qualyGapToPoleGraphData = {
        dataset : seasonData.qualifyings.results.map(r => {
          return {value : r.gapToPole}
        }),
        labels : seasonData.qualifyings.results.map(r => r.location),
        min: 0,
        max : Math.max(...seasonData.qualifyings.results.map(r => r.gapToPole)) + 0.1,
        stepSize : 0.2
      }
    },

    toggleStatView(){
      if(this.statsType === "Qualys"){
        document.querySelector("#qualy-data").classList.remove("hidden")
        document.querySelector("#race-data").classList.add("hidden")
      }
      else if(this.statsType === "Races") {
        document.querySelector("#qualy-data").classList.add("hidden")
        document.querySelector("#race-data").classList.remove("hidden")
      }

    },

    onSeasonChange(){
      this.driverChange()
    }
  }
};
</script>
