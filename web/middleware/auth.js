export default ({ store, route, redirect }) => {
  if (store.isLogin) {
    const token = store.userToken
    this.$axios.post('/api/auth', {
      token
    }).then(res => {
      const { data } = res
      // todo: token reset有り得る
      if (!data.error_text && data.checks[0].status === 'VALID') {
        redirect('/secret')
      } else {
        redirect('/')
      }
      // todo: end
      redirect('/')
    }).catch(err => console.log(err))
  } else {
    redirect('/')
  }
}
