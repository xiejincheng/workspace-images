FROM gitpod/workspace-full

# Install custom tools, runtime, etc.
RUN curl -OL https://github.com/moby/buildkit/releases/download/v0.9.0/buildkit-v0.9.0.linux-amd64.tar.gz \
    && tar xzfv buildkit-v0.9.0.linux-amd64.tar.gz \
    && mkdir /workspace/buildkit \
    && sudo mv bin/* /workspace/buildkit
