function createTable(rows, columns) {
    const table = document.createElement('table');
    table.border = '1';

    for (let i = 0; i < rows; i++) {
        const row = document.createElement('tr');

        for (let j = 0; j < columns; j++) {
            const cell = document.createElement('td');
            cell.textContent = `Row ${i + 1}, Col ${j + 1}`;
            row.appendChild(cell);
        }

        table.appendChild(row);
    }

    return table;
}

// Create the table
const tableElement = createTable(5, 9);

// Append the table to a container div with the id 'table-container'
const container = document.getElementById('table-container');
container.appendChild(tableElement);
