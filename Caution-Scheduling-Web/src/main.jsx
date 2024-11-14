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
import Labschedule from "./routes/labschedule/labschedule.jsx"
import Schedulemeeting from './routes/schedulemeeting/schedulemeeting.jsx';
import Signinlab from './routes/signinlab/signinlab.jsx'
import Aboutus from './routes/aboutus/aboutus.jsx'
import Users from './routes/users/users.jsx'
import Labs from './routes/labs/labs.jsx'
import Timeslots from './routes/timeslots/timeslots.jsx'
import Meetings from './routes/my-meetings/my-meetings.jsx'
import MyMeetings from './routes/my-meetings/my-meetings.jsx'
import AdminMeetings from './routes/admin/meetings/meetings.jsx';
import MyAvailability from './routes/my-availability/my-availability.jsx'
import UserTimeslots from './routes/usertimeslots/usertimeslots'
import TimeslotManagement from './routes/tutortimeslots/tutortimeslots.jsx'


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
        element: <Labschedule /> 
      },
      {
        path: "schedulemeeting",
        element: <Schedulemeeting /> 
      },
      {
        path: "signinlab",
        element: <Signinlab /> 
      },
      {
        path: "aboutus",
        element: <Aboutus /> 
      },
      {
        path: "meetings",
        element: <Meetings />
      },
      {
        path: "users",
        element: <Users />
      },
      {
        path: "labs",
        element: <Labs />
      },
      {
        path: "timeslots",
        element: <Timeslots />
      },
      {
        path: "timeslotmanagement",
        element: <TimeslotManagement />
      },
      {
        path: "meetings",
        element: <MyMeetings />
      },
      {
        path: "my-availability",
        element: <MyAvailability />
      },{
        path: "user-timeslots",
        element: <UserTimeslots/>
      },{
        path: "admin",
        children:[
          {
            path: "adminmeetings",
            element: <AdminMeetings/>
        }
      ]
      }
    ]
  }
]);

ReactDOM.createRoot(document.getElementById('root')).render(
  <AuthProvider>
    <RouterProvider router={router} />
  </AuthProvider>
);