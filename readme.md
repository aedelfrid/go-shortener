# URL Shortener

GOlang URL shortener with Redis as the storage backend.

## Features

- Shorten URLs
- Redirect to original URLs
- Store URL mappings in Redis

## Plan

1. Set up a basic HTTP server in Go.
2. Create an endpoint to accept URL shortening requests.
3. Generate a unique short code for each URL.
4. Store the mapping of short code to original URL in Redis.
5. Create an endpoint to handle redirection based on the short code.
6. Implement error handling and edge cases (e.g., invalid URLs, duplicate short codes).

