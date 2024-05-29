<template>
  <div>
    <PrimeCard class="bg-gray-50 mx-2 mb-4">
        <template #content>
             <div class="flex flex-wrap md:flex-nowrap justify-between items-center mx-10">
                <span v-for="(value, key) in stats" :key="key" class="px-2 py-1 rounded-md font-semibold hover:bg-gray-200 hover:cursor-pointer">
                    <span :style="{ color: color }">{{ value }}</span> {{ key }}
                </span>
             </div>
        </template>
    </PrimeCard>
  </div>
</template>





<script>
export default {
  props: {
    color : String,
    stats : Object
  },
  mounted() {
    this.setStats()
  },
  watch: {
    stats : "setStats"
  },
  methods : {
    setStats(){
        console.log("Setting stats")
        console.log(this.stats)
    }
  }
};
</script>
