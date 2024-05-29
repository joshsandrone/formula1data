<template>
  <div class="mx-10">
    <div class="flex mb-5 items-center justify-between">
      <SeasonSelectDropdown  @season-selected="onSeasonChange" />
      <div>
        <!-- <h3 class="text-xl font-semibold"> {{ selectedDriver }}</h3>
        <p class="text-sm font-medium">{{ selectedDriverTeam.name }} </p> -->
      </div>
      
      <p></p>
    </div>

      <div class="flex w-full">
        <RaceMenu :menuData="raceMenuData" @driver-selected="setSelectedDriver" class="hidden md:block" style="width:230px;"/>
        <!-- <SeasonRaceStats :selectedDriver="selectedDriver" :selectedSeason="selectedSeason" :color="selectedDriverTeam.primaryColor" :teamPoints="selectedDriverTeam.points"/> -->
      </div>
  </div>
</template>




<script>
import api from '@/services/api';
import RaceMenu from './Generic/SideMenu.vue'
import SeasonSelectDropdown from './Generic/SeasonSelectDropdown.vue'


export default {
  components: {
    RaceMenu,
    SeasonSelectDropdown
  },
  data() {
    return {
      selectedSeason: "2024",
      raceMenuData: {
          name : "raceMenu",
          buttons : []
      }
    }
  },
  mounted() {
    this.loadFromURL();
  },
  methods: {
    async fetchSeasonProfile(){
      try {
        const response = await api.getSeasonProfile(this.selectedSeason);
        this.setRaceMenu(response.data);
      } catch (error) {
        console.error('Failed to fetch season profile:', error);
      }
    },

    setRaceMenu(seasonProfile){
      console.log(seasonProfile)
      this.raceMenuData.buttons = seasonProfile.races.map(race => {
        return {
          name : race.location,
          color : "red"
        }
      })
    },

    setSeasonSelectOptions(seasons){
      this.seasons = seasons;
    },

    onSeasonChange(season){
      this.selectedSeason = season;
      this.fetchSeasonProfile()
    },


    loadFromURL(){
      // const { season, race } = this.$route.params;
      // this.selectedSeason = season || "2024";
      // this.selectedRace = race || "------";
      // this.updateUrl();
      this.fetchSeasonProfile();
    }



  }
};
</script>
