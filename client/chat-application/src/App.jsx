import { useRoutes } from "react-router-dom";
import LandingPage from "./pages/landing/LandingPage";
import Login from "./pages/auth/Login";
import Register from "./pages/auth/Register";
import Verification from "./pages/auth/Verification";
import SentReqList from "./pages/conversation_request/SentReqList";
import ReceivedReqList from "./pages/conversation_request/ReceivedReqList";
import AppLayout from "./layouts/AppLayout";
import ConversationList from "./pages/conversation/ConversationList";
import Request from "./pages/conversation_request/Request";
import SettingPage from "./pages/setting-page/SettingPage";

function App() {
  const element = useRoutes([
    {
      path: "/",
      element: <LandingPage />,
    },
    {
      path: "/login",
      element: <Login />,
    },
    {
      path: "Register",
      element: <Register />,
    },
    {
      path: "/verify",
      element: <Verification />,
    },
    {
      path: "/",
      element: <AppLayout />,
      children: [
        {
          path: "conversations",
          element: <ConversationList />,
        },
        {
          path: "requests",
          element: <Request />,
          children: [
            {
              path: "sent",
              element: <SentReqList />,
            },
            {
              path: "received",
              element: <ReceivedReqList />,
            }
          ],
        },
        {
          path: 'settings',
          element: <SettingPage/>
        }
      ],
    },
  ]);
  return element;
}

export default App;
