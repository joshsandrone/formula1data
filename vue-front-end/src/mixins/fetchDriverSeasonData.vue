<script>
import api from '@/services/api';

export default {
  data() {
    return {
        seasonData: null,
    };
  },
  methods: {
    async fetchDriverSeasonData(driverName, season) {
      try {
        let [firstname, surname] = driverName.split(' ');
        const response = await api.getDriverSeasonStats(season, firstname, surname);

        // Manipulate the data and assign it to the passed data object
        return this.manipulateSeasonData(response.data);
        
      } catch (error) {
        console.error('Failed to fetch driver season stats:', error);
      }
    },
    
    manipulateSeasonData(seasonData){
        // Manipulate the data as needed
        seasonData.extra = {};
        seasonData.extra.totalPoints = seasonData.races.points + seasonData.SprintRaces.sprintPoints;
        
        return seasonData;
    }
  }
};

</script>
