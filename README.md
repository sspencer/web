# web

Minimal static web server with minimal configuration.  Just run to serve current directory or specify one as only command line argument.

## Build

    go install 

## Run

    web [-p PORT] [directory]

Examples:

Serve current directory on ':8080'

    web

Serve /doc/root on 7777

    web -p 7777 /doc/root
