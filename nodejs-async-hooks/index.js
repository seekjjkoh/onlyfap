const express = require("express");
const { AsyncLocalStorage } = require("async_hooks");

const PORT = 1801;
const asyncLocalStorage = new AsyncLocalStorage();
const app = express();

const generateRandomID = () => {
  return Math.floor(Math.random() * 1_000_000_000);
};

const sleep = (ms) => new Promise((resolve) => {
  setTimeout(() => {
    return resolve(true);
  }, ms);
});

const logger = {
  info: (msg) => {
    const tracingID = asyncLocalStorage.getStore().get("tracingID");
    console.log({ time: new Date(), tracingID, msg });
  },
};

const injectAsyncLocalStorage = (_req, _res, next) => {
  asyncLocalStorage.run(new Map(), () => {
    asyncLocalStorage.getStore().set("tracingID", generateRandomID());
    return next();
  });
};

app.get("/", injectAsyncLocalStorage, async (_req, res) => {
  logger.info("before work");
  await sleep(5_000); // mock heavy work here
  logger.info("after work");
  return res.send("Hello world");
});

app.listen(PORT, () => {
  console.log("Nodejs express server listening at port:", PORT)
});
