<template>
  <div class="mx-10 flex mt-4">
    <SelectionPanel :panelData="selectionPanel" :selectedMenuOption="selectedRace"  @dropdown-selected="onSeasonChange" @button-selected="onRaceChange" />

    <div class="flex flex-col w-full">
      <p>ss = {{ selectedRace }}</p>
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
      selectedRace: null,
      selectionPanel: {
        name: "seasonRacePanel",
        menuData: {
          name : "raceMenu",
          buttons : []
        }
      }
    }
  },
  mounted() {
    this.loadFromURL();
  },
  watch: {
    selectedSeason : 'updateUrl',
    selectedRace  : 'updateUrl'
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
      this.selectedRace = race;
    },

    updateUrl(){
      this.$router.push({name : "RaceStats", params : {season : this.selectedSeason, race: this.selectedRace}})
    },

    loadFromURL(){
      const { season, race } = this.$route.params;
      this.selectedSeason = season || "2024";
      this.selectedRace = race || "Bahrain";
      this.fetchSeasonProfile();
    }
  }
};
</script>
