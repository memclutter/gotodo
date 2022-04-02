export default function ({route, store, redirect}) {
  const exceptedIsAuth = route.meta.length > 0 ? (route.meta[0].isAuth || false) : false;
  const actualIsAuth = store.state.auth.session.user.id !== null;

  if (exceptedIsAuth !== actualIsAuth) {
    if (actualIsAuth) {
      redirect({name: 'index'})
    } else {
      redirect({name: 'login'})
    }
  }
}
