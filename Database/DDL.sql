CREATE TABLE IF NOT EXISTS Videojuegos(
	id INTEGER PRIMARY KEY NOT NULL,
	nombre varchar(50) NOT NULL,
	publicado int NOT NULL,
	genero varchar(50) not NULL,
	plataformas varchar(100) NOT NULL,
	desarrollador varchar(50) NOT NULL
);