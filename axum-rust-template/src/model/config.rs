use serde_derive::Deserialize;

#[derive(Deserialize,Debug)]
pub struct Config {
    database: Database,
}

#[derive(Deserialize,Debug)]
pub struct Database {
    host: String,
    port: u32,
    username: String,
    password: String,
    db: String,
}
