{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.docker
    pkgs.docker-compose
    pkgs.go
    pkgs.python3
    pkgs.lxc
  ];

  shellHook = ''
    export GOPATH=''${GOPATH:-''$HOME/go}
    export PATH=''${PATH}:''${GOPATH}/bin
  '';
}

