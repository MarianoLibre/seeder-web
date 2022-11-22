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
  "products",
  "employees",
  "sections",
  "sellers",
  "warehouses",
  "localities",
  "carries",
  "product_batches",
  "product_records",
  "purchase_orders",
  "inbound_orders",
]

const tables = []

const folderOutline = "fa-folder-o"
const folderSolid = "fa-folder"

const tableList = document.createElement("ul")
tableList.setAttribute("class", "table-list")
sideBar.appendChild(tableList)
for (let item of tableNames) {
  const newTable = document.createElement("li")
  newTable.className = "table"
  newTable.setAttribute("id", item)

  const icon = document.createElement("i")
  icon.classList = `fa ${folderOutline} fa-lg`

  const text = document.createElement("span")
  text.innerText = item
  text.onclick = selectTable
  
  const checked = document.createElement("i")
  checked.classList = `fa fa-lg fa-question-circle`

  const seedBtn = document.createElement("button")
  seedBtn.className = "seed-button"
  seedBtn.innerText = "Seed"

  const containerLeft = document.createElement("span")
  containerLeft.className = "left"
  containerLeft.appendChild(icon)
  containerLeft.appendChild(text)

  const containerRight = document.createElement("span")
  containerRight.className = "right"
  containerRight.appendChild(seedBtn)
  containerRight.appendChild(checked)

  newTable.appendChild(containerLeft)
  newTable.appendChild(containerRight)

  tables.push(newTable)
  tableList.appendChild(newTable)
}

const view = document.getElementById("view")
const initialMsg = `<div class="help-msg">
<p>Click "Seed all" to seed all tables.</p>
<p>Enter a number to override the default value (10)</p>
<p>Click on the <i class="fa fa-refresh"></i> icon to verify your connection.</p>
<p>Click on a table for details on that table.</p>
<p>You can seed each table individually by clicking on its "seed" button.</p>
<p>To show this message again click on the "Help" button.</p>
</div>`

view.innerHTML = initialMsg

const helpBtn = document.getElementById("help-btn")
helpBtn.onclick = () => {
  view.innerHTML = initialMsg
}

async function selectTable(e) {
  for (let t of tables) {
    if (t.id === e.target.innerText) {
      t.children[0].children[0].classList = `fa fa-lg ${folderSolid}`
      t.classList = "table selected"
      view.innerHTML = `<i class="fa fa-circle-o-notch fa-spin bluefg fa-2x"></i>`
      const chuckNorris = await fetch("https://api.chucknorris.io/jokes/random")
      const joke = await chuckNorris.json()
      view.innerHTML = `
        <div class="todo">
          <p><b>TODO: </b>${t.id}</p>
          <p>What did you expect? I only have two hands!</p>
          <p>Allright, I'll give you something in exchange...<br/>
          Here you have some <b>Chuck Norris</b> facts...</p>
          <p>(And for "facts" I mean jokes!)</p>
          <p class="cite">${joke.value}</p>
        </div>`
    } else {
      t.children[0].children[0].classList = `fa fa-lg ${folderOutline}`
      t.classList = "table"
    }
  }
}
