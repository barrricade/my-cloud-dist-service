# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.202.5/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.17-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN echo "deb http://mirrors.aliyun.com/debian/ bullseye main non-free contrib \n\
            deb-src http://mirrors.aliyun.com/debian/ bullseye main non-free contrib \n\
            deb http://mirrors.aliyun.com/debian-security/ bullseye-security main \n\
            deb-src http://mirrors.aliyun.com/debian-security/ bullseye-security main \n\
            deb http://mirrors.aliyun.com/debian/ bullseye-updates main non-free contrib \n\
            deb-src http://mirrors.aliyun.com/debian/ bullseye-updates main non-free contrib \n\
            deb http://mirrors.aliyun.com/debian/ bullseye-backports main non-free contrib \n\
            deb-src http://mirrors.aliyun.com/debian/ bullseye-backports main non-free contrib \n\
            deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free \n\
            deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free \n\
            deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free \n\
            deb https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free \n\
    " > /etc/apt/sources.list &&\
    apt-get update && apt-get upgrade -y  &&\
    apt-get install sqlite3 -y &&\
    # go env -w  GOPROXY=https://goproxy.cn,direct &&\
    mkdir /mnt/sqliteDB &&\
    if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go env -w  GOPROXY=https://goproxy.cn,direct

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1