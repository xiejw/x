# vim: filetype=dockerfile

FROM xiejw/baseimage

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        openjdk-17-jdk-headless scala \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workdir

CMD ["bash"]
