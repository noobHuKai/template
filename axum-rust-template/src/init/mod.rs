mod config;
mod tracing;

use anyhow::Result;

pub fn init() -> Result<()> {
    tracing::init_tracing()?;

    config::init_configuration()?;
    Ok(())
}
