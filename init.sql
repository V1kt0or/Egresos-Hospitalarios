CREATE TABLE IF NOT EXISTS egresosh (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT,
    historia TEXT,
    sexo TEXT,
    edad INTEGER,
    fecha_ingreso TEXT,
    fecha_egreso TEXT,
    ubigeo_lugar_residencia INTEGER,
    lugar_residencia TEXT,
    fecha_corte TEXT
);