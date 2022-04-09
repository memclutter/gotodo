export default function (error) {
  console.log('apis interceptors request rejected', error)
  return Promise.reject(error)
}
