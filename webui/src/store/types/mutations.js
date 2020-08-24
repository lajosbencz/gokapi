
import Vue from 'vue'

import state from './state'

let id = state().list.length

export default {
  create(state, type) {
    type.id = id
    state.list.push(type)
    id++
  },
  delete(state, typeId) {
    state.list = state.list.filter(i => i.id !== typeId)
  },
  edit(state, {id, fields}) {
    const i = state.list.findIndex(i => i.id === id);
    //console.log({id, list: state.list, i});
    Vue.set(state.list, i, Object.assign({}, state.list[i], fields));
    //state.list[i] = Object.assign({}, state.list[i], fields)
  },
}
