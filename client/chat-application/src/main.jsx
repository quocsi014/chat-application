import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.jsx'
import "./index.css"
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Register from './pages/Register.jsx'
import Verification from './pages/Verification.jsx'
import OnBoarding from './pages/Onboarding.jsx'
import ErrPage from './pages/ErrPage.jsx'
import Login from './pages/Login.jsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>,
    errorElement: <ErrPage/>
  },
  {
    path:"/register",
    element: <Register/>,
  },
  {
    path: "/verify",
    element: <Verification/>
  },
  {
    path: "/onboarding",
    element: <OnBoarding/>
  },
  {
    path:"/login",
    element: <Login/>
  }
]);

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router}/>
  </StrictMode>,
)
