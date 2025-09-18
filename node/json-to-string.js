const fs = require('fs');

// Read the JSON file
fs.readFile('data.json', 'utf8', (err, data) => {
    if (err) {
        console.error('Error reading the file:', err);
        return;
    }

    // Parse the JSON data
    try {
        const jsonData = JSON.parse(data);
        var jsonStr = JSON.stringify(jsonData)
        jsonStr = jsonStr.replace(/"/g, '\\"')
        // jsonStr = jsonStr.replace("role_name", "rolliePoly")
        console.log(jsonStr);
    } catch (parseErr) {
        console.error('Error parsing JSON:', parseErr);
    }
});