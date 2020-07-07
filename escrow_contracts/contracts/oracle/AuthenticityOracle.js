const axios = require('axios')
const BN = require('bn.js')
const common = require('./utils/common.js')
const SLEEP_INTERVAL = process.env.SLEEP_INTERVAL || 2000
const PRIVATE_KEY_FILE_NAME = process.env.PRIVATE_KEY_FILE || './oracle_private_key'
const CHUNK_SIZE = process.env.CHUNK_SIZE || 3
const MAX_RETRIES = process.env.MAX_RETRIES || 5
const OracleJSON = require('./AuthenticityOracle.sol')
var pendingRequests = []

async function getOracleContract (web3js) {
  const networkId = await web3js.eth.net.getId()
  return new web3js.eth.Contract(OracleJSON.abi, OracleJSON.networks[networkId].address)
}

// needs to be replaced by the actual endpoint for providers

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


async function filterEvents (oracleContract, web3js) {
  oracleContract.events.GetAuthenticityCheckEvent(async (err, event) => {
    if (err) {
      console.error('Error on event', err)
      return
    }
    await addRequestToQueue(event)
  })

  oracleContract.events.SetAuthenticityCheckEvent(async (err, event) => {
    if (err) console.error('Error on event', err)
    // Fill out
  })
}

async function addRequestToQueue (event) {
  const callerAddress = event.returnValues.callerAddress
  const id = event.returnValues.id
  pendingRequests.push({ callerAddress, id })
}

async function processQueue (oracleContract, ownerAddress) {
  let processedRequests = 0
  while (pendingRequests.length > 0 && processedRequests < CHUNK_SIZE) {
    const req = pendingRequests.shift()
    await processRequest(oracleContract, ownerAddress, req.id, req.callerAddress)
    processedRequests++
  }
}

async function processRequest (oracleContract, ownerAddress, id, callerAddress) {
  let retries = 0
  while (retries < MAX_RETRIES) {
    try {
      const authenticityCheck = await retrieveAuthenticity()
      await setAuthenticityCheck(oracleContract, callerAddress, ownerAddress, authenticityCheck, id)
      return
    } catch (error) {
      if (retries === MAX_RETRIES - 1) {
        await setAuthenticityCheck(oracleContract, callerAddress, ownerAddress, false, id)
        return
      }
      retries++
    }
  }
}

async function setAuthenticityCheck (oracleContract, callerAddress, ownerAddress, authenticityCheck, id) {
  const idInt = new BN(parseInt(id))
  try {
    await oracleContract.methods.setAuthenticityCheck(authenticityCheck, callerAddress, idInt.toString()).send({ from: ownerAddress })
  } catch (error) {
    console.log('Error encountered while calling setAuthenticityCheck.')
    // Do some error handling
  }
}

async function init () {
  const { ownerAddress, web3js, client } = common.loadAccount(PRIVATE_KEY_FILE_NAME)
  const oracleContract = await getOracleContract(web3js)
  filterEvents(oracleContract, web3js)
  return { oracleContract, ownerAddress, client }
}

(async () => {
  const { oracleContract, ownerAddress, client } = await init()
  process.on( 'SIGINT', () => {
    console.log('Calling client.disconnect()')
    client.disconnect()
    process.exit( )
  })
  setInterval(async () => {
    await processQueue(oracleContract, ownerAddress)
  }, SLEEP_INTERVAL)
})()

