import { Error as ProtobufError } from './api'

export default async function request<T>(url: string, options: RequestInit = {}, decodeFn: (binary: Uint8Array) => T): Promise<T> {
  try {
    const response = await fetch(url, options)

    const blob = await response.blob()
    const binary = new Uint8Array(await blob.arrayBuffer())

    if (!response.ok) {
      const errorMessage = ProtobufError.decode(binary)
      throw new Error(errorMessage.message || 'Unknown error occurred')
    }

    return decodeFn(binary)
  } catch (err) {
    console.error('Request error:', err)
    throw err instanceof Error ? err : new Error('An unknown error occurred')
  }
}

export function ensureError(error: unknown): Error {
  if (error instanceof Error) {
    return error
  }
  return new Error('An unknown error occurred')
}
