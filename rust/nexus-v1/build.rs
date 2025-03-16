fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(false)
        .emit_rerun_if_changed(true)
        .compile_protos(&["../../proto/nexus/v1/nexus.proto"], &["../.."])?;
    Ok(())
}
