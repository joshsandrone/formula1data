<template>
    <div>
        <div class="flex justify-center">
            <PrimeDropdown v-model="driver1" :options="drivers" class="bg-gray-50 w-30" @change="onDriverChange(leftDriver = true)" placeholder="Select a Driver" />
            <span class = "mx-20"></span>
            <PrimeDropdown v-model="driver2" :options="drivers" class="bg-gray-50 w-30" @change="onDriverChange(leftDriver = false)" placeholder="Select a Driver" />
        </div>

        <div class="flex flex-col mt-10">
            <ComparisonLine v-for="(item, index) in comparisons" :key="index" :title="item.title" :leftVal="item.leftVal" :rightVal="item.rightVal" :preferMore="item.perferMore" />
        </div>
    </div>
</template>
    
    
    
<script> 
import ComparisonLine from './Generic/ComparisonLine.vue'
import fetchDriverSeasonData from '@/mixins/fetchDriverSeasonData.vue';

export default {
    mixins: [fetchDriverSeasonData],
    components: {
        ComparisonLine
    },
    props : [
        'season',
        'drivers'
    ],
    data(){
        return{
            comparisons: [
                {
                    title: "Points",
                    id: "points",
                    perferMore: true,
                },
                {
                    title: "Race H2H",
                    id: "race-h2h",
                    perferMore: true,
                },
                {
                    title: "Qualy H2H",
                    id: "qualy-h2h",
                    perferMore: true,
                },
                {
                    title: "Wins",
                    id: "wins",
                    perferMore: true,
                },
                {
                    title: "Podiums",
                    id: "podiums",
                    perferMore: true,
                },
                {
                    title: "Average Race Pos",
                    id: "avg-race-pos",
                    perferMore: false,
                },
                {
                    title: "Average Qualy Pos",
                    id: "avg-qualy-pos",
                    perferMore: false,
                },
                {
                    title: "Highest Race Pos",
                    id: "highest-race-pos",
                    perferMore: false,
                },
                {
                    title: "Highest Qualy Pos",
                    id: "highest-qualy-pos",
                    perferMore: false,
                }
            ]
        }
    },
    methods: {
        async onDriverChange(leftDriver){
            let seasonData;

            if (leftDriver){
                seasonData = await this.fetchDriverSeasonData(this.driver1, this.season, this.leftDriverSeasonData);
            }
            else{
                seasonData = await this.fetchDriverSeasonData(this.driver2, this.season, this.rightDriverSeasonData);
            }

            this.populateComparisons(leftDriver, seasonData);
        },
        populateComparisons(leftDriver, seasonData){

            console.log("Populating comparionssssss");
            let valueKey;

            if(leftDriver){
                valueKey = "leftVal";
            }
            else{
                valueKey = "rightVal";
            }

            this.comparisons.find(r => r.id === "points")[valueKey] = seasonData.extra.totalPoints;
            this.comparisons.find(r => r.id === "wins")[valueKey] = seasonData.races.wins;
            this.comparisons.find(r => r.id === "podiums")[valueKey] = seasonData.races.podiums;
            this.comparisons.find(r => r.id === "avg-race-pos")[valueKey] = seasonData.races.avgRacePos.toFixed(2);
            this.comparisons.find(r => r.id === "avg-qualy-pos")[valueKey] = seasonData.qualifyings.avgQualyPos.toFixed(2);
            this.comparisons.find(r => r.id === "highest-race-pos")[valueKey] = seasonData.races.highestRacePos;
            this.comparisons.find(r => r.id === "highest-qualy-pos")[valueKey] = seasonData.qualifyings.highestQualyPos;
        }
    }
}

</script>
    