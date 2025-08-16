fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(false)
        .emit_rerun_if_changed(true)
        .include_file("mod.rs")
        .compile_protos(
            &["../../proto/nexus/metadata/v1/metadata.proto"],
            &["../../proto"],
        )?;
    Ok(())
}
