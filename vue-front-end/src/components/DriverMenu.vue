<template>
  <div class="">
    <div v-for="(driver, index) in driversWithColors" :key="index" class="flex">
      <span class="square" :style="{ backgroundColor: driver.color }"></span>
      <PrimeButton class="mx-2 px-1 rounded-none hover:bg-gray-200 rounded-md" :label="driver.name" @click="driverSelection(driver.name)"/>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    drivers: Array,
    teams: Array
  },
  data() {
    return {
      driversWithColors: []
    };
  },
  created() {
    this.setDriverColor();
  },
  watch: {
    drivers: 'setDriverColor',
    teams: 'setDriverColor',
  },
  methods: {
    setDriverColor() {
      // Create a local copy of drivers with colors
      this.driversWithColors = this.drivers.map(driver => {
        let driverColor = 'red'; // default color
        for (let team of this.teams) {
          if (team.drivers.map(d => d.name).includes(driver)) {
            driverColor = team.primaryColor;
            break;
          }
        }
        return {
          name: driver, 
          color: driverColor
        };
      });
    },

    driverSelection(driver) {
      this.$emit('driver-selected', driver);
    }
  }
};
</script>

<style>
.square {
  margin-top: 4px;
  height: 15px;
  width: 15px;
  background-color: red;
  border-radius: 5px;
}
</style>
