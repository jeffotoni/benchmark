// @jeffotoni
// manual: https://modules.vlang.io
module main

import net.http { CommonHeader, Request, Response, Server }
import time
//import strconv

struct StructHandler {}

const url_client = "http://localhost:3000/v1/avatar"

fn (h StructHandler) handle(req Request) Response {
	mut res := Response{
		header: http.new_header_from_map({
			CommonHeader.content_type: 'application/json'
		})
	}

	if req.url == '/v1/user' {
		mut headers := http.new_header()
    	headers.add(http.CommonHeader.content_language, 'Vlang')
    	headers.add(http.CommonHeader.content_location, '/v1/user')
    	headers.add(http.CommonHeader.date, time.now().custom_format("2006-01-02T15:04:05.000Z"))
    	headers.add(http.CommonHeader.connection, 'keep-alive')

		resp := http.get(url_client) or {
			eprintln('Erro ao fazer a chamada para /v1/avatar: $err')
			res.status_code = 500
			res.body = 'Erro ao fazer a chamada para /v1/avatar'
			return res
		}

		if resp.status_code != 200 {
			res.status_code = 500
			res.body = 'Erro ao fazer a chamada para /v1/avatar'
			return res
		}

		body := resp.body
		
		//length := strconv.f64_to_str_l(body.len)
    	//headers.add(http.CommonHeader.content_length, length)

		res.status_code = 200
		res.body = body
		res.header = headers
	} else {
		res.status_code = 404
		res.body = 'Not found'
	}

	return res
}

fn main() {
	mut server := Server{
		handler: StructHandler{}
	}
	server.listen_and_serve()
}
