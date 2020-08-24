import Vue from "vue";
import VueWamp from "vue-wamp";


export default function NuxtWampPlugin (context, inject) {
  let url;
  if(process.env.NODE_ENV !== 'development') {
    url = ((window.location.protocol === "https:") ? "wss" : "ws") + "://" + window.location.host + "/ws/"
  } else {
    url = 'ws://localhost:4000/ws/'
  }
  console.log(url);
  Vue.use(VueWamp, {
      url,
      realm: 'gokapi',
  })
  if (!context.app['wamp']) {
    context.app['wamp'] = Vue['wamp']
  }
  if (context.store && !context.store['wamp']) {
    context.store['wamp'] = Vue['wamp']
  }
}
