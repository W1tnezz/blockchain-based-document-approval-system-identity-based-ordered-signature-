const fs = require("fs");
const Registry = artifacts.require("Registry");

module.exports = async function () {

  let registry = await Registry.deployed();
  let type = 2;  // 0 : sakai, 1 : IBSAS, 2 : NOTBATCH
  let message = "0x930d2dedab40cb4c03a967aea4f54b22ba6328f7096dc44590e651de6e2a416b";
  let signOrder = [0, 1, 2, 3, 4, 5, 6]

  await registry.requestSign(type, message, signOrder);

};
