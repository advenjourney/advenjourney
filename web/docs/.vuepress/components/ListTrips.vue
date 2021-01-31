<template>
  <ul>
    <li v-for="trip in trips" :key="trip">{{trip}}</li>
  </ul>
</template>
<script>
import gql from 'graphql-tag'

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
      const trips = await this.axios.post('http://localhost:9000/query', {
        query
      })

      console.log(trips)
    }
  }
}
</script>
