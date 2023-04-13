defmodule Server.Router do
  use Plug.Router
  import Plug.Conn
  alias Server.HttpClient

  plug :match
  plug :dispatch

  get "/v1/client/get" do
    case HttpClient.fetch_data() do
      {:ok, body} ->
        conn
        |> put_resp_header("Content-Type", "application/json")
        |> put_resp_header("Engine", "Elixir")
        |> put_resp_header("Location", "/v1/client/get")
        |> put_resp_header("Date", DateTime.utc_now() |> DateTime.to_iso8601())
        |> put_resp_header("Content-Length", Integer.to_string(byte_size(body)))
        |> send_resp(200, body)
      {:error, error} ->
        conn
        |> put_resp_header("Content-Type", "application/json")
        |> put_resp_header("Engine", "Elixir")
        |> put_resp_header("Location", "/v1/client/get")
        |> put_resp_header("Date", DateTime.utc_now() |> DateTime.to_iso8601())
        |> put_resp_header("Content-Length", Integer.to_string(byte_size(error)))
        |> send_resp(500, error)
        #send_resp(conn, 500, error)
    end
  end

  match _ do
    conn
    |> put_resp_header("Content-Type", "application/json")
    |> put_resp_header("Engine", "Elixir")
    |> put_resp_header("Location", "/v1/client/get")
    |> put_resp_header("Date", DateTime.utc_now() |> DateTime.to_iso8601())
    |> send_resp(404, "Not found")
  end
end
