defmodule Server.HttpClient do
  require Logger
  import HTTPoison.Response
  @timeout 2000

  def fetch_data do
    case HTTPoison.get("http://localhost:3000/v1/customer/get", [], [timeout: @timeout]) do
      {:ok, %HTTPoison.Response{status_code: 200, body: body}} ->
        {:ok, body}
      {:error, %HTTPoison.Error{reason: reason}} ->
        Logger.error("Error fetching data from remote service: #{reason}")
        {:error, "Error fetching data from remote service"}
    end
  end
end
