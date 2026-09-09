const tBtn = document.getElementById('testbutton1')

function btnClicked(text){
    fetch('http://localhost:8080/hi', {
        method: 'POST',
        body: text
    }
    ).then((response) => console.log(response))
}