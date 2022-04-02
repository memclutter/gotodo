import {camelizeKeys} from 'humps'

export default function ({$axios, store, error}) {
  $axios.onRequest(req => {

  })

  $axios.onResponse(res => {
    res.succes = true;
    res.data = camelizeKeys(res.data)
    return Promise.resolve(res);
  });

  $axios.onError(e => {
    const response = e.response || {};
    const status = response.status || 0;
    const data = response.data || {};

    return Promise.resolve({
      success: false,
      status,
      data
    })
  });
}
