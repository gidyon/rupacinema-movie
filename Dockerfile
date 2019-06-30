FROM scratch
# Working directory
WORKDIR /

# server.bin is app binary, certs needed for TLS
COPY server.bin /
COPY certs /certs

# Entry, pass -uflag to use default config from flag
ENTRYPOINT [ "/server.bin" ]