import { UserLogin, getAutomobiles } from "./getInfoUser.js";
const UsuarioLogueado = JSON.parse(localStorage.getItem("Usu"))
console.log(UsuarioLogueado);

const spanAside = document.querySelector(".nameUser");
console.log(spanAside)
let datos;

document.addEventListener("DOMContentLoaded", async () => {
   
    datos = await UserLogin(UsuarioLogueado);
    const carritos = await getAutomobiles();
    console.log(carritos)
    const usuario = datos[0];
    spanAside.textContent = usuario.Name


})