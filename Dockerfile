FROM golang:1.18

ENV DEBIAN_FRONTEND=noninteractive

RUN dpkg-reconfigure -p critical dash

RUN for i in $(seq 1001 1500); do \
        groupadd -g $i runner$i && \
        useradd -M runner$i -g $i -u $i ; \
    done

RUN apt-get update && \
    apt-get install -y libxml2 gnupg tar coreutils util-linux libc6-dev \
    binutils build-essential locales libpcre3-dev libevent-dev libgmp3-dev \
    libncurses6 libncurses5 libedit-dev libseccomp-dev rename procps python3 \
    libreadline-dev libblas-dev liblapack-dev libpcre3-dev libarpack2-dev \
    libfftw3-dev libglpk-dev libqhull-dev libqrupdate-dev libsuitesparse-dev \
    libsundials-dev libpcre2-dev && \
    rm -rf /var/lib/apt/lists/*

RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && locale-gen

COPY build.sh build.sh

RUN bash build.sh

COPY exec.go exec.go
COPY run.sh run.sh


# docker build -t golang_container:latest .
# docker run --name golang --rm -i -t golang_container /bin/bash