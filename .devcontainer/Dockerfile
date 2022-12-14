FROM golang:1.19.3-bullseye

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN mkdir -p /home/${USERNAME}
RUN groupadd ${USERNAME} -g ${USER_UID} && useradd -l -u ${USER_UID} -g ${USERNAME} ${USERNAME}
RUN chown ${USERNAME}:${USERNAME} /home/${USERNAME}

RUN apt-get update \
    && apt-get install --no-install-recommends -y zsh less vim \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

RUN usermod -s /usr/bin/zsh ${USERNAME}

USER ${USERNAME}

RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/cweill/gotests/gotests@latest

RUN sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
COPY zshrc /home/$USERNAME/.zshrc
RUN git clone https://github.com/zsh-users/zsh-autosuggestions /home/$USERNAME/.oh-my-zsh/custom/plugins/zsh-autosuggestions
