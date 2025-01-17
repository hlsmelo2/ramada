<?php

namespace App\Http\Controllers;

use Exception;
use GuzzleHttp\Client;
use GuzzleHttp\HandlerStack;
use Psr\Http\Message\RequestInterface;

class RequestsController extends Client
{
    public function __construct(array $config = [])
    {

        $header = 'Authorization';
        $token = $_COOKIE['_token'] ?? 0;
        $value = $token !== 0 ? "Bearer $token" : 0;
        $handler = HandlerStack::create();

        $handler->push(function (callable $handler) use ($header, $value) {
                return function (RequestInterface $request, array $options) use ($handler, $header, $value) {
                    if ($value !== 0) {
                        $request = $request->withHeader($header, $value);
                    }

                    try {
                        return $handler($request, $options);
                    } catch (Exception $e) {
                        return $e;
                    }
                };
            }, 'middleWareAuth');

        $config['handler'] = $handler;

        parent::__construct($config);
    }
}
