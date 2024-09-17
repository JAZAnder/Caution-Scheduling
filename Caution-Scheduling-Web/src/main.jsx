import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import Home from './Home.jsx'; 
import ReactDOM from 'react-dom/client'
import NavigationBar from './navigationBar.jsx';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css';
import LoginTestButtons from "./routes/testing/login.jsx"

const router = createBrowserRouter([
  {
    path: "/",
    element: <NavigationBar/>,
    children:[
      {
        path: "",
        element: <Home/>
      }
    ]
  },{
    path: "testing",
    children:[
      {
        path:"login",
        element: <LoginTestButtons/>
      }
    ]
  }
]);

ReactDOM.createRoot(document.getElementById("root")).render(
    <RouterProvider router={router} />
)
