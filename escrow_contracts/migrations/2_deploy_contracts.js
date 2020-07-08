const MadEscrow = artifacts.require("escrow/MadEscrow");
const Oracle = artifacts.require("oracle/AuthenticityOracle");

module.exports = async function(deployer) {
    await deployer.deploy(MadEscrow);
    const madEscrow = await MadEscrow.deployed()

    await deployer.deploy(Oracle);
    const oracle = await Oracle.deployed()
};