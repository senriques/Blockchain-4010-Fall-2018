var UW4010 = artifacts.require("./UW4010.sol");

module.exports = (deployer) => {
  deployer.then(async () => {
    await deployer.deploy(UW4010);
  });
};

