 <template>
 
  <div class="mx-2 md:mx-10 flex flex-col md:flex-row mt-4 w-full">

    <div class="">
      <div class="flex flex-col">
      <SelectionPanel :panelData="selectionPanel" :selectedMenuOption="selectedDriver"  @dropdown-selected="onSeasonChange" @button-selected="onDriverChange" @secondary-button-selected="onViewChange" />
      </div>
    </div>

    <div class="dash-right">
      <div class="flex flex-col">
        <div class="flex flex-col md:flex-row mb-4 md:mb-0">
          <p class="flex-1"></p>
          <div class="flex-1 hidden md:block h-16">
            <h3 class="text-xl font-semibold"> {{ selectedDriver }}</h3>
            <p class="text-sm font-medium">{{ selectedDriverTeam.name }} </p>
          </div>

          <div class="flex-1 md:text-right mr-4 hidden md:block">
            <PrimeSelectButton v-model="statsType" :options="statTypes" @change="changeStatType" class="mt-2"/>
          </div>

          
        </div>
        <SeasonRaceStats id="statsDashboard" class="season-view"  :statsType="statsType"  :selectedDriver="selectedDriver" :selectedSeason="selectedSeason" :color="selectedDriverTeam.primaryColor" :teamPoints="selectedDriverTeam.points"/>
        <RacesTable id="seasonCalender" class="hidden season-view" :rows="racesTable" />
      </div>
    </div>
  </div>

</template>


<script>
import api from '@/services/api';
import SelectionPanel from './Generic/SelectionPanel.vue'
import SeasonRaceStats from './DriverSeasonData.vue'
import RacesTable from './Generic/GraphicTable.vue'


export default {
  components: {
    SeasonRaceStats,
    SelectionPanel,
    RacesTable
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
          buttons : [],
          secondaryButtons : [
            {
              name: "Season Calender",
              color: "orange"
            },
            {
              name: "Championships",
              color: "orange"
            }
          ],
        }
      },
      racesTable: [],
      statsType : "Races",
      statTypes : ["Races", "Qualys"]
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
        this.setChampionshipTable(response.data);
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

    setChampionshipTable(season){
      this.racesTable = season.races.map(r => {
        return {
          date : r.date,
          location : r.location,
          laps : `${r.laps} laps`
        }
      })
      console.log("championship table = ", this.racesTable)
    },

    onSeasonChange(season){
      this.selectedSeason = season;
      this.fetchSeasonProfile()
    },

    onDriverChange(driver){
      this.selectedDriver = driver;
      this.setTeamProfile();
      this.onViewChange("Stats Dashboard")
    },

    onViewChange(view){
      console.log("View change - ", view)

      for (const seasonView of document.querySelectorAll(".season-view")){
        seasonView.classList.add("hidden")
      }

      if (view === "Season Calender"){
        document.querySelector("#seasonCalender").classList.remove("hidden");
      }

      else if (view === "Stats Dashboard"){
        document.querySelector("#statsDashboard").classList.remove("hidden");
      }


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


<style>
/* On md and below, set card padding to 0.8 (from 1.5)*/
@media (max-width: 768px){
  .p-card-body{
    
    padding: 0.8rem !important;
  }
}

@media (min-width: 769px){
  .dash-right{
  width: calc(100vw - 325px);
  }
}

</style>
