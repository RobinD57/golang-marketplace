const Escrow = artifacts.require("/escrow/MadEscrow");
const Oracle = artifacts.require("oracle/AuthenticityOracle");

module.exports = async function(deployer) {
    await deployer.deploy(Escrow);
    const escrow = await Escrow.deployed()

    await deployer.deploy(Oracle);
    const oracle = await Oracle.deployed()
};