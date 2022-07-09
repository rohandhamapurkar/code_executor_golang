# install languages binaries




# install NodeJS v14.20.0
build_path=/usr/local/lib/node-v14.20.0-linux-x64
pkg_file=node-v14.20.0-linux-x64.tar.xz
mkdir -p $build_path
mkdir -p /tmp/tar_pkgs/
curl "https://nodejs.org/dist/v14.20.0/$pkg_file" -o /tmp/tar_pkgs/$pkg_file
tar xf /tmp/tar_pkgs/$pkg_file --strip-components=1 --directory=$build_path
rm /tmp/tar_pkgs/$pkg_file
export PATH=$build_path/bin:$PATH
echo "export PATH=$build_path/bin:$PATH" >> ~/.bashrc



# install Python 3.10.0
# build_path=/usr/local/lib/Python-3.10.0
# pkg_file=Python-3.10.0.tgz
# mkdir -p $build_path
# mkdir -p /tmp/tar_pkgs/
# curl "https://www.python.org/ftp/python/3.10.0/$pkg_file" -o /tmp/tar_pkgs/$pkg_file
# tar xzf /tmp/tar_pkgs/$pkg_file --strip-components=1 --directory=$build_path
# rm /tmp/tar_pkgs/$pkg_file
# cd $build_path
# PREFIX=$(realpath $(dirname .))
# ./configure --prefix "$PREFIX" --with-ensurepip=install
# make -j$(nproc)
# make install -j$(nproc)
# rm -rf build
# export PATH=$build_path/bin:$PATH
# echo "export PATH=$build_path/bin:$PATH" >> ~/.bashrc
# pip3 install numpy scipy pandas bcrypt




# install Java 15.0.2
build_path=/usr/local/lib/openjdk-15.0.2_linux-x64_bin
pkg_file=openjdk-15.0.2_linux-x64_bin.tar.gz
mkdir -p $build_path
mkdir -p /tmp/tar_pkgs/
curl "https://download.java.net/java/GA/jdk15.0.2/0d1cfde4252546c6931946de8db48ee2/7/GPL/$pkg_file" -o /tmp/tar_pkgs/$pkg_file
tar xzf /tmp/tar_pkgs/$pkg_file --strip-components=1 --directory=$build_path
rm /tmp/tar_pkgs/$pkg_file
export PATH=$build_path/bin:$PATH
echo "export PATH=$build_path/bin:$PATH" >> ~/.bashrc




# source ~/.bashrc
. ~/.bashrc
