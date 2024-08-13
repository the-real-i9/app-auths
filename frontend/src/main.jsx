import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import App from './App.jsx'
import './index.css'
import Home from './components/Home.jsx'
import { homeLoader, socialAuthLoader } from './loaders/loader.js'
import SocialAuths from './components/SocialAuths.jsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        index: true,
        element: <Home />,
        loader: homeLoader,
      },
      {
        path: "/auth",
        element: <SocialAuths />,
        loader: socialAuthLoader,
      },
      {
        path: "/auth/:provider/oauth_callback",
        loader: null,
      },
    ],
  },
]);

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
