const express = require('express');
const path = require('path');
const app = express();
const port = 3000;

// Middleware para servir archivos estáticos desde el directorio 'public'
app.use(express.static(path.join(__dirname, "assets")));

// Ruta para servir tu archivo HTML
app.get("/", (req: any, res: any) => {
  res.sendFile(path.join(__dirname, "index.html"));
});

// Inicia el servidor
app.listen(port, () => {
  console.log(`Servidor Express escuchando en http://localhost:${port}`);
});
