#!/bin/sh
echo "Downloading jhtree..."
curl -L -o jhtree https://github.com/cjlangan/java-class-tree-cli/releases/latest/download/jhtree-linux
chmod +x jhtree
sudo mv jhtree /usr/local/bin/
echo "Installation complete! Run 'jhtree <directory>' to start."
