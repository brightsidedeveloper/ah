import { createLazyFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import { Users } from '../api/api'
import request, { ensureError } from '../api/request'

export const Route = createLazyFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  useEffect(() => {
    request('http://localhost:8081/users', undefined, (b) => Users.decode(b))
      .then((users) => {
        console.log(users)
      })
      .catch((v) => {
        const err = ensureError(v)
        console.log(err.message)
      })
  }, [])

  return <div></div>
}
