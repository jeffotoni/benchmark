defmodule Server.Application do
  use Application

  def start(_type, _args) do
    children = [
      {Plug.Cowboy, scheme: :http, plug: Server.Router, options: [port: 8080, keepalive: false]}
    ]

    IO.puts("Run Server port: 0.0.0.0:8080")

    opts = [strategy: :one_for_one, name: Server.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
