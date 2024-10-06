import defaultAvatar from "../assets/default_avatar.png";

export default function (props) {
  const { user } = props;
  return (
    <div className="flex">
      <img
        src={user.avatar_url || defaultAvatar}
        alt={user.username}
        className="w-10 h-10 object-cover rounded-full mr-2"
      />
      <div className="flex flex-col">
        <h3 className="font-bold">
          {user.firstname} {user.lastname}
        </h3>
        <p className="text-gray-500">@{user.username}</p>
      </div>
    </div>
  );
}
