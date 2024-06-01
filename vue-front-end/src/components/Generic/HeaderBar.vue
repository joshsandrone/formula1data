<template>
  <div class="bg-gray-50 mb-4">
      <div class="flex px-5 pt-4 pb-3 items-center justify-between mx-1 md:mx-5">
        <p class="font-medium">F1 Stats</p>
        <div id="header-buttons" class="text-light ">
          <!-- <router-link to="/"> <PrimeButton id="headerButton1" class="header-button px-2" @click="headerButtonClicked('headerButton1')" label="Career Data"></PrimeButton></router-link>
          <span class="mx-3"></span> -->
          <router-link to="/seasonData"><PrimeButton id="headerButton2" class="header-button bg-gray-200 px-2" @click="headerButtonClicked('headerButton2')" label="Season Data">   </PrimeButton></router-link>
          <span class="mx-1 md:mx-3"></span>
          <router-link to="/raceData"><PrimeButton id="headerButton3" class="header-button px-2" @click="headerButtonClicked('headerButton3')" label="Race Data"></PrimeButton></router-link>
        </div>
        <p class="hidden md:block"></p>
      </div>
  </div>
</template>




<script>

export default {
  methods: {
    headerButtonClicked(id){
      const headerButtons = document.querySelector("#header-buttons").querySelectorAll(".header-button");
      for (const button of headerButtons){
        button.classList.remove("bg-gray-200");
      }

      document.getElementById(id).classList.add("bg-gray-200");
    }
  }
}

</script>
