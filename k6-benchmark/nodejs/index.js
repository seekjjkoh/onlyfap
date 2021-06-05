const express = require("express");

const PORT = 1801;

const app = express();

app.get("/", (_req, res) => {
  return res.send("Hello world");
});

app.listen(PORT, () => {
  console.log("Nodejs express server listening at port:", PORT)
});
