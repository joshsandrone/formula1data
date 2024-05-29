<template>
    <PrimeDropdown v-model="selectedSeason" :options="seasons" class="bg-gray-50 w-30" @change="onSeasonChange" placeholder="Select a Season" />
</template>


<script>
import api from '@/services/api';

export default {
  data() {
    return {
      seasons: [
        "2024"
      ],
      selectedSeason: "2024",
    };
  },
  mounted() {
    this.fetchAvailableSeasonProfiles();
  },
  methods: {
    async fetchAvailableSeasonProfiles(){
      try {
        const response = await api.getAvailableSeasonProfiles();
        this.setSeasonSelectOptions(response.data);
      } catch (error) {
        console.error('Failed to fetch available season profiles:', error);
      }
    },

    setSeasonSelectOptions(seasons){
      this.seasons = seasons;
    },

    onSeasonChange() {
      this.$emit(`season-selected`, this.selectedSeason);
    }
  }
}

</script>
