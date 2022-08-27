const express = require("express");

const app = express();
const port = 8000;

app.get("/", (req, res) => {
  res.status(200).send("Hello There");
});

app.post("/post", (req, res) => {
  const data = req.body
  res.status(200).json(data);
});

app.post("/postform", (req, res) => {
  res.status(200).send("Hello There");
});

app.listen(port, () => {
  console.log(`Listening at https://localhost:${port}`);
});
