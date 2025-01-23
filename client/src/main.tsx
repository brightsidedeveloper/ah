import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'

import { routeTree } from './routeTree.gen'
import { createRouter, RouterProvider } from '@tanstack/react-router'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

const qc = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
    },
  },
})

// Create a new router instance
const r = createRouter({ routeTree, context: { queryClient: qc } })

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof r
  }
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={qc}>
      <RouterProvider router={r} />
    </QueryClientProvider>
  </StrictMode>
)
