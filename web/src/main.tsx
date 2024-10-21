import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import ErrorPage from './errors/error-page.tsx'
import Chat from './routes/chat.tsx'
import './index.css'
import {
  createBrowserRouter,
  BrowserRouter,
  RouterProvider,
  Routes,
  Route,
} from "react-router-dom";
import * as reactRouterDom from "react-router-dom";


import Navbar from './components/navbar.tsx'
import Root from './routes/root.tsx';
import Signup from './components/signup.tsx'


const router = createBrowserRouter([
  
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/signup",
    element: < Signup />,
  }
]);


const isLoggedIn = (): boolean => {
  return !!localStorage.getItem('token');
}

createRoot(document.getElementById('root')!).render(

      <BrowserRouter>
        <Routes>
        
        <Route path="/" element={<Root />} />  {/*or other routes*/}
        <Route path="/signup" element={<Signup />} />
        </Routes>
      </BrowserRouter>
,
)
