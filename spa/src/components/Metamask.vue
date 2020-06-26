<template lang="html">
    <div id="connect">
        <button type="button" class="Login-button" id="Login-mm" onClick="createExampleContract();">Connect with Metamask</button>
    </div>
</template>

<script>
    export default {
        name: "Metamask",
        data() {
            return {
                currentAddress: "",
            }
        },
        methods: {
            async loadWeb3() {
                // Modern dapp browsers...
                if (window.ethereum) {
                    window.web3 = new Web3(window.ethereum);
                    await window.ethereum.enable()
                }
                  // Legacy dapp browsers...
                else if (window.web3) {
                    window.web3 = new Web3(window.web3.currentProvider);
                }
                // Non-dapp browsers...
                else {
                    window.alert('Non-Ethereum browser detected. You should consider trying MetaMask!');
                }
            },

            async loadBlockchainData() {
                const web3 = window.web3
                const accounts = await web3.eth.getAccounts()
                this.currentAddress = accounts[0]
                // put address into navbar ( { this.state.account } in React )
            }
        },
    }
</script>

<style scoped>

    .Login-button {
        color: white;
        display: block;
        font-size: 16px;
        height: 60px;
        margin: 10px auto;
        width: 300px;
    }

    .Login-button:hover {
        opacity: 0.9;
    }

    #Login-mm {
        background-color: rgb(255, 125, 0);
    }
</style>