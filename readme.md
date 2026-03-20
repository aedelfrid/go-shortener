# URL Shortener

GOlang URL shortener with Redis as the storage backend.

## Features

- Shorten long URLs to a fixed-length short code.
- Redirect users to the original URL when they access the short code.

## Plan

1. **Setup Redis**: Install and run Redis on your local machine or use a cloud-based Redis service.
2. **Create a Go Project**: Initialize a new Go project and set up the necessary dependencies.
3. **Implement URL Shortening Logic**: Create a function to generate a short code for a given long URL and store the mapping in Redis.

----

1. The "Identity" Upgrade (The Encoder)

Right now, every link is abc123. We need to replace that with the Base62 logic we discussed.

    The Task: Write a base62.go utility.

    The Result: Every time you hit /shorten, the counter goes up (1001→1002), and the code changes (g9→gA).

2. The "Memory" Swap (Redis Integration)

The current urlStore map is "volatile"—if you stop your Go server, all your shortened links vanish.

    The Task: Connect to Upstash (or local Redis).

    The Result: Links become permanent. You can restart your computer, and localhost:4000/g9 will still work.

3. The "Guardrail" Layer (Validation)

What if someone sends "not-a-url" or an empty string?

    The Task: Add a URL parser to the Go backend to ensure the input actually starts with http or https.

    The Result: Your server won't crash or store "garbage" data.

4. The "Face" (Next.js Frontend)

Once the API is bulletproof, we move to the browser.

    The Task: Build a sleek Tailwind UI with an input field and a "Recently Created" list.

    The Result: A shareable portfolio piece that looks like a real startup product.

5. The "Analytics" Flex (Optional Bonus)

If you want to go above and beyond for recruiters:

    The Task: Every time someone clicks a short link, increment a "click counter" in Redis.

    The Result: Your dashboard can show how many times each link was clicked. This proves you can handle write-heavy operations.