# tcpServer

### Minimum requirement:

- [x] TCP server takes in any request text per line and send a query to an external API, until client send 'quit' or timed out.
- [x] TCP server can accept multiple connections at the same time.
- [x] As for the external API, the choice is yours, or even a mock.
- [x] We do have a API request rate limit for the external API: 30 requests per second.
- [x] Try not to use 3rd-party libraries unless necessary.
- [x] Your code should be runnable and functional in happy cases.

### Going deeper:
- [x] Please consider the external API could be unreachable or unavailable.
- [ ] Mount a HTTP endpoint on your server and display some statistics: current connection count, current request rate, processed request count, remaing jobs...etc
- [ ] The rest is for you to show your strength in programming. e.g., tests, documents, user interface, architecture.
