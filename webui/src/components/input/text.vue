<template>
  <b-input type="text" :value="value" @input="v => $emit('input', v)"/>
</template>

<script>
export default {
  name: 'input-text',
  props: {
    value: {
      type: String,
    },
  },
  methods: {
    log: console.log,
  }
}
</script>

<style lang="scss">

</style>
