export default async function request(url: string, options: RequestInit = {}) {
  return fetch(url, options)
    .then(async (response) => {
      if (!response.ok) {
        throw new Error(response.statusText)
      }
      const blob = await response.blob()
      return new Uint8Array(await blob.arrayBuffer())
    })
    .catch((err) => {
      console.error(err)
      return new Error('Oops')
    })
}
