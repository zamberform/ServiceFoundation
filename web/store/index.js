export const state = () => ({
  userName: null,
  userToken: null,
  isLogin: false,
  isLimit: false
})

export const mutations = {
  user(state, userName) {
    state.userName = userName
  },
  token(state, token) {
    state.token = token
  },
  limit(state, limit) {
    state.limit = limit
  },
  setLoginState(state, isLogin) {
    state.isLogin = isLogin
  }
}
