// fetch('http://localhost:8080/items/frontend')
//   .then(response => response.json())
//   .then(data => console.log(data))

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



document.addEventListener('DOMContentLoaded', () => {
    console.log('js executed after document loaded')
    let temp = document.getElementById('cardTemplate')
    let content = temp.content.cloneNode(true)
    console.log(content)
    for (let index = 0; index < 30; index++) {
        console.log(`add ${index}th element`)

        content = temp.content.cloneNode(true)
        document.getElementById('cardContainer').appendChild(content)
    }
})