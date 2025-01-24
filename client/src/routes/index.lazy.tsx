import { createLazyFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import { User } from '../api/api'

export const Route = createLazyFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  useEffect(() => {
    const user = User.create({
      id: 2,
      name: 'John Doe',
      email: '',
    })
    const binary = User.encode(user).finish()
    const decoded = User.decode(binary)
    console.log(decoded)
  }, [])

  return <div>Hello "/"!</div>
}
