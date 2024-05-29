<template>
        <div class="w-full" style="max-width: 1200px; margin-left: auto; margin-right:auto;">
          <StatBar :stats="statBarDataRace" :color="colorr"/>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center">
            <StatCircles class="" :stats="seasonData1Race" :primaryColor="color"/>
          </div>
          <div class="flex flex-wrap md:flex-nowrap justify-center items-center" >
           <LineGraph title="Race Results" :graphInputData="raceResultsGraphData" :primaryColor="color" class="w-full"/>
           <LineGraph title="Positions Gained Per Race" :graphInputData="positionGainedGraphData" :primaryColor="color" class="w-full"/>
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
      seasonData : {},
      statBarDataRace: {},
      seasonData1Race: [
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
      positionGainedGraphData: {
        dataset : null,
        labels : null,
        average : null,
        min : null,
        max : null
      }
    };
  },
  methods: {
    async fetchDriverSeasonStats() {
      console.log("Fetching!!!!!")
      try {
        let [firstname, surname] = this.selectedDriver.split(' ');
        const response = await api.getDriverSeasonStats(this.selectedSeason,firstname,surname);

        this.seasonData = response.data;
        this.setSeasonRaceData(response.data)
      } catch (error) {
        console.error('Failed to fetch driver season stats:', error);
      }
    },

    setSeasonRaceData(seasonData){
      this.statBarDataRace = {
        "Wins" : seasonData.races.wins,
        "Podiums" : seasonData.races.podiums,
        "Points" : seasonData.races.points + seasonData.SprintRaces.sprintPoints,
        "DNFs" : seasonData.races.dnfs
      }

      let driverTeamPointPercentage = ((seasonData.races.points + seasonData.SprintRaces.sprintPoints) * 100 / (this.teamPoints)).toFixed(1);
      driverTeamPointPercentage = isNaN(driverTeamPointPercentage) ? 0.0 : driverTeamPointPercentage

      this.seasonData1Race.find(r => r.id === "avgRacePos").value = seasonData.races.avgRacePos.toFixed(2);
      this.seasonData1Race.find(r => r.id === "highestPos").value = seasonData.races.highestRacePos;
      this.seasonData1Race.find(r => r.id === "percentageTeamPoints").value = driverTeamPointPercentage;
      this.seasonData1Race.find(r => r.id === "consistency").value = seasonData.races.raceConsistency.toFixed(1);

      let positionsGainedData = seasonData.races.results.map((race) => {

        console.log("racerrrrr = ", race)

          const qualy = seasonData.qualifyings.results.find(r => r.location === race.location);
          if (!qualy) {return {value : 0};}

          if(race.Dnf){
            console.log("Found RACE WITH DNF")
            return {
              value : (qualy.position - race.position),
              color : "red"
            }
          }
          return {value : (qualy.position - race.position)}
        })

        console.log(positionsGainedData, Math.min(...positionsGainedData))

      this.positionGainedGraphData = {
        dataset : positionsGainedData,
        labels : seasonData.races.results.map(r => r.location),
        average : null,
        min : Math.min(...positionsGainedData.map(r => r.value)) - 1,
        max : Math.max(...positionsGainedData.map(r => r.value)) + 1
      }

      console.log("££££££ = ", seasonData)

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

    onSeasonChange(){
      this.fetchDriverSeasonStats()
    }
  }
};
</script>
