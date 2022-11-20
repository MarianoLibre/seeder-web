const root = document.getElementById("root")

const colors = ["black", "white", "red", "blue", "yellow", "green", "cyan", "magenta"]

const colortest = document.createElement("div")
colortest.setAttribute("id", "colortest")
colortest.style.display = "none"
const colortestInner = document.createElement("div")
colortestInner.setAttribute("id", "ct_inner")
for (let c of colors) {
  const p = document.createElement("p")
  p.innerText = "I'm a paragraph, you know?"
  p.classList += c + "fg"
  if (c === "black") p.classList += " whitebg"
  colortestInner.appendChild(p)
}

colortest.onclick = toggleColors
colortestInner.onclick = e => {e.stopPropagation()}
colortest.appendChild(colortestInner)
root.appendChild(colortest)

const showColorBtn = document.createElement("button")
showColorBtn.innerText = "Show colors"
showColorBtn.setAttribute("class", "showBtn")
root.appendChild(showColorBtn)

function toggleColors() {
  const show = colortest.style.display
  if (show == "none") {
    colortest.style.display = "flex"
    showColorBtn.innerText = "Hide colors"
  } else {
    colortest.style.display = "none"
    showColorBtn.innerText = "Show colors"
  }
}

showColorBtn.onclick = toggleColors


// Connect to Server
const refreshBtn = document.getElementById("refresh")
refreshBtn.onclick = connect

async function connect() {
  refreshBtn.style.animation = "spin 1s linear infinite"
  setTimeout(() => refreshBtn.style.animation = "none", 1000)
  const statusBtn = document.getElementById("status")
  try {
    const url = new URL("http://localhost:8080/")
    const server = await fetch(url, {
      method: "GET",
      headers: {
        token: 1234,
        cors: "no-cors",
      },
    })
    statusBtn.setAttribute("class", "statusUp")
    statusBtn.innerText = "Server is UP"
  } catch (e) {
    statusBtn.setAttribute("class", "statusDown")
    statusBtn.innerText = "Server is DOWN"
  }
}


/*
const response = document.createElement("div")
response.setAttribute("class", "response")

const btn = document.createElement("button")
btn.setAttribute("class", "btn")
btn.innerText = "Get warehouses"

root.appendChild(btn)
root.appendChild(response)

btn.onclick = async function () {
  try {
    const url = new URL("http://localhost:8080/api/v1/warehouses")
    const result = await fetch(url, {
      method: 'GET',
      headers: {
        token: 1234,
      },
    })
    data = await result.json()
    response.innerText = JSON.stringify(data)
    console.log(data.data)
  } catch (e) {
    response.innerText = "???"
  }
}
*/
const sideBar = document.getElementById("side-bar")

const tableNames = [
  "buyers",
  "carries",
  "employees",
  "inbound_orders",
  "localhost",
  "product_batches",
  "product_records",
  "products",
  "purchase_orders",
  "sections",
  "sellers",
  "warehouses",
]

const tables = []

const folderOutline = "fa-folder-o"
const folderSolid = "fa-folder"

for (let item of tableNames) {
  const newTable = document.createElement("p")
  newTable.className = "table"
  newTable.setAttribute("id", item)
  const icon = document.createElement("i")
  icon.classList = `fa ${folderOutline} fa-lg`
  icon.onclick = e => e.stopPropagation()
  const text = document.createElement("span")
  text.innerText = item
  text.onclick = e => e.stopPropagation()
  newTable.appendChild(icon)
  newTable.appendChild(text)
  newTable.onclick = selectTable
  tables.push(newTable)
  sideBar.appendChild(newTable)
}

function selectTable(e) {
  for (let t of tables) {
    t.firstChild.classList = `fa ${folderOutline} fa-lg`
    t.classList = "table"
  }
  console.log(e.target.innerHTML)
}
