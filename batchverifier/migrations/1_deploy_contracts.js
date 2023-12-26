const BatchVerifier = artifacts.require("BatchVerifier");
module.exports = function (deployer) {
  deployer.deploy(BatchVerifier);
};
