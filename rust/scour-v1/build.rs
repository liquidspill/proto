fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(true)
        .emit_rerun_if_changed(true)
        .compile_protos(&["../../scour.v1.proto"], &["../.."])?;
    Ok(())
}
