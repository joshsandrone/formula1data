<template>
    <div class="flex justify-center md:justify-normal md:flex-col md:mr-4">
      <div class="md:text-left md:h-16">
        <SeasonSelectDropdown  @season-selected="onDropDownChanged" class="mb-4"/>
      </div>
      <SideMenu :menuData="panelData.menuData" @button-selected="onButtonSelected" @secondary-button-selected="onSecondaryButtonSelected" class="hidden md:block" style="width:230px;"/>
      <div class="md:hidden ml-4">
        <PrimeDropdown v-model="MenuDropdownOption" @change="onMenuDropDownChanged" :options="panelData.menuData.buttons.map(r => r.name)" class=" bg-gray-50 w-30"  />
      </div>
    </div>
</template>




<script>
import SideMenu from './SideMenu.vue'
import SeasonSelectDropdown from './SeasonSelectDropdown.vue'

export default {
  props : [
    "panelData",
    "selectedMenuOption"
  ],
  data() {
    return {
        MenuDropdownOption: this.selectedMenuOption
    }
  },
  components: {
    SideMenu,
    SeasonSelectDropdown
  },
  methods: {
    onButtonSelected(value){
        this.$emit(`button-selected`, value);
    },
    onSecondaryButtonSelected(value){
        this.$emit(`secondary-button-selected`, value);
    },
    onDropDownChanged(value){
        this.$emit(`dropdown-selected`, value);
    },
    // On small screens, menu is designed to fold into a dropdown
    onMenuDropDownChanged(){
        this.$emit(`button-selected`, this.MenuDropdownOption);
    }
  }
};
</script>

<style>

</style>
