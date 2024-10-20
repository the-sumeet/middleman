# Middleman

Middleman is a tool that allows you to intercept and modify HTTP requests and responses. It is useful for testing and debugging web applications.

## Demo Video

[Watch the video](https://www.youtube.com/watch?v=tVchu5zLN_o)

## How to Use?

- Add CA Certificate to your browser or system. Path given in the settings.
- Start the Middleman.
- Use middleman server port as proxy in your browser or system.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.


## Rules


- Cancel, on request.
- Redirect, if not cancelled.
- Modify Request Body, if not cancelled.
- Modify Response Body, if not cancelled.
- Modify Headers, if not cancelled.
- Add Delay, if not cancelled.