import { createBrowserRouter, Navigate } from 'react-router-dom'
import { v4 as uuidv4 } from 'uuid'
import App from './App'
import HomePage from './pages/HomePage'
import WebsocketPage from './pages/WebsocketPage'
import NotFoundPage from './pages/NotFoundPage'
import SumoPage from './pages/SumoPage'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      {
        index: true,
        element: <Navigate to={`/${uuidv4()}`} replace />
      },
      {
        path: 'home',
        element: <HomePage />
      },
      {
        path: ':id',
        element: <WebsocketPage />
      },
      {
        path: 'sumo/torikumi/:division/:day',
        element: <SumoPage />
      },
      {
        path: 'not-found',
        element: <NotFoundPage />
      },
      {
        path: '*',
        element: <Navigate to="/not-found" replace />
      }
    ],
  }
])
