FROM golang:1.13

RUN mkdir /grant_n_z
WORKDIR /grant_n_z
COPY ./ /grant_n_z

WORKDIR /grant_n_z/

RUN cd gnzserver && \
    make build-linux-amd64 && \
    chmod 755 build/linux-amd64/GrantNZServer

ENTRYPOINT ["/grant_n_z/gnzserver/build/linux-amd64/entrypoint.sh"]
