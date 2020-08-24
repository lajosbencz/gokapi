<script>
export default {
  name: "types-id-edit",
  data: () => ({
    form: {},
  }),
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.form = Object.assign({}, vm.$store.state.types.list.find(i => i.id === vm.id));
    })
  },
  computed: {
    id() {
      return this.$route.params.id
    },
  },
  methods: {
    edit() {
      this.$store.commit('types/edit', {id: this.id, fields: this.form})
      this.$router.push({name: 'types'})
    },
  },
}
</script>

<template>
  <div>
    <h3>Create new Type</h3>
    <b-form-group>
      <label>Name</label>
      <b-input-group>
        <b-input type="text" v-model="form.name"/>
      </b-input-group>
    </b-form-group>
    <b-form-group class="text-right">
      <b-btn @click="edit" variant="primary">
        Save
      </b-btn>
    </b-form-group>
  </div>
</template>

<style lang="scss">

</style>
