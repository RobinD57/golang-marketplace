const common = require('./utils/common.js')
const SLEEP_INTERVAL = process.env.SLEEP_INTERVAL || 2000
const PRIVATE_KEY_FILE_NAME = process.env.PRIVATE_KEY_FILE || './escrow_private_key'
const CallerJSON = require('../escrow/MadEscrow.sol')
const OracleJSON = require('./Authenticity.json')

async function getCallerContract (web3js) {
  const networkId = await web3js.eth.net.getId()
  return new web3js.eth.Contract(CallerJSON.abi, CallerJSON.networks[networkId].address)
}

async function retrieveAuthenticityCheck () {
  const resp = await axios({
    url: 'http://dummy.restapiexample.com/api/v1/employees',
    // params: {
    //   ...
    // },
    method: 'get'
  })
  return resp
}

async function filterEvents (callerContract) {
  callerContract.events.AuthenticityCheckCompletedEvent({ filter: { } }, async (err, event) => {
    if (err) console.error('Error on event', err)
    console.log('* New AuthenticityCheckCompletedEvent event. Authenticity: ' + event.returnValues.authenticityCheck)
  })
  callerContract.events.ReceivedNewRequestIdEvent({ filter: { } }, async (err, event) => {
    if (err) console.error('Error on event', err)
  })
}

async function init () {
  const { ownerAddress, web3js, client } = common.loadAccount(PRIVATE_KEY_FILE_NAME)
  const callerContract = await getCallerContract(web3js)
  filterEvents(callerContract)
  return { callerContract, ownerAddress, client, web3js }
}

(async () => {
  const { callerContract, ownerAddress, client, web3js } = await init()
  process.on( 'SIGINT', () => {
    console.log('Calling client.disconnect()')
    client.disconnect();
    process.exit( );
  })
  const networkId = await web3js.eth.net.getId()
  const oracleAddress =  OracleJSON.networks[networkId].address
  await callerContract.methods.setOracleInstanceAddress(oracleAddress).send({ from: ownerAddress })
  setInterval( async () => {
    await callerContract.methods.updateAuthenticityCheck().send({ from: ownerAddress })
  }, SLEEP_INTERVAL);
})()
