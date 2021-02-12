import axios from 'axios'
import VueAxios from 'vue-axios'

export default ({Vue,
    options,
    router,
    siteData,
    isServer }) => {
  Vue.use(VueAxios, axios)

  let debug = false
  if (debug) {
    console.log(Vue)
    console.log(options)
    console.log(router)
    console.log(siteData)
    console.log(isServer)
  }
};