fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(false)
        .emit_rerun_if_changed(true)
        .compile_protos(&["../../nexus.v1.proto"], &["../.."])?;
    Ok(())
}
