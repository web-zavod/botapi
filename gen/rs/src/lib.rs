pub mod healthcheck {
    include!("botapi.healthcheck.v1.rs");
}

pub mod expense {
    include!("botapi.expense.v1.rs");
}

pub mod media {
    include!("botapi.media.v1.rs");
}

pub mod command {
    include!("botapi.command.v1.rs");
}
