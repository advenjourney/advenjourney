<template>
  <ul>
    <li v-for="trip in trips" :key="trip.title" v-text="trip.title" />
  </ul>
</template>
<script>
export default {
  data: () => ({
    trips: []
  }),

  created() {
    this.getTrips()
  },

  methods: {
    async getTrips() {
      const query = `
        query {
          offers{
            title
            location,
            description,
            titleImageUrl,
          }
        }
        `
      const trips = await this.axios.post('http://localhost:9000/query', { query }, { headers: { 'Content-Type': 'application/json' } })

      this.trips = trips.data.data.offers
    }
  }
}
</script>
