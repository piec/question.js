import express from "express";
import { setIntervalAsync } from "set-interval-async/fixed";

function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

// only a single execution of this function is allowed at a time
// which is not the case with the current code
async function task(reason: string) {
  console.log("do thing because %s...", reason);
  await sleep(1000);
  console.log("done");
}

const app = express();

// call task regularly
setIntervalAsync(async () => {
  await task("ticker");
}, 5000);

// call task immediately
app.get("/task", async (req, res) => {
  await task("trigger");
  res.send("ok");
});

const port = 3000;
app.listen(port, () => {
  console.log(`listen on ${port}`);
});
