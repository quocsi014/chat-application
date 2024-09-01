import { useState, useEffect, useCallback, useRef } from "react";
import { searchUsers } from "../api/userAPI";
import { getCookie } from "../utils/cookie";
import TextField from "./TextField";
import Loading from "./Loading";
import { useDebounce } from "../hooks/useDebounce";
import { useThrottle } from "../hooks/useThrottle";
import { useNavigate } from "react-router-dom";
import defaultAvatar from "../assets/default_avatar.png";
import { BiX } from "react-icons/bi";

function SearchUsersModal({ isOpen, onClose }) {
  const [searchTerm, setSearchTerm] = useState("");
  const [searchResults, setSearchResults] = useState([]);
  const [isSearching, setIsSearching] = useState(false);
  const nextPageNumber = useRef(1);
  const totalPages = useRef(null);
  const searchInputRef = useRef(null);
  const isSearchingRef = useRef(false);

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
        console.log(res.data)
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
        <input
          type="text"
          placeholder="Type username to search for users"
          className="p-2 border-2 bg-gray-200 outline-none border-gray-200 focus:border-gray-300 box-border rounded-full w-full mt-2"
          value={searchTerm}
          ref={searchInputRef}
          onChange={(e) => {
            handleSearchChange(e);
          }}
        />
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
                  className="flex items-center hover:bg-gray-100 py-2 rounded-md"
                >
                  <img
                    src={user.avatar_url || defaultAvatar}
                    alt={user.username}
                    className="w-10 h-10 rounded-full mr-2"
                  />
                  <span>{user.username}</span>
                </div>
              ))}
            </>
          )}
          {isSearching && (
            <Loading size={30} className="text-slate-500 self-center" />
          )}
          {
            totalPages.current !== null && nextPageNumber.current > totalPages.current && (
              <div className="text-gray-500 self-center">No more users</div>
            )
          }
        </div>
      </div>
    </div>
  );
}

export default SearchUsersModal;
