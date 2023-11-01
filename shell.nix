{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  buildInputs = [
    delve
    go
    gocode
    golangci-lint
    gomodifytags
    gopls
    gore
    gotests
    gotools
    nodejs
    rnix-lsp
    shellcheck
  ];
  shellHook = ''
    mkdir -p .go/tmp
  '';
}
