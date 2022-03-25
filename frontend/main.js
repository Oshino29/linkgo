

// fetch('http://localhost:8080/items', {
//     method: 'POST',
//     headers: {
//         'Content-Type': 'application/json'
//     },
//     body: JSON.stringify({
//         title: 'codepen.io trending',
//         url: 'https://codepen.io/trending',
//         tags: ['dev:frontend', 'community']
//     })
// }).then(res => {
//     return res.json()
// })
// .then(data => console.log(data))
// .catch(error => console.log('ERROR'))



// document.addEventListener('DOMContentLoaded', () => {
//     console.log('js executed after document loaded')
//     let templateCard = document.getElementById('cardTemplate').content
//     let templateItem = document.getElementById('itemTemplate').content
//     let card = document.importNode(templateCard, true)
//     let item = document.importNode(templateItem, true)

//     console.log(card)
//     for (let index = 0; index < 30; index++) {
//         console.log(`add ${index}th element`)

        

//         card.content.getElementsByClassName('item-container').appendChild(item)
//         document.getElementById('cardContainer').appendChild(card)
//     }
// })

// main functions start
document.addEventListener('DOMContentLoaded', () => showAllItems())

// const testItem = {
//     title: "ウラアカ批評空間",
//     url: "https://uraakaspace.com",
//     tags: ["community selfhosted"]
// }

// postItem(testItem).then(() => showAllItems())
// main functions end

// function definition starts here

// show all items
async function getAllItems() {
    const items = fetch('http://localhost:8080/items').then(response => response.json())  
    return items
}

async function showAllItems() {
    getAllItems().then(items => showItems(items))
}
function showItems(items) {
    let card = makeCard('testAllTag')
    items.forEach(item => {
        let itemHTML = makeItem(item)
        addItemToCard(itemHTML, card)
    })
    addCardToDom(card)
}

// make a card with tag name
function makeCard(tag) {
    const templateCard = document.getElementById('cardTemplate').content
    let card = document.importNode(templateCard, true)
    card.querySelector('.card-title').textContent = tag
    return card
}
// append a card to DOM
function addCardToDom(card) {
    document.getElementById('cardContainer').appendChild(card)
}

// make an item with item object
function makeItem(item) {
    const templateItem = document.getElementById('itemTemplate').content
    let itemHTML = document.importNode(templateItem, true)
    itemHTML.querySelector('.item-title').textContent = item.title
    itemHTML.querySelector('a').href = item.url
    return itemHTML
}
// add an item to a card made bt makeCard()
function addItemToCard(itemHTML, card) {
    card.querySelector('.item-container').appendChild(itemHTML)
}


function addItembyInput() {
    const item = readInput()
    postItem(item)
    showAllItems()
}

function readInput() {
    const item = {
        title: document.getElementById('titleInput').value,
        url: document.getElementById('urlInput').value
    }
    return item
}
// post an item with an item object
async function postItem(item) {
    const request = new Request('http://localhost:8080/items', {
        method: 'POST',
        body: JSON.stringify(item),
        headers: {
          'Content-Type': 'application/json'
        }
      });
      const response = await fetch(request);
    //   return response.json()
}