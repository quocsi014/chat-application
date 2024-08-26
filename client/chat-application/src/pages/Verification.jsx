import { useEffect, useState } from "react";
import LoadingButton from "../components/LoadingButton";
import { Link, useLocation, useNavigate, useParams } from "react-router-dom";
import { API_Verify } from "../api/authAPI";
import { setCookie } from "../utils/cookie";

function Verification() {
  const [isVerifying, setIsVerifying] = useState(false);
  const [showError, setShowError] = useState(false)
  const navigate = useNavigate()
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);

  const token = queryParams.get("token");

  useEffect(() => {
    document.title = "Verify - Chat";
  }, []);

  const handleVerify = () => {
    setIsVerifying(true);
    API_Verify(token)
      .then((result) => {
        setCookie("access_token", result.data.token, 30);
        navigate("/onboarding")
      })
      .catch((error) => {
        setShowError(true)
      })
      .finally(() => {
        setIsVerifying(false);
      });
  };

  return (
    <div className="flex h-screen w-screen justify-center bg-bg">
      <div className="bg-white h-fit px-10 py-10 rounded-xl mt-10 w-96 box-content">
        <h2 className="mb-4">Click bellow button to verify your account</h2>
        <LoadingButton value="Verify" className="w-24" handleFunction={handleVerify} isLoading={isVerifying} disabled={false} />
        <div className={`${showError? "block":"hidden"}`}>
          <h2 className="mt-4 text-red-500">ERROR</h2>
          <p>
            Verification is invalid or expired, to continue please{" "}
            <Link to="/register" className="underline font-semibold">
              register
            </Link>{" "}
            again.
          </p>
          <p>
            Or{" "}
            <Link to="/login" className="underline font-semibold">
              Login
            </Link>{" "}
            if you already have an account
          </p>
        </div>
      </div>
    </div>
  );
}

export default Verification;
