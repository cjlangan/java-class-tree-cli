#!/bin/sh
echo "Downloading myprogram..."
curl -L -o jhtree https://github.com/cjlangan/java-class-tree-cli/releases/latest/download/jhtree-linux
chmod +x jhtree
sudo mv myprogram /usr/local/bin/
echo "Installation complete! Run 'jhtree <directory>' to start."
