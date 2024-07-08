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
            <PrimeSelectButton v-model="statsType" :options="statTypes" @change="changeStatType" :allowEmpty="false" class="mt-2"/>
          </div>

          
        </div>
        <SeasonRaceStats id="stats-dashboard" class="season-view"  :statsType="statsType"  :selectedDriver="selectedDriver" :selectedSeason="selectedSeason" :color="selectedDriverTeam.primaryColor" :teamPoints="selectedDriverTeam.points"/>
        <RacesTable id="season-calender" class="hidden season-view" :rows="racesTable" />
        <Head2Head id="head-2-head" class="hidden season-view" :drivers="drivers" :season="selectedSeason" />
      </div>
    </div>
  </div>

</template>


<script>
import api from '@/services/api';
import SelectionPanel from './Generic/SelectionPanel.vue'
import SeasonRaceStats from './DriverSeasonData.vue'
import RacesTable from './Generic/GraphicTable.vue'
import Head2Head from './Head2Head.vue';


export default {
  components: {
    SeasonRaceStats,
    SelectionPanel,
    RacesTable,
    Head2Head
  },
  data() {
    return {
      selectedSeason: "2024",
      selectedDriver: "Max Verstappen",
      selectedDriverTeam: {
        name : "Red Bull Racing Honda RBPT",
        primaryColor : null,
        points : null
      },
      teams : [],
      drivers: [],
      selectionPanel: {
        name: "seasonDriverPanel",
        menuData: {
          name : "driverMenu",
          buttons : [],
          secondaryButtons : [
            {
              name: "Season Calender",
              viewlinkid: "season-calender",
              color: "orange"
            },
            {
              name: "Championships",
              viewlinkid: "championships",
              color: "orange"
            },
            {
              name: "Head-2-Head",
              viewlinkid: "head-2-head",
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
    selectedDriver(newval){
      this.onViewChange(newval, false)
    },
    selectedSeason(){
      this.onViewChange(this.currentViewId, false);
    }
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
      this.drivers = seasonProfile.teams.map(team => {
        return team.drivers.map(driver => driver.name)}).flat();
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
          row : {
            date : r.date,
            location : r.location,
            laps : `${r.laps} laps`
          },
          link: {
            name: "RaceStats",
            params : {
              season : season.season,
              race : r.location
            }
          }
        }
      })
    },

    onSeasonChange(season){
      this.selectedSeason = season;
      this.fetchSeasonProfile()
    },

    onDriverChange(driver){
      this.selectedDriver = driver;
      this.setTeamProfile();
      this.onViewChange("stats-dashboard")
    },

    onViewChange(viewid, changeOfSeasonView = true){
      console.log("View change - ", viewid)

      if (changeOfSeasonView){
        for (const seasonView of document.querySelectorAll(".season-view")){
          seasonView.classList.add("hidden")
        }

        document.getElementById(viewid).classList.remove("hidden");
      }

      this.updateUrl(viewid);
    },

    updateUrl(page){
      this.$router.push({name : "SeasonStats", params : {season : this.selectedSeason, page: page}})
      this.currentViewId = page;
    },

    loadFromURL(){
      const { season, page } = this.$route.params;
      this.selectedSeason = season || "2024";
      this.fetchSeasonProfile();

      const seasonViewIds = Array.from(document.querySelectorAll(".season-view")).map( el => el.id );
      if (seasonViewIds.includes(page)){
        this.onViewChange(page);
        return;
      }

      // if not passed a view, currenlty assume a driver has been passed, put them in view.
      this.selectedDriver = page || "Max Verstappen";
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
