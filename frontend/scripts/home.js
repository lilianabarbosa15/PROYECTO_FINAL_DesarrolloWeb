import { UserLogin, getAutomobiles } from "./getInfoUser.js";
import { toggleModal } from "./modal.js";


const UsuarioLogueado = JSON.parse(localStorage.getItem("Usu"));
const spanAside = document.querySelector(".nameUser");
const botonBuscar = document.querySelector(".fa-magnifying-glass");
const wordsearch = document.querySelector("#wordsearch");
const options = document.querySelectorAll(".options1");
const resultados = document.querySelector(".resultadoBusqueda");
const modal = document.querySelector(".modalCart");
const closeButton = document.getElementById("closeModal");
const botonReservar = document.querySelector(".botonReservar");
let datos;
let selectedOption = null;

const InsertarAficheByType = (array, searchTerm, tipo) => {
  let filtro;
  if (tipo === "Transmisión") {
    filtro = array.filter(carro =>
      carro.Transmission.toLowerCase().includes(searchTerm.toLowerCase())
    );
  } else if (tipo === "Combustible") {
    filtro = array.filter(carro =>
      carro.Fuel.toLowerCase().includes(searchTerm.toLowerCase())
    );
  } else if (tipo === "Modelo") {
    filtro = array.filter(carro =>
      carro.Model.toLowerCase().includes(searchTerm.toLowerCase())
    );
  } else if (tipo === "Precio") {
    filtro = array.filter(carro =>
      carro.Price.toLowerCase().includes(searchTerm.toLowerCase())
    );
  } else if (tipo === "Ninguno") {
    filtro = array;
  }
  return filtro;
};

function capitalizeFirstLetter(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

const InsertarCarro = (contenedor, carro) => {
  const afiche = document.createElement("div");
  afiche.className = 'AficheCarro';
  afiche.innerHTML = `
    <div id="${carro.Ref}">
      <img class="reservar" src=${carro.Image}>
      <div class="Principal">
        <h4>${capitalizeFirstLetter(carro.Model)}</h4>
        <h6> $ ${carro.Price}</h6>
      </div>
      <h6 class="solito"> Transmisión: ${capitalizeFirstLetter(carro.Transmission)}</h6>
      <div class="secundario">
        <h6> Combustible: ${capitalizeFirstLetter(carro.Fuel)}</h6>
        <h6> Marca: ${capitalizeFirstLetter(carro.Brand)}</h6>
      </div>
      <div class="secundario">
        <h6> Color: ${capitalizeFirstLetter(carro.Color)}</h6>
        <h6> Año: ${carro.Year}</h6>
      </div>
    </div>`;
  contenedor.appendChild(afiche);
};

const insertarImagen = (urlimg) => {
  const fotoDetalles = document.querySelector(".sectionDetalles");
  if (!fotoDetalles) return; // Ensure the element exists
  fotoDetalles.innerHTML = ""; // Clear previous image
  const figura = document.createElement('figure');
  const imagen = document.createElement('img');
  imagen.src = urlimg;
  imagen.alt = 'Imagen';
  figura.appendChild(imagen);
  fotoDetalles.appendChild(figura);
};

const insertarDetallesModal = (carro) => {
  const detallesEspecificos = document.querySelector(".detallesEspecificos");
  if (!detallesEspecificos) return; // Ensure the element exists
  detallesEspecificos.innerHTML = `
    <div>
      <h6>Modelo: ${capitalizeFirstLetter(carro.Model)}</h6>
      <h6>Color: ${capitalizeFirstLetter(carro.Color)}</h6>
      <h6>Sillas: ${carro.Seats}</h6>
      <h6>Combustible: ${capitalizeFirstLetter(carro.Fuel)}</h6>
      <h6>Transmisión: ${capitalizeFirstLetter(carro.Transmission)}</h6>
    </div>
    <div>
      <h6>Marca: ${capitalizeFirstLetter(carro.Brand)}</h6>
      <h6>Año: ${carro.Year}</h6>
      <h6>Precio/día: $${carro.Price}</h6>
    </div>`;
};

const postReserva = async (reserva) => {
  try {
    const response = await axios.post('http://localhost:3000/reservas', reserva);
    console.log('Reserva creada:', response.data);
  } catch (error) {
    console.error('Error al crear la reserva:', error);
  }
};

document.addEventListener("DOMContentLoaded", async () => {
  datos = await UserLogin(UsuarioLogueado);
  const carritos = await getAutomobiles();
  const usuario = datos[0];
  spanAside.textContent = usuario.name;

  options.forEach(option => {
    option.addEventListener("click", (event) => {
      event.preventDefault();
      selectedOption = event.target.textContent;

      options.forEach(opt => {
        if (opt !== option) {
          opt.classList.add('disabled');
        }
      });
    });

    if (selectedOption) {
      botonBuscar.style.display = 'flex';
      botonBuscar.style.opacity = 1;
    } else {
      botonBuscar.style.opacity = 0.8;
    }
  });

  let searchTerm;
  let tipo;
  botonBuscar.addEventListener("click", async () => {
    if (wordsearch.value && selectedOption) {
      searchTerm = wordsearch.value;
      tipo = selectedOption;

      options.forEach(opt => {
        opt.classList.remove('disabled');
      });

      const filteredCars = InsertarAficheByType(carritos, searchTerm, tipo);
      resultados.innerHTML = "";
      filteredCars.forEach((elemento) => {
        InsertarCarro(resultados, elemento);
      });

      const aficheCarrito = document.querySelectorAll(".AficheCarro img.reservar");
      let carro;
      aficheCarrito.forEach(imagen => {
        imagen.addEventListener("click", (event) => {
          event.stopPropagation();
          const carro = filteredCars.find(carro => carro.Image === imagen.src);
          if (carro) {
            insertarImagen(carro.Image);
            insertarDetallesModal(carro);
            modal.classList.add("show");
            modal.classList.remove("hidden");

            botonReservar.onclick = () => {
              const days = parseInt(document.getElementById("days").value);
              const lifeinsurance = document.getElementById("mySeguroVida").checked;
              const roadassistance = document.getElementById("Asistencia").checked;
              const babyseat = document.getElementById("SillaBebe").checked;
              const luxury = document.getElementById("EquipoLujoso").checked;

              const reserva = {
                iduser: usuario.username,
                ref: carro.Ref,
                total: days * carro.Price,
                days: days,
                lifeinsurance: lifeinsurance,
                roadassistance: roadassistance,
                babyseat: babyseat,
                luxury: luxury
              };

              postReserva(reserva);

              modal.classList.add("hidden");
              modal.classList.remove("show");
            };
          }
        });
      });

      closeButton.addEventListener("click", () => {
        modal.classList.add("hidden");
        modal.classList.remove("show");
      });
    }
  });
});
