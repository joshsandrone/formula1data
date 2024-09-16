<template>
    <div>
        <div class="p-2 max-w-screen-sm mx-auto flex mb-2" :id="comparisonId">
            <div class="flex justify-between w-full">
                <p :class="leftClass" class="px-4 py-2 rounded-md left-comparison-elem">{{ leftVal }}</p>
                <p class="mx-20 py-2 font-medium">{{ title }}</p>
                <p :class="rightClass" class="px-4 py-2 rounded-md right-comparison-elem">{{ rightVal }}</p>
            </div>
        </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      title: String,
      leftVal: Number,
      rightVal: Number,
      preferMore: Boolean
    },
    computed: {
      comparisonId() {
        return `comparison-${this.title}`;
      },
      leftClass() {
        return this.getClass(this.leftVal, this.rightVal, this.preferMore, true);
      },
      rightClass() {
        return this.getClass(this.leftVal, this.rightVal, this.preferMore, false);
      }
    },
    watch: {
      leftVal(newValue, oldValue) {
        console.log("NEW LEFT VAL!!!!!!!!")
        if (newValue !== oldValue) {
          this.updateClasses();
        }
      },
      rightVal(newValue, oldValue) {
        console.log("NEW RIGHT VALLLLLLL")
        if (newValue !== oldValue) {
          this.updateClasses();
        }
      }
    },
    methods: {
      getClass(val1, val2, preferMore, isLeft) {
        let leftCondition = isLeft ? (preferMore ? val1 > val2 : val1 < val2) : (preferMore ? val1 < val2 : val1 > val2);
        let rightCondition = isLeft ? (preferMore ? val1 < val2 : val1 > val2) : (preferMore ? val1 > val2 : val1 < val2);
  
        if (leftCondition) {
          return `bg-green-200 px-4 py-2 rounded-md ${isLeft ? 'left-comparison-elem' : 'right-comparison-elem'}`;
        } else if (rightCondition) {
          return `bg-red-400 px-4 py-2 rounded-md ${isLeft ? 'left-comparison-elem' : 'right-comparison-elem'}`;
        } else {
          return `bg-orange-200 px-4 py-2 rounded-md ${isLeft ? 'left-comparison-elem' : 'right-comparison-elem'}`;
        }
      },
      updateClasses() {
        // Force update the computed properties
        this.$forceUpdate();
      }
    }
  };
  </script>

  