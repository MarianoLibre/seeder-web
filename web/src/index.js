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
const refreshIcon = document.getElementById("refresh")
const refreshBtn = document.getElementById("status")
refreshIcon.onclick = connect
refreshBtn.onclick = connect

async function connect() {
  refreshIcon.style.animation = "spin 1s linear infinite"
  setTimeout(() => refreshIcon.style.animation = "none", 1000)
  try {
    const url = new URL("http://localhost:8080/")
    const server = await fetch(url, {
      method: "GET",
      headers: {
        token: 1234,
        cors: "no-cors",
      },
    })
    refreshBtn.setAttribute("class", "statusUp")
    refreshBtn.innerText = "Server is UP"
  } catch (e) {
    refreshBtn.setAttribute("class", "statusDown")
    refreshBtn.innerText = "Server is DOWN"
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
  "localities",
  "carries",
  "warehouses",
  "buyers",
  "employees",
  "sellers",
  "sections",
  "products",
  "product_batches",
  "inbound_orders",
  "product_records",
  "purchase_orders",
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
  seedBtn.onclick = () => seedTable(item)

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

const dropBtn = document.getElementById("drop-tables")
dropBtn.onclick = dropTables
const createBtn = document.getElementById("create-tables")
createBtn.onclick = createTables

async function dropTables() {
  view.innerHTML = `<i class="fa fa-circle-o-notch fa-spin bluefg fa-2x"></i>`
  const result = await fetch("http://localhost:8080/seeder/drop-tables", {
    method: "POST",
  })
  const msg = await result.json()
  view.innerHTML = msg.message
}

async function createTables() {
  view.innerHTML = `<i class="fa fa-circle-o-notch fa-spin bluefg fa-2x"></i>`
  const result = await fetch("http://localhost:8080/seeder/create-tables", {
    method: "POST",
  })
  const msg = await result.json()
  view.innerHTML = msg.message
}

async function seedTable(table) {
  const quantity = document.getElementById("quantity").value
  if (isNaN(Number(quantity))) {
    quantity = 10
  }
  const route = table.replace("_", "-")

  view.innerHTML = `<i class="fa fa-circle-o-notch fa-spin bluefg fa-2x"></i>`
  const result = await fetch(`http://localhost:8080/seeder/${route}?qty=${quantity}`, {
    method: "POST",
  })
  const msg = await result.text()
  view.innerHTML = msg
  for (let t of tables) {
    if (t.id == table) {
      t.children[1].children[1].classList = `fa fa-lg fa-check-circle`
    }
  }

  return result.status
}

const seedAllBtn = document.getElementById("seed-all")
seedAllBtn.onclick = seedAll

async function seedAll() {
  const localities = await seedTable("localities")
  if (localities != 201) {
    view.innerText = "Something went wrong: localities"
    return
  }
  const carries = await seedTable("carries")
  if (carries != 201) {
    view.innerText = "Something went wrong: carries"
    return
  }
  const warehouses = await seedTable("warehouses")
  if (warehouses != 201) {
    view.innerText = "Something went wrong: warehouses"
    return
  }
  const buyers = await seedTable("buyers")
  if (buyers != 201) {
    view.innerText = "Something went wrong: buyers"
    return
  }
  const employees = await seedTable("employees")
  if (employees != 201) {
    view.innerText = "Something went wrong: employees"
    return
  }
  const sellers = await seedTable("sellers")
  if (sellers != 201) {
    view.innerText = "Something went wrong: sellers"
    return
  }
  const sections = await seedTable("sections")
  if (sections != 201) {
    view.innerText = "Something went wrong: sections"
    return
  }
  const products = await seedTable("products")
  if (products != 201) {
    view.innerText = "Something went wrong: products"
    return
  }
  const product_batches = await seedTable("product_batches")
  if (product_batches != 201) {
    view.innerText = "Something went wrong: product_batches"
    return
  }
  const inbound_orders = await seedTable("inbound_orders")
  if (inbound_orders != 201) {
    view.innerText = "Something went wrong: inbound_orders"
    return
  }
  const product_records = await seedTable("product_records")
  if (product_records != 201) {
    view.innerText = "Something went wrong: product_records"
    return
  }
  const purchase_orders = await seedTable("purchase_orders")
  if (purchase_orders != 201) {
    view.innerText = "Something went wrong: purchase_orders"
    return
  }
  view.innerHTML = `<h3>All tables successfully seeded</h3>`
}
