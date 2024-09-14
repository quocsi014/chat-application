import { useEffect, useState } from "react";
import { getRequestSent } from "../../api/conversationRequestAPI";
import { getCookie } from "../../utils/cookie";
import defaultAvatar from "../../assets/default_avatar.png";
function SentReqList() {
  const [reqList, setReqList] = useState([]);

  useEffect(() => {
    let token = getCookie("access_token");
    getRequestSent(token)
      .then((res) => {
        console.table(res.data);
        setReqList(res.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  return (
    <div>
      {reqList.length > 0 ? (
        reqList.map((req) => {
          return (
            <div
              key={req.recipient.id}
              className="flex items-center py-2 rounded-md justify-between"
            >
              <div className="flex">
                <img
                  src={req.recipient.avatar_url || defaultAvatar}
                  alt={req.recipient.username}
                  className="w-10 h-10 object-cover rounded-full mr-2"
                />
                <div className="flex flex-col">
                  <h3 className="font-bold">
                    {req.recipient.firstname} {req.recipient.lastname}
                  </h3>
                  <p className="text-gray-500">@{req.recipient.username}</p>
                </div>
              </div>
              <button className="px-2 bg-gray-300 hover:bg-gray-500 text-lg text-white rounded-md" >cancel</button>
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
