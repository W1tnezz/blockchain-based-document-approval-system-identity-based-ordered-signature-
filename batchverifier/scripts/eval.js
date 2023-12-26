const fs = require("fs");
const BatchVerifier = artifacts.require("BatchVerifier");

module.exports = async function () {

  let batchVerifier = await BatchVerifier.deployed();
  let type = 1;
  let message = "0x930d2dedab40cb4c03a967aea4f54b22ba6328f7096dc44590e651de6e2a416b";
  let signOrder = [0, 1, 2]

  await batchVerifier.requestSign(type, message, signOrder);

};
