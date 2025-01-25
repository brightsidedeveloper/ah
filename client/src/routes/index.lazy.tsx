import { createLazyFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import { User, Users } from '../api/api'
import request from '../api/request'

export const Route = createLazyFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  useEffect(() => {
    request('http://localhost:8081/users', undefined, (b) => Users.decode(b))
  }, [])

  return (
    <div>
      <button
        onClick={async () => {
          const u = User.create({
            name: 'test',
            id: 3,
          })
          const body = User.encode(u).finish()
          request(
            'http://localhost:8081/user',
            {
              method: 'POST',
              body,
            },
            (b) => Users.decode(b)
          )
        }}
      >
        Click
      </button>
    </div>
  )
}
