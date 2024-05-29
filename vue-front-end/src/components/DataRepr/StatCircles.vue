<template>
  <div class="w-full">
    <PrimeCard class="card flex justify-content-center bg-gray-50 mx-2 mb-4">
      <template #content>
        <h4 class="font-semibold text-xl mb-2">{{ title }}</h4>
        <div class="flex flex-wrap justify-between mx-10">
          <div class="mx-2" v-for="(stat, index) in stats" :key="index">
            <div class="flex flex-col items-center">
              <PrimeKnob v-model="stat.value" :valueColor="primaryColor" rangeColor="#cbd5e0" :min="stat.min" :max="stat.max" :valueTemplate="stat.template" readonly />
              <span class="px-2 py-1 rounded-md font-medium hover:bg-gray-200 hover:cursor-default" @mouseover="toggle(index, $event)" @mouseleave="toggle(index, $event)">{{ stat.desc }}</span>
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
  methods: {
    toggle(index, event) {
      if (this.$refs['op' + index]){
        this.$refs['op' + index][0].toggle(event);
      } 
    },
  },
};
</script>
