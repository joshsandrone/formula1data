<template>
  <div class="bg-gray-50 mb-4">
      <div class="flex px-5 pt-4 pb-3 items-center justify-between mx-1 md:mx-5">
        <p class="font-medium">F1 Stats</p>
        <PrimeSelectButton v-model="currentPage" :options="pages" @change="changePage" :allowEmpty="false" class=""/>
        <p class="hidden md:block"></p>
      </div>
  </div>
</template>




<script>

export default {

  data(){
    return {
      currentPage: "Season Data",
      pages: ["Season Data", "Race Data"]
    }
  },

  watch:{
    $route : 'onRouteChange'
  },

  methods: {
    changePage(){
      if (this.currentPage === "Season Data"){
        this.$router.push({name : "SeasonStats"})
      }
      else if (this.currentPage === "Race Data"){
        this.$router.push({name : "RaceStats"})
      }
    },

    onRouteChange(to, from){
      if (to.name !== from.name){
        console.log("Page Switch!!!")

        if (to.name === "SeasonStats"){
          this.currentPage = "Season Data";
        }

        else if (to.name === "RaceStats"){
          this.currentPage = "Race Data";
        }


      }
    }
  }
}

</script>
