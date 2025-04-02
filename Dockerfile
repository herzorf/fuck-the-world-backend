FROM ubuntu:latest
LABEL authors="herzorf"

ENTRYPOINT ["top", "-b"]