const apiUrlMap = {
    alunos: 'http://localhost:8080/alunos',
    atividades: 'http://localhost:8080/atividades',
    notas: 'http://localhost:8080/notas',
    professores: 'http://localhost:8080/professores',
    turmas: 'http://localhost:8080/turmas'
};

const apiKey = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InZlcmJndWpqeGdmY3Fla2p6aXdlIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjQzMzY0MjAsImV4cCI6MjAzOTkxMjQyMH0.XBYdHu8DR3qaBOv9LQhhLMoNV6SSEGTC3coaZ5OHVwI'; // Se estiver usando Supabase
const headers = {
    'Content-Type': 'application/json',
    'apikey': apiKey,
    'Authorization': `Bearer ${apiKey}`
};

document.addEventListener('DOMContentLoaded', () => {
    if (document.body.id) {
        const model = document.body.id;
        loadTable(model);
        const form = document.getElementById(`${model}Form`);
        if (form) {
            form.addEventListener('submit', (event) => {
                event.preventDefault();
                const data = new FormData(form);
                const jsonData = {};
                data.forEach((value, key) => {
                    jsonData[key] = value;
                });
                createItem(model, jsonData);
            });
        }
    }
});

function loadTable(model) {
    fetch(apiUrlMap[model], { headers })
        .then(response => response.json())
        .then(data => {
            const tableBody = document.getElementById(`${model}Table`).querySelector('tbody');
            tableBody.innerHTML = '';
            data.forEach(item => {
                const row = document.createElement('tr');
                Object.keys(item).forEach(key => {
                    const cell = document.createElement('td');
                    cell.textContent = item[key];
                    row.appendChild(cell);
                });
                const actionsCell = document.createElement('td');
                actionsCell.innerHTML = `
                    <button onclick="editItem('${model}', ${item.id})">Editar</button>
                    <button onclick="deleteItem('${model}', ${item.id})">Deletar</button>
                `;
                row.appendChild(actionsCell);
                tableBody.appendChild(row);
            });
        })
        .catch(error => console.error('Erro ao carregar tabela:', error));
}

function createItem(model, data) {
    fetch(apiUrlMap[model], {
        method: 'POST',
        headers,
        body: JSON.stringify(data)
    })
    .then(() => loadTable(model))
    .catch(error => console.error('Erro ao criar item:', error));
}

function editItem(model, id) {
    const data = prompt('Digite os novos dados em formato JSON:');
    if (data) {
        fetch(`${apiUrlMap[model]}/${id}`, {
            method: 'PUT',
            headers,
            body: JSON.stringify(JSON.parse(data))
        })
        .then(() => loadTable(model))
        .catch(error => console.error('Erro ao atualizar item:', error));
    }
}

function deleteItem(model, id) {
    fetch(`${apiUrlMap[model]}/${id}`, {
        method: 'DELETE',
        headers
    })
    .then(() => loadTable(model))
    .catch(error => console.error('Erro ao deletar item:', error));
}
