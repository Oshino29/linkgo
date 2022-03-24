fetch('https://linkgo.ccllssd.com/items')
  .then(response => response.json())
  .then(data => console.log(data));