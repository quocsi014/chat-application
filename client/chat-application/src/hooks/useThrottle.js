import { useRef, useCallback } from "react";

export const useThrottle = (callback, delay) => {
  const shouldWait = useRef(false);

  return useCallback((...args) => {
    console.log("call")
    if (shouldWait.current) return;
    callback(...args);
    shouldWait.current = true;
    setTimeout(() => {
      shouldWait.current = false;
    }, delay);
  }, [callback, delay]);
};