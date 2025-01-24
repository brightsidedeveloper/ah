import { createLazyFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import { User, Users } from '../api/api'
import request from '../api/request'

export const Route = createLazyFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  useEffect(() => {
    request('http://localhost:8081/users').then((binary) => {
      if (binary instanceof Error) {
        console.error(binary)
        return
      }
      const users = Users.decode(binary)
      console.log(users)
    })
  }, [])

  return (
    <div>
      <button
        onClick={() => {
          const user = User.create({ id: 3, name: 'Alice' })
          const body = User.encode(user).finish()

          request('http://localhost:8081/user', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/x-protobuf',
            },
            body,
          }).then((binary) => {
            if (binary instanceof Error) {
              console.error(binary)
              return
            }
            const users = Users.decode(binary)
            console.log(users)
          })
        }}
      >
        Click Me!
      </button>
    </div>
  )
}
