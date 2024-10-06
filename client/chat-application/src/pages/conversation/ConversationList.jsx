import { BiSearch } from "react-icons/bi"
import Chat from "./Chat"
import { useDispatch } from "react-redux";
import { toggleSearchUser } from "../../redux/SearchUser/searchUserSlice";
import { useEffect, useState } from "react";
import { getConversations } from "../../api/conversationAPI";
import { setCurrentChat } from "../../redux/conversation/conversationSlice";

function ConversationList(){
  const [conversations, setConversations] = useState([])

  const dispatch = useDispatch()

  const sectionPadding = "p-5";

  useEffect(()=>{
    getConversations()
    .then(res => {
      setConversations(res.data.items)
      dispatch(setCurrentChat({chat: res.data.items[0]}))
    })
    .catch(error => {
      console.log(error)
    })
  },[])
  
  return(
    <div
        className={`h-full shrink-0 w-120 bg-white rounded-2xl mx-2 ${sectionPadding}`}
      >
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold">Chats</h1>
          <button
            onClick={() => {dispatch(toggleSearchUser())}}
            className="rounded-full hover:bg-gray-200 bg-gray-100 p-2 border box-border"
          >
            <BiSearch size={24}/>
          </button>
        </div>
        <input
          type="text"
          placeholder="Find a chat (Ctrl + K)"
          className="p-2 border-2 bg-gray-200 outline-none border-gray-200 focus:border-gray-300 box-border rounded-full w-full mt-2"
        />
        <div className="mt-4 space-y-2">
          {conversations.map((chat) => (
            <Chat
              key={chat.id}
              chat = {chat}
            />
          ))}
        </div>
      </div>
  )
}

export default ConversationList