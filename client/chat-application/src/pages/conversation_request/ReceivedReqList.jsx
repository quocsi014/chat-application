import {
  acceptRequest,
  getRequestReceived,
  rejectRequest,
} from "../../api/conversationRequestAPI";
import defaultAvatar from "../../assets/default_avatar.png";
import { useState, useEffect } from "react";
import BriefUser from "../../components/BriefUser";
function ReceivedReqList() {
  const [reqList, setReqList] = useState([]);

  useEffect(() => {
    getRequestReceived()
      .then((res) => {
        setReqList(res.data.items);
        console.log(res.data.items)
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
        alert("Fail, something went wrong")
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
        alert("Fail, something went wrong")
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
              <BriefUser user={req.sender} />
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
