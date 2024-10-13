# Middleman

## About

This is the official Wails Svelte template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## AnTLR

```
java -jar antlr-4.13.2-complete.jar -Dlanguage=Go -visitor Expression.g4 -o parser
```

## Gen Certificate

```
openssl genrsa -out ca.key 2048
openssl req -new -x509 -key ca.key -out ca.crt
```

## Rules


- Cancel, on request.
- Redirect, if not cancelled.
- Modify Request Body, if not cancelled.
- Modify Response Body, if not cancelled.
- Modify Headers, if not cancelled.
- Add Delay, if not cancelled.