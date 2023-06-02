with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "detahardd-go-env";
  buildInputs = [
    fpm
    gcc
    gnupg
    go
    nsis
    osslsigncode
    pkgsCross.aarch64-multiplatform.buildPackages.gcc
    pkgsCross.mingwW64.buildPackages.gcc
    rpm
  ];
}
