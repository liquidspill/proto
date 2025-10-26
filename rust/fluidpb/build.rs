fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(true)
        .emit_rerun_if_changed(true)
        .include_file("mod.rs")
        .compile_protos(
            &["../../proto/fluid/agent/v1/agent.proto"],
            &["../../proto"],
        )?;
    Ok(())
}
