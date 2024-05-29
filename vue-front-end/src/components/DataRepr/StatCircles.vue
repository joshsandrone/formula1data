<template>
  <div class="w-full">
      <PrimeCard class="card flex justify-content-center bg-gray-50 mx-2 mb-4">
          <template #content>
            <h4 class="font-semibold text-xl mb-2">{{ title }}</h4>
            <div class="flex flex-wrap justify-between mx-10">
              <div class="mx-5" v-for="(stat,index) in stats" :key="index">
                <PrimeKnob v-model="stat.value" :valueColor="primaryColor" :min="stat.min" :max="stat.max" :valueTemplate="stat.template" readonly/>
                <span>{{ stat.desc }}</span>
              </div>

            </div>
          </template>
      </PrimeCard>
  </div>
</template>

<script>
export default {
  props: [
    'stats',
    'primaryColor',
  ]
};
</script>
