import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://10.41.66.123:7999',
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
