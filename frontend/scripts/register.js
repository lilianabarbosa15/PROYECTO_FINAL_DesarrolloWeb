const form = document.getElementById("form");
const entradaNombre = document.querySelector("#name");
const entradaEmail = document.querySelector("#email");
const entradaNombreUsuario = document.querySelector("#nameUsuario");
const entradaContraseña = document.querySelector("#Contraseña");
const errorName = document.getElementById("error-name");
const errorEmail = document.getElementById("error-email");
const errorNameUsuario = document.getElementById("error-nameUsuario");
const errorContraseña = document.getElementById("error-Contraseña");
let valid;
document.addEventListener("DOMContentLoaded", async () => {


    form.addEventListener("submit", async (event) => {
        event.preventDefault();



        const nameRegExp = /^[a-zA-Z0-9\s]{4,}$/; // Al menos 4 caracteres, incluyendo letras y números
        const emailRegExp = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/; // Formato de correo electrónico
        const userRegExp = /^[a-zA-Z0-9]{4,}$/; // Al menos 4 caracteres, incluyendo letras y números
        const passwordRegExp = /^[a-zA-Z0-9]{4,}$/; // Al menos 4 caracteres, incluyendo letras y números

        // Validate full name
        if (!nameRegExp.test(entradaNombre.value.trim())) {
            errorName.textContent = "El nombre debe tener al menos 4 caracteres y contener solo letras.";
            errorName.setAttribute('style', 'display:flex');

        }

        // Validate email
        if (!emailRegExp.test(entradaEmail.value.trim())) {
            errorEmail.textContent = "Ingrese un correo electrónico válido.";
            errorEmail.setAttribute('style', 'display:flex');

        }

        // Validate username
        if (!userRegExp.test(entradaNombreUsuario.value.trim())) {
            errorNameUsuario.textContent = "El nombre de usuario debe tener al menos 4 caracteres y contener solo números y letras.";
            errorNameUsuario.setAttribute('style', 'display:flex');

        }

        // Validate password
        if (!passwordRegExp.test(entradaContraseña.value.trim())) {
            errorContraseña.textContent = "La contraseña debe tener al menos 4 caracteres y contener solo números y letras.";
            errorContraseña.setAttribute('style', 'display:flex');

        }

        if (nameRegExp.test(entradaNombre.value.trim()) && emailRegExp.test(entradaEmail.value.trim()) && userRegExp.test(entradaNombreUsuario.value.trim()) && passwordRegExp.test(entradaContraseña.value.trim())) {
            const newUser = {
                username: entradaNombreUsuario.value.trim(),
                name: entradaNombre.value.trim(),
                email: entradaEmail.value.trim(),
                password: entradaContraseña.value.trim()
            };

            try {
                const response = await fetch('http://localhost:3000/users', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(newUser)
                });

                if (response.ok) {
                    Swal.fire({
                        title: 'Registro Exitoso',
                        text: '¡Tu cuenta ha sido creada!',
                        icon: 'success'
                    });
                    location.href= '../index.html';
                } else {
                    throw new Error('Error en el registro');
                }
            }catch (error) {
                console.error('Error:', error);
                Swal.fire({
                    title: 'Error',
                    text: 'Hubo un problema al crear tu cuenta. Inténtalo de nuevo.',
                    icon: 'error'
                });
            }
        }
    })
})

