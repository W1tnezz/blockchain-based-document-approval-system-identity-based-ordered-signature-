const BN256G1 = artifacts.require("BN256G1");
const Registry = artifacts.require("Registry");
const Sakai = artifacts.require("Sakai");
const IBSAS = artifacts.require("IBSAS");
module.exports = function (deployer) {
    deployer.deploy(Registry).then(
      function(){
        deployer.deploy(BN256G1);
      }
    )
    
    deployer.link(BN256G1, Sakai)
    deployer.deploy(Sakai, Registry.address);

    deployer.link(BN256G1, IBSAS)
    deployer.deploy(IBSAS, Registry.address);

};
