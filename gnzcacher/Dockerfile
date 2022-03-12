FROM golang:1.15

RUN mkdir /grant_n_z
WORKDIR /grant_n_z
COPY ./ /grant_n_z

WORKDIR /grant_n_z/

RUN cd gnzcacher && \
    make build-linux-amd64 && \
    chmod 755 build/linux-amd64/GrantNZCacher

ENTRYPOINT ["/grant_n_z/gnzcacher/build/linux-amd64/entrypoint.sh"]
