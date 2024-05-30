 <template>
 
  <div class="mx-10 flex mt-4">
    <SelectionPanel :panelData="selectionPanel"  @dropdown-selected="onSeasonChange" @button-selected="onDriverChange" />

    <div class="flex flex-col w-full">
      <div style="height: 70px;">
        <h3 class="text-xl font-semibold"> {{ selectedDriver }}</h3>
        <p class="text-sm font-medium">{{ selectedDriverTeam.name }} </p>
      </div>
      <SeasonRaceStats :selectedDriver="selectedDriver" :selectedSeason="selectedSeason" :color="selectedDriverTeam.primaryColor" :teamPoints="selectedDriverTeam.points"/>
    </div>
  </div>

</template>


<script>
import api from '@/services/api';
import SelectionPanel from './Generic/SelectionPanel.vue'
import SeasonRaceStats from './DriverSeasonData.vue'


export default {
  components: {
    SeasonRaceStats,
    SelectionPanel
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
        name : "Red Bull Racing Honda RBPT",
        primaryColor : null,
        points : null
      },
      teams : [],
      selectionPanel: {
        name: "seasonDriverPanel",
        menuData: {
          name : "driverMenu",
          buttons : {}
        }
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

    setDrivers(seasonProfile){
      this.drivers = seasonProfile.teams.map((team) => {
        return team.drivers.map(driver => driver.name);
      }).flat()
    },

    setDriverMenu(seasonProfile){
      this.selectionPanel.menuData.buttons = seasonProfile.teams.map(team => {
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

    onDriverChange(driver){
      this.selectedDriver = driver;
      this.setTeamProfile();
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
      this.onDriverChange(this.selectedDriver)
    }
  }
};
</script>
