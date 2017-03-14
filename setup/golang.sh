echo ""
echo "------------------------------------------------"
echo "------------ INSTALLING GOLANG -------------------"
echo "--------------------------------------------------"

cd ~
sudo curl -O https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
sudo tar -xvf go1.7.3.linux-amd64.tar.gz
sudo mv go /usr/local


cd ~

sleep 1

cat > .goenv <<'endmsg1989'
export GOROOT=/usr/local/go
export GOPATH=$HOME/goprojects
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
endmsg1989

sleep 1

cd ~
cat >> .bashrc <<'endmsg1989'
if [ -f ~/.goenv ]; then
    . ~/.goenv
fi
endmsg1989
