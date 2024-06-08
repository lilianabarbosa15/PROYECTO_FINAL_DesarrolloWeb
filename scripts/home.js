import { UserLogin, getAutomobiles } from "./getInfoUser.js";
import {
    toggleModal,
}from "./modal.js";
const UsuarioLogueado = JSON.parse(localStorage.getItem("Usu"))
console.log(UsuarioLogueado);

const spanAside = document.querySelector(".nameUser");
const botonBuscar = document.querySelector(".fa-magnifying-glass");
const wordsearch = document.querySelector("#wordsearch");
const options = document.querySelectorAll(".options1");
const resultados = document.querySelector(".resultadoBusqueda");
const modal = document.querySelector(".modalCart");
const closeButton = document.getElementById("closeModal");
let datos;
let selectedOption = null;
console.log(botonBuscar)
console.log(spanAside)

const InsertarAficheByType = (array, searchTerm, tipo) => {
    let filtro;
    if (tipo == "Transmisión") {
        filtro = array.filter(carro =>
            carro.Type_transmission.toLowerCase().includes(searchTerm.toLowerCase())
        );
    } else if (tipo == "Combustible") {
        filtro = array.filter(carro =>
            carro.Type_fuel.toLowerCase().includes(searchTerm.toLowerCase())
        );
    } else if (tipo == "Modelo") {
        filtro = array.filter(carro =>
            carro.Model.toLowerCase().includes(searchTerm.toLowerCase())
        );
    } else if (tipo == "Precio") {
        filtro = array.filter(carro =>
            carro.Price.toLowerCase().includes(searchTerm.toLowerCase())
        );
    } else if (tipo == "Ninguno") {
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
    afiche.innerHTML =
        `<div id="${carro.Ref}">
      <img class="reservar" src=${carro.Image}>
      <div class="Principal"> 
      <h4>${capitalizeFirstLetter(carro.Model)}</h4>
      <h6> $ ${carro.Price}</h6>
      </div>
      <h6 class="solito"> Transmisión: ${capitalizeFirstLetter(carro.Type_transmission)}</h6>
      <div class="secundario"> 
      <h6> Combustible: ${capitalizeFirstLetter(carro.Type_fuel)}</h6>
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
    const figura = document.createElement('figure');
    const imagen = document.createElement('img');
    imagen.src = urlimg;
    imagen.alt = 'Imagen';
    figura.appendChild(imagen);
    fotoDetalles.appendChild(figura);
   
    
};

const insertarDetallesModal = (carro) => {
    const detallesEspecificos = document.querySelector(".detallesEspecificos");
    detallesEspecificos.innerHTML = `
        <div>
            <h6>Modelo: ${capitalizeFirstLetter(carro.Model)}</h6>
            <h6>Color: ${capitalizeFirstLetter(carro.Color)}</h6>
            <h6>Sillas: ${carro.Seats}</h6>
            <h6>Combustible: ${capitalizeFirstLetter(carro.Type_fuel)}</h6>
            <h6>Transmisión: ${capitalizeFirstLetter(carro.Type_transmission)}</h6>
        </div>
        <div>
            <h6>Marca: ${capitalizeFirstLetter(carro.Brand)}</h6>
            <h6>Año: ${carro.Year}</h6>
            <h6>Precio/día: $${carro.Price}</h6>
        </div>`;
};


document.addEventListener("DOMContentLoaded", async () => {

    datos = await UserLogin(UsuarioLogueado);
    const carritos = await getAutomobiles();
    console.log(carritos)
    const usuario = datos[0];
    spanAside.textContent = usuario.Name


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
            console.log(searchTerm)
            console.log(tipo)
            options.forEach(opt => {
                opt.classList.remove('disabled');
            });
            const filteredCars = InsertarAficheByType(carritos, searchTerm, tipo);
            console.log(filteredCars)
            resultados.innerHTML = ""; // Limpiar resultados anteriores
            filteredCars.forEach((elemento) => {
                InsertarCarro(resultados, elemento);
            });
            
            

            const aficheCarrito = document.querySelectorAll(".AficheCarro img.reservar");
            aficheCarrito.forEach(imagen => {
                imagen.addEventListener("click", (event) => {
                    event.stopPropagation();
                    const carro = filteredCars.find(carro => carro.Image === imagen.src);
                    console.log(carro.Image)
                    insertarImagen(carro.Image);
                    insertarDetallesModal(carro);
                    modal.classList.add("show");
                    modal.classList.remove("hidden");
                });
            });

            closeButton.addEventListener("click", () => {
                modal.classList.add("hidden");
                modal.classList.remove("show");
            });
            
        }
    });
    







});