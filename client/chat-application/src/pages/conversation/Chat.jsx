import { useNavigate } from "react-router-dom";
import defaultAvatar from "../../assets/default_avatar.png";
import defaultGroupAvatar from "../../assets/default_group_avatar.png";
import { useDispatch } from "react-redux";
import { setCurrentChat } from "../../redux/conversation/conversationSlice";

const Chat = (props) => {
  const { chat } = props;
  const navigate = useNavigate()
  const dispatch = useDispatch()

  const handleOpenChat = ()=>{
    dispatch(setCurrentChat({chat: chat}))
    navigate(`/conversations/${chat.id}`)
  }

  return (
    <div onClick={()=>{handleOpenChat()}} key={chat.id} className="flex items-center  py-3 hover:bg-gray-100 rounded-lg cursor-pointer">
      <img
        src={chat.avatar_url}
        className="w-12 h-12 rounded-full mr-3 object-cover"
        onError={(e) => {
          e.target.src = chat.is_group ? defaultGroupAvatar : defaultAvatar;
        }}
      />
      <div className="flex-1 min-w-0">
        <h3 className="text-sm font-semibold text-gray-900 truncate">{chat.name}</h3>
        <p className="text-sm text-gray-500 truncate">
          {chat.user_name_sender}: {chat.last_message}
        </p>
      </div>
    </div>
  );
};

export default Chat;
