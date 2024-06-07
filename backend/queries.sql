CREATE TABLE IF NOT EXISTS users(
  usu TEXT PRIMARY KEY,
  name TEXT,
  email TEXT,
  password TEXT
);

INSERT INTO users 
(usu, name, email, password) VALUES
('lilo', 'Laura', 'lau.gomez@gmail.com', '15647k'),
('Lau153','Laura Gomez','lau.gomez@gmail.com','15647k'),
('Lily1526','Lily Betancur','lily_vetancur@hotmail.com','1234'),
('li6lo2','Lily Betancur1', 'lily_betancur@hotmail.com', '1234'),
('l5ilo2','Lily Betanc2ur', 'lily_betancur@hotmail.com', '1234'),
('lilo42','ily Betancu3r', 'lily_betancur@hotmail.com', '1234'),
('lil3o2','Lily Bet4ancur', 'lily_betancur@hotmail.com', '1234'),
('li2lo2','Lily Beta5ncur', 'lily_betancur@hotmail.com', '1234'),
('l1ilo2','Lily Betan6cur', 'lily_betancur@hotmail.com', '1234');

CREATE TABLE IF NOT EXISTS automobiles(
  ref TEXT PRIMARY KEY,
	type_transmission TEXT,
	type_fuel TEXT,
	year INT,
	model TEXT,
	color TEXT,
	price INT,
	seats INT,
	brand TEXT,
	image TEXT,
	quantity INT
);

INSERT INTO automobiles 
(ref, type_transmission, type_fuel, year, model, color, price, seats, brand, image, quantity)
VALUES
('ref0001', 'automatico', 'gasolina', 2024, 'mazda3', 'rojo', 30, 5, 'mazda', 'https://www.mazda.com.co/globalassets/cars/mazda-3-2024/versiones-y-menu/desplegable-vehiculos.png', 20),
('ref0002', 'manual', 'gasolina', 2023, 'sandero life+', 'azul', 20, 5, 'renault', 'https://acroadtrip.blob.core.windows.net/catalogo-imagenes/l/RT_V_9e1b416d37674588b9959680073c1e92.webp', 10),
('ref0003', 'automatico', 'gasolina', 2010, 'escape 3.0', 'negro', 16, 5, 'ford', 'https://www.vehiclehistory.com/uploads/2010-Ford-Escape.jpg', 15),
('ref0004', 'automatico', 'gasolina', 2010, 'journey 2.4', 'rojo', 20, 7, 'dodge', 'https://media.ed.edmunds-media.com/dodge/journey/2010/oem/2010_dodge_journey_4dr-suv_rt_fq_oem_2_1600.jpg', 5);
