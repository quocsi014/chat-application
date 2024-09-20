import { useState } from "react";
import { MdArrowBackIos, MdArrowForwardIos } from "react-icons/md";

function ChatBox() {
  const [isChatInfoOpen, setIsChatInfoOpen] = useState(true);
  const handleCloseChatInfo = () => {
    setIsChatInfoOpen(!isChatInfoOpen);
  };
  
  const sectionPadding = "p-5";

  return (
    <>
      <div
        className={`h-full w-full bg-white rounded-2xl mx-2 ${sectionPadding} relative`}
      >
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
