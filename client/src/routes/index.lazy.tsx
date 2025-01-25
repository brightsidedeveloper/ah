import { createLazyFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import { User, Users } from '../api/api'
import { get, post } from '../api/request'

export const Route = createLazyFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  useEffect(() => {
    get('/api/users', undefined, (b) => Users.decode(b))
  }, [])

  return (
    <div>
      <button
        onClick={async () => {
          post(
            '/api/user',
            User.create({
              name: 'test',
              id: 3,
            }),
            (u) => User.encode(u).finish(),
            (b) => Users.decode(b)
          )
        }}
      >
        Click
      </button>
    </div>
  )
}
