import { formatToTime } from "../utils/timeFormat"

export default function(props){
  const {message, isMine} = props
  return(
    <div key={message.id} className={`flex flex-col ${isMine ? "self-end items-end" : "self-start items-start"}`}>
      {
        isMine?
        <></>
        :
        <div className="mb-2">
          <img src={message.user.avatar_url} alt="" className="size-10 rounded-full object-cover" />
        </div>
      }
      <div className={`p-2  max-w-52 w-fit ${isMine?"bg-blue-600 text-white":"bg-gray-300"} rounded-xl`}>
        {message.message}
      </div>
      <span className="text-gray-400">{formatToTime(message.sending_time)}</span>
    </div>
  )
}