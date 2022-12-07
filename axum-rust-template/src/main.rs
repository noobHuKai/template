mod apis;
mod init;
mod router;
mod utils;
mod model;

use std::net::SocketAddr;
use tokio::signal;

#[tokio::main]
async fn main() {
    // ========================= init =========================
    init::init().unwrap_or_else(|err| {
        tracing::error!("{}", err);
    });

    return;
    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));

    // server start...
    tracing::info!("listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(router::init_router().into_make_service())
        .with_graceful_shutdown(shutdown_signal())
        .await
        .unwrap();
}

// graceful shutdown
async fn shutdown_signal() {
    let ctrl_c = async {
        signal::ctrl_c()
            .await
            .expect("failed to install Ctrl+C handler");
    };

    #[cfg(unix)]
    let terminate = async {
        signal::unix::signal(signal::unix::SignalKind::terminate())
            .expect("failed to install signal handler")
            .recv()
            .await;
    };

    #[cfg(not(unix))]
    let terminate = std::future::pending::<()>();

    tokio::select! {
        _ = ctrl_c => {},
        _ = terminate => {},
    }

    println!("signal received, starting graceful shutdown");
}
