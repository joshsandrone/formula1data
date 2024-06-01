<template>
  <div class="">
    <div v-for="(button, index) in menuData.buttons" :key="index" class="flex">
      <span class="square" :style="{ backgroundColor: button.color }"></span>
      <PrimeButton class="mx-2 px-1 hover:bg-gray-200 rounded-md" :label="button.name" @click="buttonSelection(button.name)"/>
    </div>

    <div class="mt-4"></div>

    <div v-for="(button, index) in menuData.secondaryButtons" :key="index" class="flex">
      <span class="square" :style="{ backgroundColor: button.color }"></span>
      <PrimeButton class="mx-2 px-1 hover:bg-gray-200 rounded-md" :label="button.name" @click="secondaryButtonSelection(button.name)"/>
    </div>
  </div>
</template>

<script>
export default {
  props: [
    'menuData'
  ],
  data() {
    return {
      driversWithColors: []
    };
  },
  methods: {
    buttonSelection(button) {
      this.$emit(`button-selected`, button);
    },

    secondaryButtonSelection(button) {
      this.$emit(`secondary-button-selected`, button);
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

.circle {
  margin-top: 4px;
  height: 15px;
  width: 15px;
  background-color: red;
  border-radius: 50%;
}
</style>
