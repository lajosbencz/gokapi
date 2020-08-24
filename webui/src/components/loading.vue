<template>
  <div class="page-loading" v-if="loading">
    <div class="background"></div>
    <img class="image" src="/admin/loading.svg" alt="Loading..."/>
    <h3 class="title">loading</h3>
    <div class="placeholder"></div>
  </div>
</template>

<script>
export default {
  name: "loading",
  data: () => ({
    loading: false
  }),
  methods: {
    start() {
      this.loading = true
    },
    finish() {
      this.loading = false
    }
  },
}
</script>

<style lang="scss">
.page-loading {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  font-size: 2rem;
  display: flex;
  justify-content: center;
  flex-direction: column;
  z-index: 999999;

  > * {
    align-self: center;
  }

  .background {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 68, 204, 0.6);
  }

  .image {
    z-index: 10;
  }

  .title {
    text-transform: uppercase;
    font-weight: bold;
    color: #fff;
    z-index: 10;
  }

  .placeholder {
    display: block;
    width: 100%;
    height: 25vh;
  }
}
</style>
