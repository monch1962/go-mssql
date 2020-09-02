FROM gitpod/workspace-full

# Install custom tools, runtimes, etc.
# For example "bastet", a command-line tetris clone:
# RUN brew install bastet
#
# More information: https://www.gitpod.io/docs/config-docker/

#USER root
RUN wget -qO- https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add - \
    && sudo add-apt-repository "$(wget -qO- https://packages.microsoft.com/config/ubuntu/18.04/mssql-server-2019.list)" \
    && sudo apt-get update \
    && sudo apt-get install -y mssql-server

USER root