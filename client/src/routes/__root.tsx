import { QueryClient } from '@tanstack/react-query'
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'

export const Route = createRootRouteWithContext<QueryContext>()({
  component: RootComponent,
})

function RootComponent() {
  return <Outlet />
}

export type QueryContext = {
  queryClient: QueryClient
}
