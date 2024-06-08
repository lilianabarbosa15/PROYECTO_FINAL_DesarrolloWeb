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
('ref0002', 'manual', 'gasolina', 2023, 'sandero life+', 'azul', 20, 5, 'renault', 'https://cdn.group.renault.com/ren/co/vehicles/sandero/home/sandero03.png.ximg.xsmall.png/99a39699cc.png', 10),
('ref0003', 'automatico', 'gasolina', 2010, 'escape 3.0', 'negro', 16, 5, 'ford', 'https://www.vehiclehistory.com/uploads/2010-Ford-Escape.jpg', 15),
('ref0004', 'automatico', 'gasolina', 2010, 'journey 2.4', 'rojo', 20, 7, 'dodge', 'https://media.ed.edmunds-media.com/dodge/journey/2010/oem/2010_dodge_journey_4dr-suv_rt_fq_oem_2_1600.jpg', 5),
('ref0005', 'automatico', 'electrico', 2019, 'tesla 3', 'gris', 20, 5, 'tesla', 'https://c0.klipartz.com/pngpicture/772/1018/gratis-png-sedan-plateado-vista-lateral-gris-modelo-tesla-3-thumbnail.png', 7),
('ref0006', 'automatico', 'gasolina', 2023, 'corolla 2023', 'gris', 25, 5, 'toyota', 'https://w7.pngwing.com/pngs/198/691/png-transparent-toyota-corolla-altis-v-car-honda-toyota-corolla-altis-v-toyotaaltis-compact-car-sedan-car-thumbnail.png', 15),
('ref0007', 'automatico', 'gasolina', 2022, 'civic 2022', 'rojo', 30,  5, 'honda', 'https://w7.pngwing.com/pngs/453/773/png-transparent-honda-civic-type-r-car-manila-international-auto-show-honda-today-honda-compact-car-sedan-car.png', 5),
('ref0008', 'manual', 'gasolina', 2023, 'ford mustang 2023', 'verde', 20, 5, 'ford', 'https://www.vdm.ford.com/content/dam/vdm_ford/live/en_us/ford/nameplate/mustang/2023/collections/cyp/BYO_Mustang_2023.png',  5),
('ref0009', 'automatico', 'electrico', 2023, 'bolt ev', 'gris', 20,  5, 'chevrolet', 'https://di-uploads-pod30.dealerinspire.com/sweeneycars/uploads/2022/09/2023-Chevrolet-Bolt-EV-Grey-Ghost-Metallic.png', 6),
('ref0010', 'automatico', 'gasolina', 2023, 'bmw x5 2023', 'rojo', 20,  5, 'bmw', 'https://c0.klipartz.com/pngpicture/831/757/gratis-png-ilustracion-bmw-roja-serie-suv-bmw-rojo-x5.png', 16),
('ref0011', 'automatico', 'gasolina', 2023, 'audi a4 2023', 'blanco', 15,  5, 'audi', 'https://e7.pngegg.com/pngimages/215/409/png-clipart-white-audi-coupe-audi-a5-car-audi-a4-audi-a8-audi-car-compact-car-sedan.png', 8),
('ref0012', 'automatico', 'gasolina', 2023, 'e-class', 'negro', 18,  5, 'mercedes-benz', 'https://w7.pngwing.com/pngs/223/220/png-transparent-mercedes-benz-e-class-luxury-vehicle-car-2014-mercedes-benz-s-class-mercedes-compact-car-sedan-car.png', 12),
('ref0013', 'automatico', 'gasolina', 2023, 'sonata', 'rojo', 30,  5, 'hyundai', 'https://c0.klipartz.com/pngpicture/210/698/gratis-png-hyundai-elantra-rojo-2015-hyundai-sonata-car-hyundai-santa-fe-hyundai-tucson-hyundai-sonata-red-car.png', 5),
('ref0014', 'automatico', 'gasolina', 2023, 'sorento', 'rojo', 40,  5, 'kia', 'https://e7.pngegg.com/pngimages/482/259/png-clipart-2018-kia-sorento-2019-kia-sorento-kia-motors-2016-kia-sorento-kia.png', 40);

CREATE TABLE IF NOT EXISTS reservas(
  id SERIAL PRIMARY KEY, --no está en la estructura
  iduser TEXT,
  ref TEXT,
  total INTEGER,
  days INTEGER,
  lifeinsurance BOOLEAN,
  roadassistance BOOLEAN,
  babyseat BOOLEAN,
  luxury BOOLEAN
);

INSERT INTO reservas
(iduser, ref, total, days, lifeinsurance, roadassistance, babyseat, luxury)
VALUES
('Lau153','Ref0001',30,1,false,false,false,false),
('Lau153','Ref0004',40,2,false,false,false,false);
