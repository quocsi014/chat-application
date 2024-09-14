function ServerError(){
  return (
    <div className="flex flex-col items-center justify-center h-screen text-red-600">
      <h1 className="text-5xl font-bold">500</h1>
      <p className="text-2xl font-bold">Internal Server Error</p>
    </div>
  )
}

export default ServerError