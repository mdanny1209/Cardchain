const http = require('http');
const express = require('express');
const { mainRouter } = require('./routers');

const PORT = 4500
const app = express();
app.use('/api/', mainRouter);

const server = http.createServer(app);
server.listen(PORT, () => {
    console.log(`Server started on port ${PORT}`);
});
