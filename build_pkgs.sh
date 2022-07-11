# install languages binaries

mkdir -p /pkg/

# install NodeJS v14.20.0
BUILD_FOLDER_NAME=node_v14.20.0
PKG_NAME=node-v14.20.0-linux-x64.tar.xz
PKG_SOURCE=https://nodejs.org/dist/v14.20.0/$PKG_NAME
mkdir -p /pkg/$BUILD_FOLDER_NAME
wget $PKG_SOURCE -P /pkg
tar xf /pkg/$PKG_NAME --strip-components=1 --directory=/pkg/$BUILD_FOLDER_NAME
rm /pkg/$PKG_NAME
echo "PATH=/pkg/$BUILD_FOLDER_NAME/bin:$PATH" >> /pkg/$BUILD_FOLDER_NAME/.env
ls /pkg/node_v14.20.0


# install Python 3.10.0
# BUILD_FOLDER_NAME=python_v3.10.0
# PKG_NAME=Python-3.10.0.tgz
# PKG_SOURCE=https://www.python.org/ftp/python/3.10.0/$PKG_NAME
# mkdir -p /pkg/$BUILD_FOLDER_NAME
# wget $PKG_SOURCE -P /tmp
# tar xzf /pkg/$PKG_NAME --strip-components=1 --directory=/pkg/$BUILD_FOLDER_NAME
# rm /pkg/$PKG_NAME
# cd /pkg/$BUILD_FOLDER_NAME
# PREFIX=$(realpath $(dirname .))
# ./configure --prefix "$PREFIX" --with-ensurepip=install
# make -j$(nproc)
# make install -j$(nproc)
# rm -rf build
# echo "PATH=/pkg/$BUILD_FOLDER_NAME/bin:$PATH" >> /pkg/$BUILD_FOLDER_NAME/.env
# bin/pip3 install numpy scipy pandas bcrypt
# cd /tmp
