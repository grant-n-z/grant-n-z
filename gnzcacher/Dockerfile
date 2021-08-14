FROM golang:1.13

RUN mkdir /grant_n_z
WORKDIR /grant_n_z
COPY ./ /grant_n_z

WORKDIR /grant_n_z/

RUN cd gnzcacher && \
    make build-linux-amd64 && \
    chmod 755 build/linux-amd64/GrantNZCacher

WORKDIR /gnzcacher/build/linux-amd64/

ENTRYPOINT ["./entrypoint.sh"]
