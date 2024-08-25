import { useEffect, useState } from "react";
import TextField from "../components/TextField";
import { API_Login } from "../api/authAPI";
import { Link, useNavigate } from "react-router-dom";
import logo from "../assets/logo.png";
import { setCookie } from "../utils/cookie";

function Login() {
  const [account, setAccount] = useState("");
  const [accountErrorMessage, setAccountErrorMessage] = useState("");
  const [password, setPassword] = useState("");
  const [passwordErrorMessage, setPasswordErrorMessage] = useState("");

  const BLANK_FIELD_QUOTE = "This field can not be blank";
  let navigate = useNavigate();

  useEffect(()=>{
    document.title = "Login - Chat"
  }, [])

  const handleLogin = () => {
    API_Login(account, password)
      .then((result) => {
        setCookie("access_token", result.data.token, 30);
        navigate("/");
      })
      .catch((error) => {
        if (error.status == 401) {
          setAccountErrorMessage(error.response.data.message);
          setPasswordErrorMessage(error.response.data.message);
        }
      });
  };

  const handleLoginWithKey = (event) => {
    if (event.key == "Enter") {
      if (account == "") {
        setAccountErrorMessage(BLANK_FIELD_QUOTE);
        return;
      }
      if (password == "") {
        setPasswordErrorMessage(BLANK_FIELD_QUOTE);
        return;
      }
      handleLogin();
    }
  };

  return (
    <div
      className="flex h-screen w-screen justify-center bg-bg"
      onKeyDown={(e) => {
        handleLoginWithKey(e);
      }}
    >
      <div className="mt-20 px-10 py-14 w-fit h-fit flex flex-col items-center rounded-2xl bg-white shadow-xl">
        <img src={logo} alt="" className="w-52" />
        <h1 className="font-extrabold text-5xl mb-10">Login</h1>
        <TextField
          title="Email or username"
          type="text"
          value={account}
          setValue={setAccount}
          errorMessage={accountErrorMessage}
          setErrorMessage={setAccountErrorMessage}
          id="email_or_password"
        />
        <TextField
          title="Password"
          type="password"
          value={password}
          setValue={setPassword}
          errorMessage={passwordErrorMessage}
          setErrorMessage={setPasswordErrorMessage}
          id="password"
        />
        <div className="flex flex-col items-start w-full mb-8">
          <span>
            You don&apos;t have an account yet?{" "}
            <Link to="/register" className="underline font-semibold">
              register
            </Link>
          </span>
          <a href="/forgotpassword" className="underline font-semibold">
            Forgot password
          </a>
        </div>
        {account != "" && password != "" ? (
          <button
            className="w-full bg-blue-500 py-2 text-xl font-semibold text-white rounded-xl"
            onClick={() => {
              handleLogin();
            }}
          >
            Login
          </button>
        ) : (
          <button
            className="w-full bg-blue-200 py-2 text-xl font-semibold text-white rounded-xl"
            onClick={() => {
              handleLogin();
            }}
            disabled
          >
            Login
          </button>
        )}
      </div>
    </div>
  );
}

export default Login;
