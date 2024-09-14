const Chat = ({ user, lastMessage }) => {
  return (
    <div className="flex items-center  py-3 hover:bg-gray-100 rounded-lg cursor-pointer">
      <img src={user.avatar} alt={user.name} className="w-12 h-12 rounded-full mr-3" />
      <div className="flex-1 min-w-0">
        <h3 className="text-sm font-semibold text-gray-900 truncate">{user.name}</h3>
        <p className="text-sm text-gray-500 truncate">{lastMessage}</p>
      </div>
    </div>
  );
};

export default Chat;