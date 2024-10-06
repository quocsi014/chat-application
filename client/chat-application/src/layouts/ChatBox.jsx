import { useEffect, useState } from "react";
import { MdArrowBackIos, MdArrowForwardIos, MdSend } from "react-icons/md";
import { useSelector } from "react-redux";
import { getMessages } from "../api/messageAPI";
import Message from "../components/Message";

function ChatBox() {
  const [isChatInfoOpen, setIsChatInfoOpen] = useState(true);
  const [message, setMessage] = useState("");
  const handleCloseChatInfo = () => {
    setIsChatInfoOpen(!isChatInfoOpen);
  };
  const [messageList, setMessageList] = useState([]);
  const chat = useSelector((state) => state.currentChat.chat);
  const me = useSelector((state) => state.me.me);

  useEffect(() => {
    if (!chat) return;
    getMessages(chat.id)
      .then((res) => {
        setMessageList(res.data.items);
      })
      .catch((error) => {
        console.log(error);
      });
  }, [chat]);

  const sectionPadding = "p-5";

  return (
    <>
      <div
        className={`h-full w-full bg-white rounded-2xl mx-2 ${sectionPadding} relative`}
      >
        <div className="flex flex-col w-full h-full ">
          <div className="w-full flex">
            <img
              src={chat?.avatar_url}
              className="rounded-full size-16 object-cover"
              alt=""
            />
            <div className="p-2">
              <h3 className="font-semibold text-xl">{chat?.name}</h3>
              <span className="text-gray-400">Không hoạt động</span>
            </div>
          </div>
          <div className="w-full h-full flex flex-col justify-end py-4">
            {messageList.map((message, index) => {
              let isFirstMessageOnDay = true
              const sendingTime = new Date(message.sending_time);
              if (index == messageList.length - 1) {
                const preSendingTime = new Date(messageList[index + 1]);
                if (sendingTime.toDateString() != preSendingTime.toDateString()){
                  isFirstMessageOnDay = false
                }
              }
              if (!isFirstMessageOnDay) {
                return (
                  <Message
                    isMine={message.user_id == me.id}
                    message={message}
                  />
                );
              }
              return (
                <>
                  <span className="self-center text-gray-400">{sendingTime.toDateString()}</span>
                  <Message
                    isMine={message.user_id == me.id}
                    message={message}
                  />
                </>
              );
            })}
          </div>
          <div className="w-full flex items-center">
            <input
              className="w-full h-10 rounded-full border-gray-400 border-2 outline-none focus:border-gray-600 px-4"
              placeholder="Message..."
              value={message}
              onChange={(e) => {
                setMessage(e.target.value);
              }}
              type="text"
            />
            <button
              className="w-fit ml-2 text-gray-600 disabled:text-gray-400"
              disabled={message == ""}
            >
              <MdSend size={34} className="" />
            </button>
          </div>
        </div>
        <div className="absolute top-0 right-0 translate-x-1/2 h-full flex items-center opacity-0 hover:opacity-100 transition-opacity duration-300">
          <button
            className="bg-gray-200 p-2 rounded-full border hover:opacity-100 opacity-60"
            onClick={handleCloseChatInfo}
          >
            {isChatInfoOpen ? (
              <MdArrowForwardIos size={30} />
            ) : (
              <MdArrowBackIos size={30} />
            )}
          </button>
        </div>
      </div>

      <div
        className={`h-full shrink-0 ${
          isChatInfoOpen ? `w-120 ${sectionPadding}` : "w-0"
        } bg-white rounded-2xl mx-2 flex flex-col transition-all duration-300 overflow-hidden`}
      ></div>
    </>
  );
}

export default ChatBox;
