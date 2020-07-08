const MadEscrow = artifacts.require("MadEscrow");
const Oracle = artifacts.require("AuthenticityOracle");

module.exports = async function(deployer) {
    await deployer.deploy(MadEscrow);
    const madEscrow = await MadEscrow.deployed()

    await deployer.deploy(Oracle);
    const oracle = await Oracle.deployed()
};