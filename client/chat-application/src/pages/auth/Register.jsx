import { useEffect, useState } from "react"
import logo from "../../assets/logo.png"
import TextField from "../../components/TextField"
import { Link } from "react-router-dom"
import { API_Register } from "../../api/authAPI"
import LoadingButton from "../../components/LoadingButton"
import { useMemo } from "react"

function Register(){

  const [email, setEmail] = useState("")
  const [emailErrorMessage, setEmailErrorMessage] = useState("")
  const [password, setPassword] = useState("")
  const [passwordErrorMessage, setPasswordErrorMessage] = useState("")
  const [confirm, setConfirm] = useState("")
  const [confirmErrorMessage, setConfirmErrorMessage] = useState("")

  const [isRegistering, setIsRegistering] = useState(false)

  const handleRegister = ()=>{
    setIsRegistering(true)
    if (password != confirm){
      setConfirmErrorMessage("Confirmed password is not match")
      return
    }

    API_Register(email, password)
    .then(result =>{
      console.log(result)
    })
    .catch(error =>{
      console.log(error)
      let responseData = error.response.data
      if(responseData.key == "EMAIL_EXIST"){
        setEmailErrorMessage(responseData.message)
      }
      if(responseData.key == "INVALID_EMAIL"){
        setEmailErrorMessage(responseData.message)
      }
      if(responseData.key == "INVALID_PASSWORD"){
        setPasswordErrorMessage(responseData.message)
        setConfirmErrorMessage(responseData.message)
      }
    })
    .finally(()=>{
      setIsRegistering(false)
    })

  }

  useEffect(()=>{
    document.title = "Register - Chat"
  }, [])

  const isFormValid = useMemo(() => {
    return email !== "" && password !== "" && confirm !== "";
  }, [email, password, confirm]);

  return(
    <div className="flex h-screen w-screen justify-center bg-bg">
      <div className="mt-16 px-10 py-14 w-fit h-fit flex flex-col items-center rounded-2xl bg-white shadow-xl">
        <img src={logo} alt="" className="w-52" />
        <h1 className="text-5xl font-extrabold mb-10">Register</h1>
        <TextField
          title="Email"
          type="text"
          value={email}
          setValue={setEmail}
          errorMessage={emailErrorMessage}
          setErrorMessage={setEmailErrorMessage}
          id="email"
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
        <TextField
          title="Confirm"
          type="password"
          value={confirm}
          setValue={setConfirm}
          errorMessage={confirmErrorMessage}
          setErrorMessage={setConfirmErrorMessage}
          id="confirm"
        />
        <span className="self-start mb-2">
            You have an account?{" "}
            <Link to="/login" className="underline font-semibold">
              login
            </Link>
          </span>
        {isFormValid ? (
          <LoadingButton
            value="Register"
            handleFunction={handleRegister}
            disabled={false}
            isLoading={isRegistering}
          />
        ) : (
          <LoadingButton
            value="Register"
            handleFunction={handleRegister}
            disabled={true}
            isLoading={isRegistering}
          />
        )}
      </div>
    </div>
  )
}

export default Register