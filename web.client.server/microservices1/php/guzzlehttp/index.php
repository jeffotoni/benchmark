<?php
/** 
 * Foi testado usando  php -S localhost:8080 main.php > /dev/null 2>&1 &
 * E foi usado também nginx para nosso web server
 * php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
 * php -r "if (hash_file('sha384', 'composer-setup.php') === '55ce33d7678c5a611085589f1f3ddf8b3c52d662cd01d4ba75c0ee0459970c2200a51f492d557530c71c15d8dba01eae') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
 * php composer-setup.php
 * php -r "unlink('composer-setup.php');"
 * sudo mv composer.phar /usr/local/bin/composer
 */
require_once 'vendor/autoload.php';

use GuzzleHttp\Client;
use GuzzleHttp\Promise;
use GuzzleHttp\Exception\RequestException;

if ($_SERVER['REQUEST_METHOD'] === 'GET' && $_SERVER['REQUEST_URI'] === '/v1/user') {

$client = new Client([
    'headers' => [
        //'Connection' => 'keep-alive',
        'User-Agent' => 'Guzzle/PHP',
    ],
    'timeout' => 3,
]);

$url1 = 'http://127.0.0.1:3000/v1/avatar';
$token = 'x3939393939x39393';

try {
    $response = $client->get($url1, ['headers' => ['Authorization' => "Bearer $token"]]);
    $content = $response->getBody();
    $length = strlen($content);
    header("HTTP/1.1 200 OK");
    header("Content-Type: application/json");
    header("Engine: PHP");
    header("Date: " . date("D, d M Y H:i:s T"));
    header("Content-Length: " . $length);
    echo $content;
} catch (RequestException $e) {
    header("HTTP/1.1 500 Internal Server Error");
    echo "RequestException: " . $e->getMessage() . "\n";
    if ($e->hasResponse()) {
        echo "Response: " . $e->getResponse()->getBody() . "\n";
    }
} catch (Exception $e) {
    header("HTTP/1.1 500 Internal Server Error");
    echo "Exception: " . $e->getMessage() . "\n";
}

} else {
    header("Engine: PHP");
    header("HTTP/1.1 404 Not Found");
}
