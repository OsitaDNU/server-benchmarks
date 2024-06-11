<?php

use Swoole\Http\Server;
use Swoole\Http\Request;
use Swoole\Http\Response;

$server = new Server("0.0.0.0", 8080);

$server->on("request", function (Request $request, Response $response) {
    $post = [
        'title' => 'Exploring the Tranquil Beauty of Nature: A Journey to Serendipity',
        'content' => "There is something incredibly captivating about the serene embrace of nature. From the gentle rustling of leaves to the melodic chirping of birds, nature offers a symphony of tranquility that beckons us to step away from the hustle and bustle of daily life. This blog post is a celebration of nature's beauty and the profound peace it brings to our souls.",
        'url' => 'http://localhost:8080/blog/1',
        'created_at' => date('c'),
        'tags' => ['PHP', 'Swoole', 'Benchmark', 'Web'],
        'views' => 123
    ];

    $response->header("Content-Type", "application/json");
    $response->end(json_encode($post));
});

$server->start();
