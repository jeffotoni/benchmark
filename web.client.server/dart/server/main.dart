import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;

void main() async {
  final server = await HttpServer.bind(InternetAddress.anyIPv4, 8080);
  print('Serving at http://${server.address.host}:${server.port}');

  await for (HttpRequest request in server) {
    handleRequest(request);
  }
}

Future<void> handleRequest(HttpRequest request) async {
  if (request.method == 'GET' && request.uri.path == '/v1/client/get') {
    try {
      var response = await http.get(Uri.parse('http://localhost:3000/v1/customer/get'));
      if (response.statusCode == 200) {
        var jsonResponse = json.decode(response.body);

        request.response
          ..statusCode = HttpStatus.ok
          ..headers.contentType = ContentType.json
          ..write(json.encode(jsonResponse))
          ..close();
      } else {
        request.response
          ..statusCode = HttpStatus.badGateway
          ..write('Error fetching data from the customer API')
          ..close();
      }
    } catch (e) {
      request.response
        ..statusCode = HttpStatus.internalServerError
        ..write('An error occurred while fetching data from the customer API: $e')
        ..close();
    }
  } else {
    request.response
      ..statusCode = HttpStatus.notFound
      ..write('Not found')
      ..close();
  }
}
