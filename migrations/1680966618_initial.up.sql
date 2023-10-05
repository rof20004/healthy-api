CREATE TABLE IF NOT EXISTS pacientes(
    id         text        not null,
    nome       text        not null,
    avatar     text        not null,
    idade      integer     not null,
    created_at timestamptz not null,
    primary key(id)
);

CREATE TABLE IF NOT EXISTS profissionais(
    id         text        not null,
    nome       text        not null,
    cpf        text        not null,
    email      text        not null,
    foto       text        not null,
    crp        text        not null,
    senha      text        not null,
    created_at timestamptz not null,
    primary key(id)
);

CREATE TABLE IF NOT EXISTS profissional_agenda(
    id              text        not null,
    profissional_id text        not null,
    data            timestamptz not null,
    created_at      timestamptz not null,
    primary key(id)
);

CREATE TABLE IF NOT EXISTS consultas(
    id              text        not null,
    paciente_id     text        not null,
    profissional_id text        not null,
    data            timestamptz not null,
    created_at      timestamptz not null,
    primary key(id)
);
