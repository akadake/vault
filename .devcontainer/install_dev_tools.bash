#!/usr/bin/env bash

apt-get install -y nodejs npm vim

npm install --global yarn

echo '"\e[A": history-search-backward
"\e[B": history-search-forward
set show-all-if-ambiguous on
set completion-ignore-case on' > /root/.inputrc

echo "export VAULT_ADDR=http://127.0.0.1:8200" >> /root/.bashrc