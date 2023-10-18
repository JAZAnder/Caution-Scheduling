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

importElements('readOnlyTable', './assets/elements/readOnlyTable.html');

