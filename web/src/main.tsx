import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";


import Signup from './components/signup.tsx'
import App from './App.tsx';
import Home from './components/home.tsx';

/*
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
*/

/*
const isLoggedIn = (): boolean => {
  return !!localStorage.getItem('token');
}
*/

createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <BrowserRouter>
        <Routes>
        
        <Route path="/" element={<App />} />  {/*or other routes*/}
        <Route path="/signup" element={<Signup />} />
        <Route path="/home" element={<Home />} />
        </Routes>
      </BrowserRouter>
  </StrictMode>
,
)
