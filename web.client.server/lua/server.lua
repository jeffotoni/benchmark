-- @jeffotoni
-- sudo apt-get install lua5.3
-- sudo apt-get install luarocks
-- sudo luarocks install luasocket
-- sudo luarocks install luasec

local socket = require("socket")
local http = require("socket.http")
local url = require("socket.url")
-- local https = require("ssl.https")
local ltn12 = require("ltn12")

local function getRemoteData()
    local response_body = {}
    local url = "http://127.0.0.1:3000/v1/customer/get"
    local result, response_code, response_headers, response_status = http.request {
        url = url,
        method = "GET",
        sink = ltn12.sink.table(response_body),
    }
    return table.concat(response_body), response_code
end

local function onRequest(client)
    local request, error = client:receive('*l')
    local path = string.match(request, 'GET (.-) HTTP')

    if path == "/v1/client/get" then
        local data, code = getRemoteData()

        if code == 200 then
            client:send("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: " .. #data .. "\r\n\r\n")
            client:send(data)
        else
            client:send("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\nContent-Length: 13\r\n\r\nFailed fetch.")
        end
    else
        client:send("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\nContent-Length: 9\r\n\r\nNot found.")
    end

    client:close()
end

local server = socket.bind("localhost", 8080)
print("Server listening at http://localhost:8080")
print("\027[0;33m[GET]\027[0m \027[0;33m/v1/customer/get\027[0m")

while true do
    local client = server:accept()
    onRequest(client)
end

