use anyhow::Result;
use chrono::Local;
use tracing::Level;
use tracing_subscriber::{
    fmt::{self, format::Writer, time::FormatTime},
    prelude::*,
    Registry,
};

pub fn init_tracing() -> Result<()> {
    #[derive(Clone)]
    struct LocalTimer;

    impl FormatTime for LocalTimer {
        fn format_time(&self, w: &mut Writer<'_>) -> std::fmt::Result {
            write!(w, "{}", Local::now().format("%Y-%m-%d %H:%M:%S"))
        }
    }

    let info_file_appender =
        tracing_appender::rolling::daily("logs/info", "info.log").with_min_level(Level::WARN);
    let error_file_appender =
        tracing_appender::rolling::daily("logs/error", "error.log").with_max_level(Level::ERROR);

    let all_files = info_file_appender.and(error_file_appender);

    let format = fmt::format()
        .with_line_number(true)
        .with_file(true)
        .with_target(false)
        .with_timer(LocalTimer);

    let subscriber = Registry::default()
        .with(
            fmt::Layer::new()
                .event_format(format.clone())
                .with_writer(std::io::stdout),
        )
        .with(
            fmt::Layer::new()
                .event_format(format.clone())
                .with_ansi(false)
                .with_writer(all_files),
        );

    tracing::subscriber::set_global_default(subscriber)?;

    Ok(())
}
