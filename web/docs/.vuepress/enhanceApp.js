import axios from 'axios'
import VueAxios from 'vue-axios'

export default ({ Vue }) => {
  Vue.use(VueAxios, axios)
};