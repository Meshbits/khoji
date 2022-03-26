FROM amd64/ubuntu

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -y \
	make git software-properties-common wget

RUN add-apt-repository ppa:longsleep/golang-backports && \
    apt install -y golang-go

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

RUN source /etc/lsb-release && echo "deb https://download.rethinkdb.com/repository/ubuntu-$DISTRIB_CODENAME $DISTRIB_CODENAME main" | tee /etc/apt/sources.list.d/rethinkdb.list && \
	wget -qO- https://download.rethinkdb.com/repository/raw/pubkey.gpg | apt-key add - && \
	apt-get update && \
	apt-get install -y rethinkdb

WORKDIR /root

RUN apt autoremove -y

RUN git clone https://github.com/meshbits/khoji.git && \
	cd khoji && \
	make
