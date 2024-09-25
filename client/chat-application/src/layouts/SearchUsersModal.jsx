import { useState, useEffect, useCallback, useRef } from "react";
import { searchUsers } from "../api/userAPI";
import { getCookie } from "../utils/cookie";
import TextField from "../components/TextField";
import Loading from "../components/Loading";
import { useDebounce } from "../hooks/useDebounce";
import { useThrottle } from "../hooks/useThrottle";
import { useNavigate } from "react-router-dom";
import defaultAvatar from "../assets/default_avatar.png";
import { BiX } from "react-icons/bi";
import { useDispatch, useSelector } from "react-redux";
import { toggleSearchUser } from "../redux/SearchUser/searchUserSlice";

function SearchUsersModal() {
  const isOpen = useSelector((state) => state.searchUser.isSearchUserOpen);
  const [searchTerm, setSearchTerm] = useState("");
  const [searchResults, setSearchResults] = useState([]);
  const [isSearching, setIsSearching] = useState(false);
  const nextPageNumber = useRef(1);
  const totalPages = useRef(null);
  const searchInputRef = useRef(null);
  const isSearchingRef = useRef(false);

  const disPatch = useDispatch();

  const navigate = useNavigate();

  useEffect(() => {
    if (isOpen) {
      setSearchTerm("");
      setSearchResults([]);
      setIsSearching(false);
      nextPageNumber.current = 1;
      totalPages.current = null;
      searchInputRef.current.focus();
    }
  }, [isOpen]);

  const handleSearch = useCallback((searchTerm) => {
    if (isSearchingRef.current) return;
    isSearchingRef.current = true;

    searchUsers(searchTerm, nextPageNumber.current, 10)
      .then((res) => {
        console.log(res.data);
        setSearchResults((prev) => [...prev, ...res.data.users]);
        nextPageNumber.current = nextPageNumber.current + 1;
        if (totalPages.current === null) {
          totalPages.current = res.data.paging.total_page;
        }
      })
      .catch((error) => {
        navigate("/500");
      })
      .finally(() => {
        setIsSearching(false);
        isSearchingRef.current = false;
      });
  }, []);

  const debouncedSearch = useDebounce(handleSearch, 500);
  const throttledSearch = useThrottle(handleSearch, 1000);

  const handleSearchChange = (e) => {
    let newSearchTerm = e.target.value.trim();
    setSearchTerm(newSearchTerm);
    setSearchResults([]);
    nextPageNumber.current = 1;
    totalPages.current = null;
    if (newSearchTerm == "") {
      return;
    }
    setIsSearching(true);
    debouncedSearch(newSearchTerm);
  };

  const handleLoadNextPage = () => {
    if (
      totalPages.current !== null &&
      nextPageNumber.current > totalPages.current
    ) {
      return;
    }
    setIsSearching(true);
    throttledSearch(searchTerm);
  };

  if (!isOpen) return null;

  const onClose = () => {
    disPatch(toggleSearchUser());
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-10 flex items-center justify-center z-50">
      <div className="bg-white p-6 rounded-3xl w-152 h-120 flex flex-col">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-bold">Find user</h1>
          <button
            className="hover:bg-gray-100 p-2 rounded-full"
            onClick={onClose}
          >
            <BiX size={24} />
          </button>
        </div>
        <div className="relative mt-2">
          <div className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
            <span className="text-gray-500">@</span>
          </div>
          <input
            type="text"
            placeholder="Type username to search for users"
            className="p-2 pl-7 border-2 bg-gray-200 outline-none border-gray-200 focus:border-gray-300 box-border rounded-full w-full"
            value={searchTerm}
            ref={searchInputRef}
            onChange={handleSearchChange}
          />
        </div>
        <div
          className="size-full overflow-y-auto mt-4 flex flex-col"
          onScroll={() => {
            handleLoadNextPage();
          }}
        >
          {searchTerm == "" && (
            <div className="text-gray-500 self-center">
              Find others to chat with
            </div>
          )}
          {searchTerm != "" && !isSearching && searchResults.length == 0 && (
            <div className="text-gray-500 self-center">No users found</div>
          )}
          {searchTerm != "" && searchResults.length > 0 && (
            <>
              {searchResults.map((user) => (
                <div
                  key={user.id}
                  className="flex justify-between hover:bg-gray-100 py-2 rounded-md"
                >
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
                  <div className="flex h-full items-center px-2">
                    {
                      !user.user_relationship.status?
                      <button className="ml-2 px-2 py-1 min-w-20 bg-blue-400 hover:bg-blue-500 text-white rounded-md">request</button>
                      :
                      user.user_relationship.status == 'ACCEPTED'?
                      <>
                      <button className="ml-2 px-2 py-1 min-w-20 bg-gray-400 hover:bg-gray-500 text-white rounded-md">block</button>
                      <button className="ml-2 px-2 py-1 min-w-20 bg-gray-100 hover:bg-white rounded-md">chat</button>
                      </>
                      :
                      user.user_relationship.status == 'PENDING' && user.id == user.user_relationship.user_id ?
                      <button className="ml-2 px-2 py-1 min-w-20 bg-gray-300 rounded-md">cancel</button>
                      :
                      user.user_relationship.status == 'PENDING' && user.id != user.user_relationship.user_id ?
                      <>
                      <button className="ml-2 px-2 py-1 min-w-20 bg-red-400 hover:bg-red-500 text-white rounded-md">reject</button>
                      <button className="ml-2 px-2 py-1 min-w-20 bg-blue-400 hover:bg-blue-500 text-white rounded-md">accept</button>
                      </>
                      :
                      user.user_relationship.status == 'BLOCKED' && user.id == user.user_relationship.blocked_user_id?                      
                      <button className="ml-2 px-2 py-1 min-w-20 bg-gray-300 rounded-md">unblock</button>
                      :
                      <button className="ml-2 px-2 py-1 min-w-20 bg-gray-300 rounded-md">view</button>
                    }
                  </div>
                </div>
              ))}
            </>
          )}
          {isSearching && (
            <Loading size={30} className="text-slate-500 self-center" />
          )}
          {totalPages.current !== null &&
            searchResults.length > 0 &&
            nextPageNumber.current > totalPages.current && (
              <div className="text-gray-500 self-center">No more users</div>
            )}
        </div>
      </div>
    </div>
  );
}

export default SearchUsersModal;
