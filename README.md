# uws
Universal Web Server for development purposes. It is similar to `python -m SimpleHTTPServer` or
`php -S 127.0.0.0:8080` commands but requires less typing and tries to detect the stack to serve it accordingly.

# Features

* `uws` to serve a web app located in the current directory
* Detect PHP stack and launch `php -S` command
* Detect Rack-based Ruby app and launch `rackup` command via `bundler`
* Detect NodeJS-based app and launch `yarn` or `npm` command
* Launch web server to serve a static site or any other directory

# Installation

If you have Go toolchain installed, you can use the following command to install `uws`:
```
go install github.com/sibprogrammer/uws@latest
```
