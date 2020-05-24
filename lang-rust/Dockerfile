ARG base
FROM ${base}

USER gitpod
ENV HOME=/home/gitpod
WORKDIR $HOME
RUN sudo apt-get update \
    && DEBIAN_FRONTEND=noninteractive sudo apt-get install -yq \
        # Enable Rust static binary builds
        musl \
        musl-dev \
        musl-tools \
    && sudo cp /var/lib/dpkg/status /var/lib/apt/dazzle-marks/lang-rust.status \
    && sudo apt-get clean \
    && sudo rm -rf /var/lib/apt/lists/* /tmp/*

RUN cp /home/gitpod/.profile /home/gitpod/.profile_orig && \
    curl -fsSL https://sh.rustup.rs | sh -s -- -y \
    && .cargo/bin/rustup toolchain install 1.43.1 \
    && .cargo/bin/rustup default 1.43.1 \
    # Save image size by removing now-redudant stable toolchain
    && .cargo/bin/rustup toolchain uninstall stable \
    && .cargo/bin/rustup component add \
        rls \
        rust-analysis \
        rust-src \
    && .cargo/bin/rustup completions bash | sudo tee /etc/bash_completion.d/rustup.bash-completion > /dev/null \
    && .cargo/bin/rustup completions bash cargo | sudo tee /etc/bash_completion.d/rustup.cargo-bash-completion > /dev/null \
    && .cargo/bin/rustup target add x86_64-unknown-linux-musl \
    && grep -v -F -x -f /home/gitpod/.profile_orig /home/gitpod/.profile > /home/gitpod/.bashrc.d/80-rust

RUN bash -lc "cargo install cargo-watch cargo-edit cargo-tree"
