import http from 'axios'

const ip = 'http://127.0.0.1:8080'
const url = ip + '/v1/db/index/index'
const url_kv = ip + '/v1/db/kv'

export default {
  getIndex() {
    return http.get(url)
  },
  changeIndex(args) {
    return http.post(url_kv, args)
  }
}
