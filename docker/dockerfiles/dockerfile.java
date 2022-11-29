# vim: filetype=dockerfile

FROM xiejw/ubuntu

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        openjdk-18-jdk-headless \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workdir

CMD ["bash"]
