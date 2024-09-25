import defaultAvatar from "../../assets/default_avatar.png";
import defaultGroupAvatar from "../../assets/default_group_avatar.png";

const Chat = (props) => {
  const { avatarURL, name, lastMessage, userNameSender, isGroup } = props;

  return (
    <div className="flex items-center  py-3 hover:bg-gray-100 rounded-lg cursor-pointer">
      <img
        src={avatarURL }
        className="w-12 h-12 rounded-full mr-3 object-cover"
        onError={(e) => {
          e.target.src = isGroup ? defaultGroupAvatar : defaultAvatar;
        }}
      />
      <div className="flex-1 min-w-0">
        <h3 className="text-sm font-semibold text-gray-900 truncate">{name}</h3>
        <p className="text-sm text-gray-500 truncate">
          {userNameSender}: {lastMessage}
        </p>
      </div>
    </div>
  );
};

export default Chat;
