<template>
  <div>
    <h1>content</h1>
  </div>
</template>

<script>
export default {
  name: 'content',
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
