export default (ctx) => {
  const urlRequiresAuth = /^\/limit(\/|$)/.test(ctx.route.fullPath)
  if (urlRequiresAuth) {
    const userIsLoggedIn = ctx.store.state.userName
    const token = ctx.store.state.userToken
    console.log(token)
    ctx.$axios.$post('api/auth', {
      token: token
    }).then(res => {
      const data = res
      if (data.status === 1) {
        ctx.store.commit('token', data.new_token)
      } else if (data.status > 10) {
        ctx.redirect('/error')
      }
    }).catch(err => console.log(err))
    if (!userIsLoggedIn) {
      return ctx.redirect('/login')
    }
  }

  return Promise.resolve()
}
