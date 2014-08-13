SimpleServer
============

Simple HTTP Server in Go

Requires having Go 1.3 or higher installed.

Build the SimpleServer binary, Execute 
```
Go build 
```
or 
```
Go install
```

Usage:
```
SimpleServer
``` 


By default the server will listen on port 8080 and serve static files from the directory where the SimpleServer resides

To specify a port or directory:
```
SimpleServer -port=PORT -dir=DIRECTORY
``` 
