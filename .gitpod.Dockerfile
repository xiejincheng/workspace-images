FROM gitpod/workspace-full

# Install custom tools, runtime, etc.
RUN cd /tmp \
    && curl -OL https://github.com/moby/buildkit/releases/download/v0.9.0/buildkit-v0.9.0.linux-amd64.tar.gz \
    && tar xzfv buildkit-v0.9.0.linux-amd64.tar.gz \
    && sudo mv bin/* /usr/bin \
    && cd -
