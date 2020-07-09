import Web3 from 'web3';
export default {
  methods: {
    async loadWeb3() {
      // Modern dapp browsers...
        if (window.ethereum) {
          window.web3 = new Web3(window.ethereum);
          await window.ethereum.enable();
          const comp = this;
          loadBlockchainData(comp);
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
    loginAndSetAddress() {
      if (!this.currentAddress) {
        this.loadWeb3();
      }
    }
  }
}
async function loadBlockchainData(comp) {
  const web3 = window.web3
  const accounts = await web3.eth.getAccounts()
  comp.currentAddress = accounts[0];
}
