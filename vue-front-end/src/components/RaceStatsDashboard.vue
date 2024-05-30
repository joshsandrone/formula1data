<template>
  <div class="mx-10 flex mt-4">
    <SelectionPanel :panelData="selectionPanel"  @dropdown-selected="onSeasonChange" @button-selected="onRaceChange" />

    <div class="flex flex-col w-full">
    </div>
  </div>


</template>




<script>
import api from '@/services/api';
import SelectionPanel from './Generic/SelectionPanel.vue'


export default {
  components: {
    SelectionPanel
  },
  data() {
    return {
      selectedSeason: "2024",
      selectionPanel: {
        name: "seasonRacePanel",
        menuData: {
          name : "raceMenu",
          buttons : {}
        }
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
      this.selectionPanel.menuData.buttons = seasonProfile.races.map(race => {
        return {
          name : race.location,
          color : "red"
        }
      })
    },

    onSeasonChange(season){
      this.selectedSeason = season;
      this.fetchSeasonProfile()
    },

    onRaceChange(race){
      console.log(race);
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
