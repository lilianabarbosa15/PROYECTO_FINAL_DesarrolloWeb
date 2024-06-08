

import { getAllUsers } from "./getInfoUser.js";

const entradaNombre = document.querySelector("#usuario");
const entradaContraseña = document.querySelector("#password");
const form = document.getElementById("form");
const errorUsuario = document.getElementById("error-usuario");
const errorPassword = document.getElementById("error-password");
let user;


const validateUser = (usuarios, username, password) => {
    return usuarios.find(user => user.Usu === username && user.Password === password);
};
document.addEventListener("DOMContentLoaded", async () => {
    
    user = await getAllUsers();


    
    form.addEventListener("submit", async (event) => {

        event.preventDefault();
        const userRegExp = /^[a-zA-Z0-9]{4,}$/;
        const passwordRegExp = /^[a-zA-Z0-9]{4,}$/;


        if (!userRegExp.test(entradaNombre.value.trim())) {
            errorUsuario.textContent = "El campo de usuario debe tener más de 3 caracteres y contener números o letras.";
            errorUsuario.setAttribute('style', 'display:flex');
        }

        if (!passwordRegExp.test(entradaContraseña.value.trim())) {
            errorPassword.textContent = "El campo de contraseña debe tener al menos 4 caracteres, contener números o letras.";
            errorPassword.setAttribute('style', 'display:flex');
        }
        if (userRegExp.test(entradaNombre.value.trim()) && passwordRegExp.test(entradaContraseña.value.trim())) {
            const usuarioLogueado = validateUser(user,entradaNombre.value, entradaContraseña.value);
            
            if (usuarioLogueado) {
                Swal.fire({
                    title: 'Login Exitoso',
                    text: '¡Bienvenido!',
                    icon: 'success'
                  });
                localStorage.setItem("Usu", JSON.stringify(entradaNombre.value));
                location.href= '../pages/home.html';
                
            } else {
                Swal.fire({
                    title: 'Error',
                    text: 'Usuario o contraseña incorrectos',
                    icon: 'error'
                  });
                
            }

        }




    });

});