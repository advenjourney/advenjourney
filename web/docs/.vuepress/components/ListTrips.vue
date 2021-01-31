<template>
  <ul>
    <li v-for="trip in trips" :key="trip">{{trip}}</li>
  </ul>
</template>
<script>
export default {
  data: () => ({
    trips: ["Trip 1", "Trip2", "Trip3"]
  })
}
</script>