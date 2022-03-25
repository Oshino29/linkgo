fetch('http://localhost:8080/items')
  .then(response => response.json())
  .then(data => console.log(data));