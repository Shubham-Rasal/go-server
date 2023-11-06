const http = require('http');
const fs = require('fs');
const csv = require('csv-parser');

const port = 8080;

const server = http.createServer((req, res) => {
  if (req.url === '/') {
    const results = [];

    fs.createReadStream('prices.csv')
      .pipe(csv())
      .on('data', (row) => {
        // console.log(row)
        results.push({
            Date: row.Date,
            Close: parseFloat(row[' Close/Last'].trim().replace('$', '')),
            Volume: parseInt(row[' Volume'].trim()),
            Open: parseFloat(row[' Open'].trim().replace('$', '')),
            High: parseFloat(row[' High'].trim().replace('$', '')),
            Low: parseFloat(row[' Low'].trim().replace('$', '')),
          });
      })
      .on('end', () => {
        res.setHeader('Content-Type', 'application/json');
        res.end(JSON.stringify(results));
      });
  } else {
    res.statusCode = 404;
    res.end('Not Found');
  }
});

server.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
