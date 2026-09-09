const btn = document.getElementById('test_btn')
const tst = document.getElementById('test')

let innerInfo = tst.value

tst.addEventListener('change', ()=>{
    innerInfo = tst.value
    console.log(innerInfo);
})

function handleClick(text){
    fetch('http://localhost:8080/hi',{
        method:'POST',
        headers: {'Content-type': 'application/json'},
        body: JSON.stringify({'key': innerInfo})
    }).then((response) => response.text())
    .then(data => console.log(data))
    .catch(error => console.log(error))
}

btn.addEventListener('click', handleClick)