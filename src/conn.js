// Connect to Server
async function connect() {
  try {
    const url = new URL("http://localhost:8080/")
    const server = await fetch(url, {
      method: 'GET',
      headers: {
        token: 1234,
        cors: 'no-cors',
      },
    })
    result = await server.text()
    console.log(result)
  } catch (e) {
    console.log("ERROR: connect")
    console.log(e)
  }
}

connect()

async function getWarehouses() {
  try {
    const url = new URL("http://localhost:8080/api/v1/warehouses")
    const result = await fetch(url, {
      method: 'GET',
      headers: {
        token: 1234,
      },
    })
    data = await result.json()
    console.log(data)
  } catch (e) {
    console.log("ERROR: Warehouses")
  }
}

getWarehouses()

