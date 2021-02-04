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

  mounted() {
    this.apiUrl = this.$site.themeConfig.custom.url
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
      const trips = await this.axios.post(this.apiUrl, { query }, { headers: { 'Content-Type': 'application/json' } })

      this.trips = trips.data.data.offers
    }
  }
}
</script>
