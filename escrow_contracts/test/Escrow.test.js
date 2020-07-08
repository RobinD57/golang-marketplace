const MadEscrow = artifacts.require("MadEscrow");
const Oracle = artifacts.require("AuthenticityOracle");


require("chai")
    .use(require("chai-as-promised"))
    .should()

contract("MadEscrow", (accounts) => {
    let oracle, madEscrow

    before(async () => {
        madEscrow = await MadEscrow.new()
        oracle = await Oracle.new()
    });

    describe("MadEscrow deployment works properly", async () => {
        it("contract has a name", async() => {
            const name = await madEscrow.name()
            assert.equal(name, "MadEscrow Solidity contract")
        })
    })
});