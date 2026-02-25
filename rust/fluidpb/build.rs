fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(false)
        .emit_rerun_if_changed(true)
        .include_file("mod.rs")
        .compile_protos(&["../../proto/fluid/rpc/v1/rpc.proto"], &["../../proto"])?;
    Ok(())
}
