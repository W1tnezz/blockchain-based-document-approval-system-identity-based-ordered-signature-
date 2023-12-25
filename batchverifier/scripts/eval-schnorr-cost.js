const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const OracleContract = artifacts.require("OracleContract");

module.exports = async function () {

  let oracleContract = await OracleContract.deployed();
  let tx = "0x930d2dedab40cb4c03a967aea4f54b22ba6328f7096dc44590e651de6e2a416b";
  let size = 10;   // 总量阈值
  let minRank = 5;  // 个人阈值
  let fee = await oracleContract.totalFee(size);

  await oracleContract.validateTransaction(tx, size, minRank,{
    value: fee,
  });

};
