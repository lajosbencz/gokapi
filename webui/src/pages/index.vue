<template>
  <div class="container">
    <div>
      <b-input-group>
        <b-input-group-prepend>
          {{ time }}
        </b-input-group-prepend>
        <b-btn @click="getTime">
          Get Current Time
        </b-btn>
      </b-input-group>
    </div>
  </div>
</template>

<script>
export default {
  name: 'home',
  data() {
    return {
      time: '',
    }
  },
  methods: {
    async getTime() {
      try {
        this.time = await this.$wamp.call('worldtime');
      }
      catch(e) {
        this.time = e.toString()
      }
    },
  },
}
</script>

<style>
</style>
