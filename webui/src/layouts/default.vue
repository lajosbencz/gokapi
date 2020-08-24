<script>
import Navbar from "./../components/navbar";

export default {
  components: {
    Navbar,
  },
  computed: {
    keepAliveExclude: () => ([
      // 'types-create',
      'types-id-edit',
    ]),
  },
}
</script>

<template>
  <div class="layout-default">
    <navbar/>
    <div class="container-fluid">
      <nuxt keep-alive :keep-alive-props="{ exclude: keepAliveExclude }"/>
    </div>
  </div>
</template>

<style lang="scss">
</style>
