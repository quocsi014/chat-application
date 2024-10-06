export const formatToTime = (dateTime)=>{
  const date = new Date(dateTime)
  return date.getHours() + ":" + date.getMinutes()
}