<template>
    <div class="flex flex-col mb-5 mr-4">
      <div class="text-left" style="height: 70px;">
        <SeasonSelectDropdown  @season-selected="onDropDownChanged" class="mb-4"/>
      </div>
      <SideMenu :menuData="panelData.menuData" @button-selected="onButtonSelected" class="hidden md:block" style="width:230px;"/>
    </div>
</template>



<script>
// import api from '@/services/api';
import SideMenu from './SideMenu.vue'
import SeasonSelectDropdown from './SeasonSelectDropdown.vue'


export default {
  props : [
    "panelData"
  ],
  components: {
    SideMenu,
    SeasonSelectDropdown
  },
  data() {
    return {
    };
  },
//   mounted() {
//     this.loadFromURL();
//   },
//   watch: {
//     // selectedDriver : 'updateUrl',
//     // selectedSeason : 'updateUrl'
//   },
  methods: {
    // setSelectedDriver(driver) {
    //     this.selectedDriver = driver;
    //     this.setTeamProfile()
    // },

    // setSideMenu(seasonProfile){
    //   this.driverMenuData.buttons = seasonProfile.teams.map(team => {
    //     return team.drivers.map(driver => {
    //       return {
    //         name : driver.name,
    //         color : team.primaryColor
    //       }
    //     }).flat()
    //   }).flat()
    // },

    // onSeasonChange(season){
    //   this.selectedSeason = season;
    //   this.fetchSeasonProfile()
    // },

    onButtonSelected(button){
        console.log("Button selected ", button)
        this.$emit(`button-selected`, button);
    },

    onDropDownChanged(button){
        console.log("Dropdown used: ", button)
        this.$emit(`dropdown-selected`, button);
    }

  }
};
</script>
