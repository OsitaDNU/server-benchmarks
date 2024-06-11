const http = require("http");

const server = http.createServer((req, res) => {
  if (req.url === "/") {
    const post = {
      title: "Exploring the Tranquil Beauty of Nature: A Journey to Serendipity",
      content:
        "There is something incredibly captivating about the serene embrace of nature. From the gentle rustling of leaves to the melodic chirping of birds, nature offers a symphony of tranquility that beckons us to step away from the hustle and bustle of daily life. This blog post is a celebration of nature's beauty and the profound peace it brings to our souls.",
      url: "http://localhost:8080/blog/1",
      author: "Author Name",
      created_at: new Date().toISOString(),
      tags: ["Node.js", "Benchmark", "Web"],
      views: 123,
    };

    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(JSON.stringify(post));
  } else {
    res.writeHead(404, { "Content-Type": "text/plain" });
    res.end("Not Found");
  }
});

const PORT = 8080;
server.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
