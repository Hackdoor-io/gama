defmodule GamaWeb.Router do
  use GamaWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/", GamaWeb do
    pipe_through :api
    
    get "/", PageController, :index
    get "/:target/:id/:path", PageController, :toExternalResource
    get "/*path", PageController, :catchAllPath
  end
end
