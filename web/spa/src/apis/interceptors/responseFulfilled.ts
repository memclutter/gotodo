export default function (response) {
  console.log('apis interceptors response fulfilled', response)
  return {
    success: true,
    data: response.data
  }
}
