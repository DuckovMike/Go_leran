// ------------- Логика заголовка ------------- //

const tableNum = document.getElementById('table-num')


const tableVar = document.getElementById('table-var')


const tableDate = document.getElementById('table-date')
const tableYear = document.getElementById('table-year')

const date = new Date()

let year = date.getFullYear()

for (let i of [-1,0,1,2]){
    const optYr = document.createElement('option')
    optYr.textContent = year + i
    optYr.value = year + i
    if (i == 0){
        optYr.selected = true
    }

    tableYear.appendChild(optYr)
}



const febLastDate = new Date(year, 2, 0).getDate()
const dateOptions = [
    { text: "Январь 1-15", value: 1},
    { text: "Январь 16-31", value: 2},
    { text: "Февраль 1-15", value: 3},
    { text: `Февраль 16-${febLastDate}`, value: 4, id: 'table-date-feb'},
    { text: "Март 1-15", value: 5},
    { text: "Март 16-31", value: 6},
    { text: "Апрель 1-15", value: 7},
    { text: "Апрель 16-30", value: 8},
    { text: "Май 1-15", value: 9},
    { text: "Май 16-31", value: 10},
    { text: "Июнь 1-15", value: 11},
    { text: "Июнь 16-30", value: 12},
    { text: "Июль 1-15", value: 13},
    { text: "Июль 16-31", value: 14},
    { text: "Август 1-15", value: 15},
    { text: "Август 16-31", value: 16},
    { text: "Сентябрь 1-15", value: 17},
    { text: "Сентябрь 16-30", value: 18},
    { text: "Октябрь 1-15", value: 19},
    { text: "Октябрь 16-31", value: 20},
    { text: "Ноябрь 1-15", value: 21},
    { text: "Ноябрь 16-30", value: 22},
    { text: "Декабрь 1-15", value: 23},
    { text: "Декабрь 16-31", value: 24},
]



for (const opt of dateOptions) {
    const optEl = document.createElement('option')
    optEl.textContent = opt.text
    optEl.value = opt.value
    if ('id' in opt){
        optEl.id = 'table-date-feb'
    }
    tableDate.appendChild(optEl)
}

tableYear.addEventListener("change", () => {
    const feb = document.getElementById('table-date-feb')
    year = tableYear.value
    const febLastDate = new Date(year, 2, 0).getDate()
    feb.textContent = `Февраль 16-${febLastDate}`
})

tableDate.addEventListener("change", () => {
    tableNum.textContent = "Табель №" + `${tableDate.value}`
})



// ------------- Логика таблицы ------------- //

const container = document.getElementById("Table")

for (let i = 1; i <= 31; i++){
    const tabelcell = document.createElement('div')
    tabelcell.textContent = i
    tabelcell.id = `day-${i}`
    tabelcell.className = "days"

    container.appendChild(tabelcell)
}
