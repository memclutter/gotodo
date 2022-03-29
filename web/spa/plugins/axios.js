
export default function ({$axios, store, error}) {
  $axios.onError(e => {
    const response = e.response || {};
    const status = response.status || 0;
    const data = response.data || {};

    return Promise.resolve({
      status,
      data
    })
  })
}
