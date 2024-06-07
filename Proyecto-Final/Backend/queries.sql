CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nombre TEXT,
    correo TEXT,
    password TEXT
);

INSERT INTO usuarios (nombre, correo, password) 
VALUES
('Carolina', 'carolina@gmail.com', '12234C'),
('Salome', 'salome@gamil.com', '1234'),
('Cristian', 'cristian@gmail.com', '123456');


CREATE TABLE autos (
    id SERIAL PRIMARY KEY,
    tipo VARCHAR(50),
    color VARCHAR(50),
    modelo VARCHAR(50),
    marca VARCHAR(50),
    precio DECIMAL(10, 2),
    disponible BOOLEAN DEFAULT TRUE
);

CREATE TABLE reservas (
    id SERIAL PRIMARY KEY,
    usuario_id INT REFERENCES usuarios(id),
    automovil_id INT REFERENCES autos(id),
    precio_total DECIMAL(10, 2)
);
