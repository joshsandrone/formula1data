<template>
  <div class="mx-10">
    <div class="flex mb-5 items-center justify-between">
      <PrimeDropdown v-model="selectedSeason" :options="seasons" class="bg-gray-50 w-30" @change="onSeasonChange" placeholder="Select a Season" />
      <div>
        <h3 class="text-xl font-semibold"> {{ selectedDriver }}</h3>
        <p class="text-sm font-medium">{{ selectedDriverTeam.name }} </p>
      </div>
      
      <p></p>
    </div>

      <div class="flex w-full">
        <DriverMenu :drivers="drivers" :teams="teams" @driver-selected="setSelectedDriver" class="hidden md:block" style="width:230px;"/>
        <SeasonRaceStats :selectedDriver="selectedDriver" :selectedSeason="selectedSeason" :color="selectedDriverTeam.primaryColor"/>
      </div>
  </div>
</template>




<script>
import api from '@/services/api';
import SeasonRaceStats from './DriverSeasonData.vue'
import DriverMenu from './DriverMenu.vue'


export default {
  components: {
    DriverMenu,
    SeasonRaceStats
  },
  data() {
    return {
      seasons: [
        "2024"
      ],
      selectedSeason: "2024",
      drivers: [
        "Max Verstappen"
      ],
      selectedDriver: "Max Verstappen",
      selectedDriverTeam: {
        name : null,
        primaryColor : null
      },
      teams : []
    };
  },
  mounted() {
    this.loadFromURL();
  },
  watch: {
    selectedDriver : 'updateUrl',
    selectedSeason : 'updateUrl'
  },
  methods: {

    async fetchAvailableSeasonProfiles(){
      try {
        const response = await api.getAvailableSeasonProfiles(this.selectedSeason);
        this.setSeasonSelectOptions(response.data);
      } catch (error) {
        console.error('Failed to fetch available season profiles:', error);
      }
    },

    async fetchSeasonProfile(){
      try {
        const response = await api.getSeasonProfile(this.selectedSeason);
        this.setTeamProfiles(response.data);
        this.setTeamProfile();
        this.setDriverSelectOptions(response.data);
      } catch (error) {
        console.error('Failed to fetch season profile:', error);
      }
    },

    setSelectedDriver(driver) {
        this.selectedDriver = driver;
        this.setTeamProfile()
    },

    setDriverSelectOptions(seasonProfile){
      this.drivers = seasonProfile.teams.map((team) => {
        return team.drivers.map(driver => driver.name);
      }).flat()
    },

    setSeasonSelectOptions(seasons){
      this.seasons = seasons;
    },

    setTeamProfiles(seasonProfile){
      this.teams = seasonProfile.teams
    },

    setTeamProfile(){
      for ( const team of this.teams ){
        for ( const driver of team.drivers ){
          if (driver.name === this.selectedDriver){
            this.selectedDriverTeam.name = team.name;
            this.selectedDriverTeam.primaryColor = team.primaryColor
          }
        }
      }
    },

    onSeasonChange(){
      this.fetchSeasonProfile()
    },

    updateUrl(){
      const newUrl = `/${this.selectedSeason}/${this.selectedDriver}`;
      // Update the URL without reloading the page
      this.$router.push(newUrl);
    },

    loadFromURL(){
      console.log("333333333", this.$route)
      console.log("44444444", window.location)
      console.log("5555555", this.$route.params)
      const { season, driver } = this.$route.params;

      console.log("66666666", season, driver)

      this.selectedSeason = season || "2024";
      this.selectedDriver = driver || "Max Verstappen";

      // this.updateUrl();
      this.fetchAvailableSeasonProfiles();
      this.fetchSeasonProfile();
      this.setSelectedDriver(this.selectedDriver)
    }
  }
};
</script>
