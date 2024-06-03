<template>
<div>
  <div  v-for="(row, index) in rows" :key="index" class="bg-gray-50 mb-1 flex rounded-lg px-3 py-1 hover:bg-gray-100" >
      <p class="w-1/3 cursor-pointer" v-for="(value, key) in row.row" :key="key" @click="onRowClicked(row)">{{ value }} </p>
  </div>
</div>


</template>




<script>

export default {
  props : [
    "rows"
  ],
  methods: {
    onRowClicked(row){
      this.$router.push({name : row.link.name, params : row.link.params})
    }
  }
}

</script>
