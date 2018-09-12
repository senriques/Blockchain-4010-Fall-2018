var UW4010 = artifacts.require("./UW4010.sol");

module.exports = (deployer) => {
  deployer.then(async () => {
    await deployer.deploy(UW4010);
    await UW4010.at(UW4010.address).newStudent("0x885765a2fcfB72e68D82D656C6411b7D27BacDD7", 761058);
    await UW4010.at(UW4010.address).newStudent("0x6E06bF940bb57adE69Cb03153d1C3842411bD3c1", 133000);
    await UW4010.at(UW4010.address).newStudent("0x3B65b88e4256C8926358551072f17460EFe5452B", 879191);
    await UW4010.at(UW4010.address).newStudent("0x40681739B0ef568acce20F5575aD4cf24223926f", 189681);
    await UW4010.at(UW4010.address).newStudent("0x4aF64cd87A47aaB7cffdbAda6bfD6aef47036c03", 115724);
    await UW4010.at(UW4010.address).newStudent("0xB0D533A8064ed180967Aa4Dafa453deAb107961C", 875844);
    await UW4010.at(UW4010.address).newStudent("0x31568CC92115a2EBE6eB37E9a7c7f6334B988196", 394673);
    await UW4010.at(UW4010.address).newStudent("0x9d41E5938767466af28865e1c33071f1561D57a8", 364175);
    await UW4010.at(UW4010.address).newStudent("0xe7B8A518bf1B5c4f01b2A7eE39a2800A982E06Ee", 169029);
    await UW4010.at(UW4010.address).newStudent("0x4e27d9C8bA3a6904F7A7Cb31Eae5cCCe8BF33300", 176706);
    await UW4010.at(UW4010.address).newStudent("0x5AE7B3Cf64Adc3D7FEf099319a9be4acb8bD73ED", 113328);
    await UW4010.at(UW4010.address).newStudent("0x0C34a1A3c5Ae302cB41F9Cfd999E7950B8eBf40f", 129368);
    await UW4010.at(UW4010.address).newStudent("0x42F487a6d5c86962310D5aB5aFE5CaD7bC80805B", 183749);
    await UW4010.at(UW4010.address).newStudent("0x7e3aFEc048bC7be745d0fA0F5af97D3978C40E9A", 237801);
  });
};

