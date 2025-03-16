@default:
  just --list

generate:
  cp proto/flow/v1/flow.proto go/flow/v1/flow.proto
  buf generate
  protoc -I /usr/local/include \
  -I . \
  -I /Users/senyo/go/bin \
  -I /Users/senyo/.cache/buf/v3/modules/b5/buf.build/srikrsna/protoc-gen-gotag/7a85d3ad2e7642c198480e92bf730c14/files \
  --gotag_out=paths=source_relative,outdir=:. go/flow/v1/flow.proto 
  rm go/flow/v1/flow.proto