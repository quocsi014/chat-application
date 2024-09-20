import { BiSearch } from "react-icons/bi"
import Chat from "../../components/Chat"
import { useDispatch } from "react-redux";
import { toggleSearchUser } from "../../redux/SearchUser/searchUserSlice";

function ConversationList(props){

  const {setIsSearchModalOpen} = props
  const dispatch = useDispatch()

  const sectionPadding = "p-5";

  const fakeChats = [
    {
      id: 1,
      user: {
        name: "Nguyễn Văn A",
        avatar: "https://i.pravatar.cc/150?img=1",
      },
      lastMessage: "Xin chào, bạn khỏe không?",
    },
    {
      id: 2,
      user: {
        name: "Trần Thị B",
        avatar: "https://i.pravatar.cc/150?img=2",
      },
      lastMessage: "Hẹn gặp lại nhé!",
    },
    {
      id: 3,
      user: {
        name: "Lê Văn C",
        avatar: "https://i.pravatar.cc/150?img=3",
      },
      lastMessage: "Ok, tôi sẽ gửi file sau.",
    },
  ];

  const openSearchUser = ()=>{
    dispatch(toggleSearchUser())
  }


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
          {fakeChats.map((chat) => (
            <Chat
              key={chat.id}
              user={chat.user}
              lastMessage={chat.lastMessage}
            />
          ))}
        </div>
      </div>
  )
}

export default ConversationList