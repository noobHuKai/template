use chrono::Local;
use tracing::subscriber::SetGlobalDefaultError;
use tracing::Level;
use tracing_appender::non_blocking::WorkerGuard;
use tracing_subscriber::{
    fmt::{self, format::Writer, time::FormatTime},
    prelude::*,
    Registry,
};

pub fn init_tracing() -> Result<Vec<WorkerGuard>, SetGlobalDefaultError> {
    #[derive(Clone)]
    struct LocalTimer;

    impl FormatTime for LocalTimer {
        fn format_time(&self, w: &mut Writer<'_>) -> std::fmt::Result {
            write!(w, "{}", Local::now().format("%Y-%m-%d %H:%M:%S"))
        }
    }

    let mut guards = Vec::new();
    let error_file_appender = tracing_appender::rolling::daily("logs/error", "error.log");
    let info_file_appender = tracing_appender::rolling::daily("logs/info", "info.log");

    let (info_non_blocking, info_guard) = tracing_appender::non_blocking(info_file_appender);
    let (error_non_blocking, error_guard) = tracing_appender::non_blocking(error_file_appender);
    guards.push(error_guard);
    guards.push(info_guard);

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
                .with_writer(info_non_blocking.with_min_level(Level::WARN)),
        )
        .with(
            fmt::Layer::new()
                .event_format(format)
                .with_ansi(false)
                .with_writer(error_non_blocking.with_max_level(Level::ERROR)),
        );

    tracing::subscriber::set_global_default(subscriber)?;

    Ok(guards)
}
