

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

async function getAllItems() {
    const response =  await fetch('http://localhost:8080/items').then(response => response.json())  
    return response
}

function makeCard(tag) {
    const templateCard = document.getElementById('cardTemplate').content
    let card = document.importNode(templateCard, true)
    card.querySelector('.card-title').textContent = tag
    return card
}

function addCardToDom(card) {
    document.getElementById('cardContainer').appendChild(card)
}

function makeItem(item) {
    const templateItem = document.getElementById('itemTemplate').content
    let itemHTML = document.importNode(templateItem, true)
    itemHTML.querySelector('.item-title').textContent = item.title
    itemHTML.querySelector('a').href = item.url
    return itemHTML
}

function addItemToCard(itemHTML, card) {
    card.querySelector('.item-container').appendChild(itemHTML)
}

getAllItems().then(items => {
    let card = makeCard('testAllTag')
    items.forEach(item => {
        let itemHTML = makeItem(item)
        addItemToCard(itemHTML, card)
    })
    addCardToDom(card)
})