import {
  acceptRequest,
  getRequestReceived,
  rejectRequest,
} from "../../api/conversationRequestAPI";
import defaultAvatar from "../../assets/default_avatar.png";
import { useState, useEffect } from "react";
function ReceivedReqList() {
  const [reqList, setReqList] = useState([]);

  useEffect(() => {
    getRequestReceived()
      .then((res) => {
        setReqList(res.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  const handleAccept = (senderId) => {
    acceptRequest(senderId)
      .then(() => {
        setReqList((prevReqList) =>
          prevReqList.filter((req) => req.sender.id !== senderId)
        );
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleReject = (senderId) => {
    rejectRequest(senderId)
      .then(() => {
        setReqList((prevReqList) =>
          prevReqList.filter((req) => req.sender.id !== senderId)
        );
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <div>
      {reqList.length > 0 ? (
        reqList.map((req) => {
          return (
            <div
              key={req.sender.id}
              className="flex items-center py-2 rounded-md justify-between"
            >
              <div className="flex">
                <img
                  src={req.sender.avatar_url || defaultAvatar}
                  alt={req.sender.username}
                  className="w-10 h-10 object-cover rounded-full mr-2"
                />
                <div className="flex flex-col">
                  <h3 className="font-bold">
                    {req.sender.firstname} {req.sender.lastname}
                  </h3>
                  <p className="text-gray-500">@{req.sender.username}</p>
                </div>
              </div>
              <div>
                <button
                  onClick={(e) => {
                    handleReject(req.sender.id);
                  }}
                  className="px-2 bg-gray-300 hover:bg-gray-500 text-lg text-white rounded-md"
                >
                  reject
                </button>
                <button
                  onClick={(e) => {
                    handleAccept(req.sender.id);
                  }}
                  className="px-2 bg-blue-400 hover:bg-blue-600 ml-2 text-lg text-white rounded-md"
                >
                  accept
                </button>
              </div>
            </div>
          );
        })
      ) : (
        <span className="text-gray-300">No requests received</span>
      )}
    </div>
  );
}

export default ReceivedReqList;
