import { Error as ProtobufError } from './api'
import debug from 'debug'

const loggers = {
  GET: debug('request:GET'),
  POST: debug('request:POST'),
  PUT: debug('request:PUT'),
  DELETE: debug('request:DELETE'),
  PATCH: debug('request:PATCH'),
}

export default async function request<T>(url: string, options: RequestInit = {}, decodeFn?: (binary: Uint8Array) => T): Promise<T> {
  const method = options.method ?? 'GET'
  const log = loggers[method as keyof typeof loggers]
  try {
    const response = await fetch(url, options)

    const blob = await response.blob()
    const binary = new Uint8Array(await blob.arrayBuffer())

    if (!response.ok) {
      const errorMessage = ProtobufError.decode(binary)
      throw new Error(errorMessage.message || 'Unknown error occurred')
    }
    if (!decodeFn) {
      return {} as unknown as T
    }
    const res = decodeFn(binary)
    log(url, res)
    return res
  } catch (err) {
    log('Error: ', options.method ?? 'GET', url, err)
    throw err instanceof Error ? err : new Error('An unknown error occurred')
  }
}

export function ensureError(error: unknown): Error {
  if (error instanceof Error) {
    return error
  }
  return new Error('An unknown error occurred')
}
