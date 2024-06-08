import { UserLogin, getReservas, getAutomobiles } from "./getInfoUser.js";
import { toggleModal } from "./modal.js";

const sectionReservas = document.querySelector(".Reservas");
const modal = document.querySelector(".modalCart");
const closeButton = document.getElementById("closeModal");

const UsuarioLogueado = JSON.parse(localStorage.getItem("Usu"));
console.log(UsuarioLogueado);

const renderReservas = (reservas, automoviles) => {
    sectionReservas.innerHTML = ""; // Clear previous content

    if (reservas.length === 0) {
        sectionReservas.innerHTML = "<p>No hay autos alquilados</p>";
        return;
    }

    reservas.forEach((reserva) => {
        const car = automoviles.find(auto => auto.Ref === reserva.ref);
        if (car) {
            reserva.car = car; // Attach car details to the reservation object
            const reservaDiv = document.createElement("div");
            reservaDiv.className = "reserva";

            reservaDiv.innerHTML = `
                <img src="${reserva.car.Image}">
                <div class="encabezado">
                    <h4>${reserva.car.Model}</h4>
                    <button class="deleteReserva" data-id="${reserva.ref}" data-user="${reserva.iduser}">Eliminar reserva</button>
                </div>
                <h6>Días: ${reserva.days}</h6>
                <h6>Total: $${reserva.total}</h6>
                <button class="openModal" data-id="${reserva.ref}">Detalles</button>
            `;

            sectionReservas.appendChild(reservaDiv);
        }
    });

    document.querySelectorAll(".openModal").forEach(button => {
        button.addEventListener("click", (event) => {
            const ref = event.target.dataset.id;
            const reserva = reservas.find(reserva => reserva.ref === ref);
            showModal(reserva);
        });
    });

    document.querySelectorAll(".deleteReserva").forEach(button => {
        button.addEventListener("click", async (event) => {
            const ref = event.target.dataset.id;
            const userId = event.target.dataset.user;
            await deleteReserva(userId, ref);
            reservas = reservas.filter(reserva => reserva.ref !== ref);
            renderReservas(reservas, automoviles);
        });
    });
};

const showModal = (reserva) => {
    const sectionDetalles = document.querySelector(".sectionDetalles");
    sectionDetalles.innerHTML = `
        <div class="Caracteristicas">
            <h4>Modelo: ${reserva.car.Model}</h4>
            <h4>Color: ${reserva.car.Color}</h4>
            <h4>Transmisión: ${reserva.car.Transmission}</h4>
            <h4>Combustible: ${reserva.car.Fuel}</h4>
        </div>
        <div class="Reserva">
            <h4>Días: ${reserva.days}</h4>
            <h4>Total: $${reserva.total}</h4>
        </div>
        <div class="ServiciosAdicionales">
            <h4>Seguro de vida: ${reserva.lifeinsurance}</h4>
            <h4>Silla para bebe: ${reserva.babyseat}</h4>
            <h4>Asistencia en carretera: ${reserva.roadassistance}</h4>
            <h4>Equipo de lujo: ${reserva.luxury}</h4>
        </div>
    `;

    modal.classList.add("show");
    modal.classList.remove("hidden");
};

const deleteReserva = async (userId, ref) => {
    try {
        await axios.delete(`http://localhost:3000/reservas`, {
            data: {
                iduser: userId,
                ref: ref
            }
        });
        console.log(`Reservation with ref: ${ref} and userId: ${userId} deleted successfully.`);
    } catch (error) {
        console.error(`Error deleting reservation: ${error}`);
    }
};

closeButton.addEventListener("click", () => {
    modal.classList.add("hidden");
    modal.classList.remove("show");
});

document.addEventListener("DOMContentLoaded", async () => {
    const usuario = await UserLogin(UsuarioLogueado);
    const reservas = await getReservas();
    const automoviles = await getAutomobiles();
    renderReservas(reservas, automoviles);
});
