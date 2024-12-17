document.addEventListener("DOMContentLoaded", () => {
    const apiUrl = "http://localhost:8080/clientes";
    
    // Obtención de los clientes
    async function getClientes() {
        const response = await fetch(apiUrl);
        const clientes = await response.json();
        renderClientes(clientes);
    }

    // Mostrar clientes en la tabla
    function renderClientes(clientes) {
        const tableBody = document.querySelector("#clientesTable tbody");
        tableBody.innerHTML = ""; // Limpiar tabla

        clientes.forEach(cliente => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${cliente.id}</td>
                <td>${cliente.nombre}</td>
                <td>${cliente.correo_electronico}</td>
                <td>${cliente.numero_telefono}</td>
                <td>
                    <button onclick="deleteCliente(${cliente.id})">Eliminar</button>
                    <button onclick="updateCliente(${cliente.id})">Actualizar</button>
                </td>
            `;
            tableBody.appendChild(row);
        });
    }

    // Crear un cliente
    document.querySelector("#addClienteForm").addEventListener("submit", async (e) => {
        e.preventDefault();

        const nombre = document.querySelector("#nombre").value;
        const correo = document.querySelector("#correo").value;
        const telefono = document.querySelector("#telefono").value;

        const newCliente = { nombre, correo_electronico: correo, numero_telefono: telefono };

        const response = await fetch(apiUrl, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newCliente),
        });

        const result = await response.json();
        if (response.status === 201) {
            alert("Cliente agregado correctamente");
            getClientes(); // Refrescar la lista de clientes
        } else {
            alert("Error al agregar cliente");
        }
    });

    // Eliminar un cliente
    async function deleteCliente(id) {
        const response = await fetch(`${apiUrl}/${id}`, { method: "DELETE" });
        if (response.status === 204) {
            alert("Cliente eliminado");
            getClientes(); // Refrescar la lista
        } else {
            alert("Error al eliminar cliente");
        }
    }

    // Actualizar un cliente
    async function updateCliente(id) {
        const nombre = prompt("Nuevo nombre:");
        const correo = prompt("Nuevo correo electrónico:");
        const telefono = prompt("Nuevo número de teléfono:");

        const updatedCliente = { nombre, correo_electronico: correo, numero_telefono: telefono };

        const response = await fetch(`${apiUrl}/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(updatedCliente),
        });

        if (response.ok) {
            alert("Cliente actualizado");
            getClientes(); // Refrescar la lista
        } else {
            alert("Error al actualizar cliente");
        }
    }

    // Llamar a getClientes al cargar la página
    getClientes();
});
