<template>
<div>
  <div  v-for="(row, index) in rows" :key="index" class="bg-gray-50 mb-1 flex rounded-lg px-3 py-1 hover:bg-gray-100" >
      <router-link to="/raceData" class="flex w-full font-medium text-md cursor-pointer justify-between"><p  class="w-1/3" v-for="(value, key) in row" :key="key">{{ value }} </p></router-link>
  </div>
</div>


</template>




<script>

export default {
  props : [
    "rows"
  ]
}

</script>
