import { useEffect, useState } from "react";
import {
  deleteRequest,
  getRequestSent,
} from "../../api/conversationRequestAPI";
import { getCookie } from "../../utils/cookie";
import defaultAvatar from "../../assets/default_avatar.png";
import Declaration from "postcss/lib/declaration";
import BriefUser from "../../components/BriefUser";
function SentReqList() {
  const [reqList, setReqList] = useState([]);

  useEffect(() => {
    let token = getCookie("access_token");
    getRequestSent(token)
      .then((res) => {
        setReqList(res.data.items);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  const handleCancelRequest = (recipientId) => {
    console.log(recipientId);
    deleteRequest(recipientId)
      .then((res) => {
        let newReqList = reqList.filter((req) => {
          return req.recipient.id != recipientId;
        });
        setReqList(newReqList);
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
              key={req.recipient.id}
              className="flex items-center py-2 rounded-md justify-between"
            >
              <BriefUser user={req.recipient}/>
              <button
                onClick={() => {
                  handleCancelRequest(req.recipient.id);
                }}
                className="px-2 bg-gray-300 hover:bg-gray-500 text-lg text-white rounded-md"
              >
                cancel
              </button>
            </div>
          );
        })
      ) : (
        <span className="text-gray-300">No requests sent</span>
      )}
    </div>
  );
}

export default SentReqList;
