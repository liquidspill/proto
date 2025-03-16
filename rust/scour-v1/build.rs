fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(true)
        .emit_rerun_if_changed(true)
        .compile_protos(&["../../proto/scour/v1/scour.proto"], &["../.."])?;
    Ok(())
}
