const Registry = artifacts.require("Registry");
const Sakai = artifacts.require("Sakai");
module.exports = function (deployer) {
    deployer.deploy(Registry).then(
      function (){
        return deployer.deploy(Sakai, Registry.address);
      }
    );
};
