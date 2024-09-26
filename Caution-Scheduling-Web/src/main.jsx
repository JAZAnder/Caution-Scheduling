import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { Outlet } from "react-router-dom";
import { AuthProvider } from './context/AuthContext'; 
import Home from './Home.jsx'; 
import ReactDOM from 'react-dom/client';
import NavigationBar from './navigationBar.jsx';
import FooterBar from './footerBar.jsx'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css';
import Login from "./routes/login/login.jsx"
import Otherlink from "./routes/otherlink/otherlink.jsx"
import LabSchedule from "./routes/labschedule/labschedule.jsx"
import ScheduleMeeting from './routes/schedulemeeting/schedulemeeting.jsx';
import SigninLab from './routes/signinlab/signinlab.jsx'
import JoinVirtually from './routes/joinvirtually/joinvirtually.jsx';



const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <NavigationBar />
        <Outlet/>
        <FooterBar/>
      </>
    ),
    children: [
      {
        path: "",
        element: <Home />
      },
      {
        path: "login",
        element: <Login />
      },
      {
        path: "otherlink",
        element: <Otherlink /> 
      },
      {
        path: "labschedule",
        element: <LabSchedule /> 
      },
      {
        path: "schedulemeeting",
        element: <ScheduleMeeting /> 
      },
      {
        path: "signinlab",
        element: <SigninLab /> 
      },
      {
        path: "joinvirtually",
        element: <JoinVirtually /> 
      }
    ]
  }
]);

ReactDOM.createRoot(document.getElementById('root')).render(
  <AuthProvider>
    <RouterProvider router={router} />
  </AuthProvider>
);