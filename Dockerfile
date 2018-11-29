FROM ubuntu:16.04
MAINTAINER annp

COPY ./hello_world /app/hello_world
RUN chmod +x /app/hello_world

ENV PORT 8000
EXPOSE 800

ENTRYPOINT /app/hello_world
