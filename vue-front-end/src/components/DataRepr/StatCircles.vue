<template>
  <div class="w-full">
    <PrimeCard class="card flex justify-content-center bg-gray-50 mx-2">
      <template #content>
        <div class="flex flex-wrap justify-between mx-10">
          <div class="md:mx-2" v-for="(stat, index) in stats" :key="index">
            <div class="flex flex-col items-center my-1 md:my-0">
              <PrimeKnob v-model="stat.value" :valueColor="primaryColor" rangeColor="#cbd5e0" :min="stat.min" :max="stat.max" :size="knobSize" :valueTemplate="stat.template" readonly />
              <span class="px-2 py-1 rounded-md font-medium hover:bg-gray-200 hover:cursor-default text-xs" @mouseover="toggle(index, $event)" @mouseleave="toggle(index, $event)">{{ stat.desc }}</span>
              <PrimeOverlay v-if="stat.overlay" :ref="'op' + index" class="bg-gray-200">
                <h2>OVERLAY OVERLAY OVERLAY OVERLAY</h2>
              </PrimeOverlay>
            </div>
          </div>
        </div>
      </template>
    </PrimeCard>
  </div>
</template>

<script>
export default {
  props: ['stats', 'primaryColor'],
  data(){
    return {
      knobSize: 110
    }
  },
  mounted() {
    this.computePrimeKnobSize()
  },
  methods: {
    toggle(index, event) {
      if (this.$refs['op' + index]){
        this.$refs['op' + index][0].toggle(event);
      } 
    },
    // depending on the size of the screen,
    // the prime knobs need to change size
    computePrimeKnobSize() {
      let width = window.innerWidth;
      if (width < 768) {
        this.knobSize =  84;
      } 
      else {
        this.knobSize =  115;
      }
    }
  }
}
</script>

