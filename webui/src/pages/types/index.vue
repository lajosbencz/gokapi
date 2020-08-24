<script>
import InputText from "./../../components/input/text";

export default {
  name: 'types',
  components: {InputText},
  data() {
    return {
      txt: 'foo',
    };
  },
  methods: {
    deleteType(id) {
      this.$store.commit('types/delete', id)
    },
    editType(id) {
      this.$router.push({name: 'types-id-edit', params: {id}})
    },
  },
}
</script>

<template>
  <div>
    <div class="text-right">
      <b-input type="text"/>
      <b-btn :to="{name:'types-create'}" variant="success">
        <b-icon-plus-square/>
        New Type
      </b-btn>
    </div>
    <table class="table table-borderless table-striped">
      <tbody>
      <tr v-for="type in $store.state.types.list">
        <th>{{ type.name }}</th>
        <td class="text-right">
          <b-btn-group>
            <b-btn size="sm"
                   variant="primary"
                   @click="editType(type.id)"
                   v-b-tooltip
                   title="Edit Type">
              <b-icon-pencil/>
            </b-btn>
            <b-btn size="sm"
                   variant="danger"
                   @click="deleteType(type.id)"
                   v-b-tooltip
                   title="Delete Type">
              <b-icon-x/>
            </b-btn>
          </b-btn-group>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<style lang="scss">
</style>
