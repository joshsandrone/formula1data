<template>
  <div class="mx-10">
    <div class="flex mb-5 items-center justify-between">
      <SeasonSelectDropdown  @season-selected="onSeasonChange" />
      <div>
        <h3 class="text-xl font-semibold"> {{ selectedDriver }}</h3>
        <p class="text-sm font-medium">{{ selectedDriverTeam.name }} </p>
      </div>
      
      <p></p>
    </div>

      <div class="flex w-full">
        <DriverMenu :menuData="driverMenuData" @driverMenu-button-selected="setSelectedDriver" class="hidden md:block" style="width:230px;"/>
        <SeasonRaceStats :selectedDriver="selectedDriver" :selectedSeason="selectedSeason" :color="selectedDriverTeam.primaryColor" :teamPoints="selectedDriverTeam.points"/>
      </div>
  </div>
</template>


<script>
import api from '@/services/api';
import SeasonRaceStats from './DriverSeasonData.vue'
import DriverMenu from './Generic/SideMenu.vue'
import SeasonSelectDropdown from './Generic/SeasonSelectDropdown.vue'


export default {
  components: {
    DriverMenu,
    SeasonRaceStats,
    SeasonSelectDropdown
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
        primaryColor : null,
        points : null
      },
      teams : [],
      driverMenuData: {
        name : "driverMenu",
        buttons : {}
      }
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

    async fetchSeasonProfile(){
      try {
        const response = await api.getSeasonProfile(this.selectedSeason);
        this.setTeamProfiles(response.data);
        this.setTeamProfile();
        this.setDrivers(response.data);
        this.setDriverMenu(response.data);
      } catch (error) {
        console.error('Failed to fetch season profile:', error);
      }
    },

    setSelectedDriver(driver) {
        this.selectedDriver = driver;
        this.setTeamProfile()
    },

    setDrivers(seasonProfile){
      this.drivers = seasonProfile.teams.map((team) => {
        return team.drivers.map(driver => driver.name);
      }).flat()
    },

    setDriverMenu(seasonProfile){
      this.driverMenuData.buttons = seasonProfile.teams.map(team => {
        return team.drivers.map(driver => {
          return {
            name : driver.name,
            color : team.primaryColor
          }
        }).flat()
      }).flat()
    },

    setTeamProfiles(seasonProfile){
      this.teams = seasonProfile.teams
    },

    setTeamProfile(){
      for ( const team of this.teams ){
        for ( const driver of team.drivers ){
          if (driver.name === this.selectedDriver){
            this.selectedDriverTeam.name = team.name;
            this.selectedDriverTeam.primaryColor = team.primaryColor;
            this.selectedDriverTeam.points = team.points;
          }
        }
      }
    },

    onSeasonChange(season){
      this.selectedSeason = season;
      this.fetchSeasonProfile()
    },

    updateUrl(){
      // const newUrl = `seasonData/${this.selectedSeason}/${this.selectedDriver}`;
      // console.log("new new url ==== ", newUrl);
      // // Update the URL without reloading the page
      // this.$router.push(newUrl);
    },

    loadFromURL(){
      const { season, driver } = this.$route.params;
      this.selectedSeason = season || "2024";
      this.selectedDriver = driver || "Max Verstappen";

      // this.updateUrl();
      this.fetchSeasonProfile();
      this.setSelectedDriver(this.selectedDriver)
    }
  }
};
</script>
