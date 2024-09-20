import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.jsx";
import "./index.css";
import {
  createBrowserRouter,
  RouterProvider,
  Navigate,
} from "react-router-dom";
import Register from "./pages/auth/Register.jsx";
import Verification from "./pages/auth/Verification.jsx";
import OnBoarding from "./pages/auth/Onboarding.jsx";
import ErrPage from "./pages/ErrPage.jsx";
import Login from "./pages/auth/Login.jsx";
import ServerError from "./pages/ServerError.jsx";
import ConversationList from "./pages/conversation/ConversationList.jsx";
import Request from "./pages/conversation_request/Request.jsx";
import { Provider } from "react-redux";
import { store } from "./redux/store.js";
import SentReqList from "./pages/conversation_request/SentReqList.jsx";
import ReceivedReqList from "./pages/conversation_request/ReceivedReqList.jsx";
import LandingPage from "./pages/landing/LandingPage.jsx";
import { BrowserRouter as Router } from "react-router-dom";

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <Router>
      <Provider store={store}>
        <App></App>
      </Provider>
    </Router>
  </StrictMode>
);
