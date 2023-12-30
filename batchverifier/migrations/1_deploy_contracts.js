const Registry = artifacts.require("Registry");
const Sakai = artifacts.require("Sakai");
const IBSAS = artifacts.require("IBSAS");
module.exports = function (deployer) {
    deployer.deploy(Registry).then(
      function (){
        return deployer.deploy(Sakai, Registry.address);
      }
    ).then(
      function (){
        return deployer.deploy(IBASA, Registry.address);
      }
    );
};
