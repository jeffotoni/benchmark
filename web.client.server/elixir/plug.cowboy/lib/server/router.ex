defmodule Server.Router do
  use Plug.Router
  import Plug.Conn
  #import Poison
  alias Server.HttpClient

  plug :match
  plug :dispatch

  get "/v1/client/get" do
    case HttpClient.fetch_data() do
      {:ok, body} ->
        send_resp(conn, 200, body)
      {:error, error} ->
        send_resp(conn, 500, error)
    end
  end

  # suas outras rotas

  match _ do
    send_resp(conn, 404, "Not found")
  end
end
