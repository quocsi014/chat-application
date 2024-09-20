import { useEffect, useRef, useState } from "react";
import TextField from "../../components/TextField";
import LoadingButton from "../../components/LoadingButton";
import defaultAvatar from "../../assets/default_avatar.png";
import { BiSolidEdit } from "react-icons/bi";
import axios from "axios";
import { createUser } from "../../api/userAPI";
import { getCookie } from "../../utils/cookie";
import { useNavigate } from "react-router-dom";

function OnBoarding() {
  const [firstname, setFirstname] = useState("");
  const [firstnameErrorMessage, setFirstnameErrorMessage] = useState("");
  const [lastname, setLastname] = useState("");
  const [lastnameErrorMessage, setLastnameErrorMessage] = useState("");
  const [avatar, setAvatar] = useState(null);
  const [preview, setPreview] = useState(defaultAvatar);
  const [isCreating, setIsCreating] = useState(false);
  const [dialogMessage, setDialogMessage] = useState("");
  const [username, setUsername] = useState("");
  const [usernameErrorMessage, setUsernameErrorMessage] = useState("");

  const dialogRef = useRef(null)
  const closeDialogButtonRef = useRef(null)

  const showDialog = (doAfterClose)=>{
    dialogRef.current.showModal()
    closeDialogButtonRef.current.addEventListener('click', ()=>{
      dialogRef.current.close()
      if(doAfterClose){
        doAfterClose()
      }
    })
  }

  const navigate = useNavigate();
  const token = getCookie("access_token");

  useEffect(() => {
    document.title = "Initialize - Chat";
  }, []);

  const handleChooseAvatar = (event) => {
    const file = event.target.files[0];
    if (file) {
      if (file.type.startsWith("image/")) {
        setAvatar(file);
        setPreview(URL.createObjectURL(file));
      } else {
        alert("You must choose an image");
      }
    }
  };

  function CallCreateUserAPI(avatar_url) {
    createUser(firstname, lastname, username, avatar_url, token)
      .then((res) => {
        navigate("/");
      })
      .catch((err) => {
        if (err.response.data.key == "INITIALIZED") {
          setDialogMessage(err.response.data.message + "\nYou will be redirected to the app.")
          showDialog(()=>{navigate("/")})
          return;
        }
        if (err.response.data.key == "INVALID_FIRSTNAME") {
          setFirstnameErrorMessage(err.response.data.message);
          return;
        }
        if (err.response.data.key == "INVALID_LASTNAME") {
          setLastnameErrorMessage(err.response.data.message);
          return;
        }
        if (err.response.data.key == "INVALID_USERNAME") {
          setUsernameErrorMessage(err.response.data.message);
          return;
        }
        if (err.response.data.key == "USERNAME_TAKEN") {
          setUsernameErrorMessage(err.response.data.message);
          return;
        }
      })
      .finally(() => {
        setIsCreating(false);
      });
  }

  const handleCreateInformation = () => {
    setIsCreating(true);

    if (avatar == null) {
      CallCreateUserAPI(null);
    } else {
      const formData = new FormData();
      formData.append("file", avatar);
      formData.append("upload_preset", "rpekp7w3");
      axios
        .post(
          `https://api.cloudinary.com/v1_1/dobwiw6lm/image/upload`,
          formData
        )
        .then((response) => {
          CallCreateUserAPI(response.data.secure_url);
        })
        .catch((err) => {
          alert(
            "Fail to upload avatar, pls come back later or you can create without avatar"
          );
        })
        .finally(() => {
          setIsCreating(false);
        });
    }
  };

  return (
    <div className="flex h-screen w-screen justify-center bg-bg">
      <div className="mt-20 px-10 py-14 w-fit h-fit flex flex-col items-center rounded-2xl bg-white shadow-xl">
        <h1 className="font-extrabold text-5xl mb-10">Initialize</h1>
        <div className="relative mb-6">
          <img
            src={preview}
            className="w-52 h-52 rounded-full border-4  border-gray-400 border-dashed"
            alt=""
          />
          <label
            className="absolute bottom-0 translate-y-1/3 left-1/2 -translate-x-1/2 bg-white rounded-full p-1"
            htmlFor="avatar"
          >
            <BiSolidEdit />
          </label>
          <input
            type="file"
            className="hidden"
            id="avatar"
            accept="image/*"
            onChange={(e) => {
              handleChooseAvatar(e);
            }}
          />
        </div>
        <TextField
          title="Username"
          value={username}
          setValue={setUsername}
          errorMessage={usernameErrorMessage}
          setErrorMessage={setUsernameErrorMessage}
          type="text"
          id="username"
        />
        <TextField
          title="Firstname"
          value={firstname}
          setValue={setFirstname}
          errorMessage={firstnameErrorMessage}
          setErrorMessage={setFirstnameErrorMessage}
          type="text"
          id="firstname"
        />
        <TextField
          title="Lastname"
          value={lastname}
          setValue={setLastname}
          errorMessage={lastnameErrorMessage}
          setErrorMessage={setLastnameErrorMessage}
          type="text"
          id="lastname"
        />
        {firstname != "" && lastname != "" ? (
          <LoadingButton
            value="Create"
            handleFunction={handleCreateInformation}
            disabled={false}
            isLoading={isCreating}
          />
        ) : (
          <LoadingButton
            value="Create"
            handleFunction={handleCreateInformation}
            disabled={true}
            isLoading={isCreating}
          />
        )}
      </div>
      <dialog ref={dialogRef} className="bg-white rounded-lg z-50 p-5" id="myDialog">
        <h2>Error</h2>
        <p>{dialogMessage}</p>
        <button ref={closeDialogButtonRef} className="outline-none bg-gray-300 p-2">Close</button>
      </dialog>
    </div>
  );
}

export default OnBoarding;
