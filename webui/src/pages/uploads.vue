<template>
  <div>
    <h1>uploads</h1>
  </div>
</template>

<script>
export default {
  name: 'uploads',
  async asyncData() {
    return new Promise(resolve => {
      setTimeout(() => {
        resolve({});
      }, 2000)
    })
  },
}
</script>

<style lang="scss">
</style>
