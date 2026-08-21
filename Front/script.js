document.getElementById('btn').addEventListener('click', function(){
    fetch('http://localhost:8080/hello', {
        method: 'GET'
    })
    .then(response => console.log("Ответ от сервера:", response))
    .catch(err => console.error("Ошибка:", err))
})