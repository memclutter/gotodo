export default function (error: any) {
  console.log('apis interceptors request rejected', error)
  return Promise.reject(error)
}
