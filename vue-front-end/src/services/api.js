import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:7999', // Update with your API base URL
  headers: {
    'Content-Type': 'application/json'
  }
});

export default {
  getAvailableSeasonProfiles(){
    return apiClient.get(`availableSeasonProfiles`)
  },
  getSeasonProfile(season){
    return apiClient.get(`/seasonTeamDriverProfile/${season}`)
  },
  getDriverSeasonStats(season, firstname, surname){
    return apiClient.get(`/driverSeasonData/${season}/${firstname}/${surname}`)
  }
};
