<script>
const formDefaults = {
  name: '',
  fields: [],
};
export default {
  name: "types-create",
  data: () => ({
    form: {...formDefaults},
  }),
  methods: {
    create() {
      this.$store.commit('types/create', this.form)
      const id = this.form.id;
      this.reset();
      this.$router.push({name: 'types-id-edit', params: {id}});
    },
    reset() {
      this.$set(this, 'form', {...formDefaults});
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
      <b-btn @click="create" variant="primary">
        Create
      </b-btn>
    </b-form-group>
  </div>
</template>

<style lang="scss">
</style>
