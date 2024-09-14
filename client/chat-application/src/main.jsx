import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.jsx";
import "./index.css";
import { createBrowserRouter, RouterProvider, Navigate } from "react-router-dom";
import Register from "./pages/Register.jsx";
import Verification from "./pages/Verification.jsx";
import OnBoarding from "./pages/Onboarding.jsx";
import ErrPage from "./pages/ErrPage.jsx";
import Login from "./pages/Login.jsx";
import ServerError from "./pages/ServerError.jsx";
import ConversationList from "./layouts/ConversationList.jsx";
import Request from "./layouts/conversation_request/Request.jsx";
import { Provider } from "react-redux";
import { store } from "./redux/store.js";
import SentReqList from "./layouts/conversation_request/SentReqList.jsx";
import ReceivedReqList from "./layouts/conversation_request/ReceivedReqList.jsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrPage />,
    children: [
      {
        path: "/conversations",
        element: <ConversationList />,
      },
      {
        path: "/requests",
        element: <Request />,
        children: [
          {
            index: true,
            element: <Navigate to="/requests/received" replace />,
          },
          {
            path: "sent",
            element: <SentReqList />,
          },
          {
            path: "received",
            element: <ReceivedReqList />,
          },
        ],
      },
    ],
  },
  {
    path: "/register",
    element: <Register />,
  },
  {
    path: "/verify",
    element: <Verification />,
  },
  {
    path: "/onboarding",
    element: <OnBoarding />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/500",
    element: <ServerError />,
  },
]);

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <Provider store={store}>
      <RouterProvider router={router} />
    </Provider>
  </StrictMode>
);
