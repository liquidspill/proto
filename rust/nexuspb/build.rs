fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(false)
        .emit_rerun_if_changed(true)
        .include_file("mod.rs")
        .compile_protos(
            &[
                "../../proto/nexus/catalog/v1/catalog.proto",
                "../../proto/nexus/controlplane/v1/control_plane.proto",
            ],
            &["../../proto"],
        )?;
    Ok(())
}
