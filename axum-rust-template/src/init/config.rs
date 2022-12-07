use crate::model::Config;
use anyhow::{Ok, Result};

pub fn init_configuration() -> Result<()> {
    let content = std::fs::read_to_string("config/config.toml")?;

    let config: Config = toml::from_str(&content)?;
    println!("{:?}", config);
    Ok(())
}
