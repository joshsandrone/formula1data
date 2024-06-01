<template>
<div>
  <div  v-for="(row, index) in rows" :key="index" class="bg-gray-50 mb-1 flex rounded-lg px-3 py-1" >
      <p class="flex-1 font-medium text-md" v-for="(value, key) in row" :key="key">{{ value }} </p>
  </div>
</div>

      <!-- <PrimeCard  v-for="(row, index) in rows" :key="index" class="bg-gray-50 my-1">
        <template #content>
        <div class="flex">
            <p class="flex-1 font-semibold" v-for="(value, key) in row" :key="key">{{ value }} </p>
        </div>
        </template>
    </PrimeCard> -->


</template>




<script>

export default {
  props : [
    "rows"
  ]
}

</script>
