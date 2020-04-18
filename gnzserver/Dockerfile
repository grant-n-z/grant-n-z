FROM golang:1.13

RUN mkdir /grant_n_z
WORKDIR /grant_n_z
COPY ./gnzserver /grant_n_z
COPY ./entrypoint.sh /grant_n_z
COPY ./grant_n_z_server.yaml /grant_n_z
RUN chmod 755 /grant_n_z/gnzserver

ENTRYPOINT ["./entrypoint.sh"]
