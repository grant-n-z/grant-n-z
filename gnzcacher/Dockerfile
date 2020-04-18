FROM golang:1.13

RUN mkdir /grant_n_z
WORKDIR /grant_n_z
COPY ./gnzcacher /grant_n_z
COPY ./entrypoint.sh /grant_n_z
COPY ./grant_n_z_cacher.yaml /grant_n_z
COPY ./grant_n_z_cacher.txt /grant_n_z
RUN chmod 755 /grant_n_z/gnzcacher

ENTRYPOINT ["./entrypoint.sh"]
