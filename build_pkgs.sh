# install languages binaries

mkdir -p /tmp/pkg/

# install NodeJS v14.20.0
BUILD_FOLDER_NAME=node_v14.20.0
PKG_NAME=node-v14.20.0-linux-x64.tar.xz
PKG_SOURCE=https://nodejs.org/dist/v14.20.0/$PKG_NAME
mkdir -p /tmp/pkg/$BUILD_FOLDER_NAME
wget $PKG_SOURCE -P /tmp
tar xf /tmp/$PKG_NAME --strip-components=1 --directory=/tmp/pkg/$BUILD_FOLDER_NAME
rm /tmp/$PKG_NAME
echo "PATH=/tmp/pkg/$BUILD_FOLDER_NAME/bin:$PATH" >> /tmp/pkg/$BUILD_FOLDER_NAME/.env



# install Python 3.10.0
# BUILD_FOLDER_NAME=python_v3.10.0
# PKG_NAME=Python-3.10.0.tgz
# PKG_SOURCE=https://www.python.org/ftp/python/3.10.0/$PKG_NAME
# mkdir -p /tmp/pkg/$BUILD_FOLDER_NAME
# wget $PKG_SOURCE -P /tmp
# tar xzf /tmp/$PKG_NAME --strip-components=1 --directory=/tmp/pkg/$BUILD_FOLDER_NAME
# rm /tmp/$PKG_NAME
# cd /tmp/pkg/$BUILD_FOLDER_NAME
# PREFIX=$(realpath $(dirname .))
# ./configure --prefix "$PREFIX" --with-ensurepip=install
# make -j$(nproc)
# make install -j$(nproc)
# rm -rf build
# echo "PATH=/tmp/pkg/$BUILD_FOLDER_NAME/bin:$PATH" >> /tmp/pkg/$BUILD_FOLDER_NAME/.env
# bin/pip3 install numpy scipy pandas bcrypt
# cd /tmp



# install Java 15.0.2
# build_path=/usr/local/lib/openjdk-15.0.2_linux-x64_bin
# pkg_file=openjdk-15.0.2_linux-x64_bin.tar.gz
# mkdir -p $build_path
# mkdir -p /tmp/tar_pkgs/
# curl "https://download.java.net/java/GA/jdk15.0.2/0d1cfde4252546c6931946de8db48ee2/7/GPL/$pkg_file" -o /tmp/tar_pkgs/$pkg_file
# tar xzf /tmp/tar_pkgs/$pkg_file --strip-components=1 --directory=$build_path
# rm /tmp/tar_pkgs/$pkg_file
# export PATH=$build_path/bin:$PATH
# echo "export PATH=$build_path/bin:$PATH" >> ~/.bashrc