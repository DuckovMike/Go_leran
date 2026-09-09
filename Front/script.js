// ------------- Логика заголовка ------------- //

function createUser(fio, profesion, rate, appearence){
    return {
        fio: fio,
        profesion: profesion,
        rate: rate,
        appearence: appearence
    }
}

function createAppearence(year, month, part){
    const dict = {}
    const fisrtDate = part == 1 ? 1 : 16
    const endDate = part == 1 ? 15: new Date(year, month, 0).getDate()

    for (let date = fisrtDate; date <= endDate; date++){
        dict[String(date)] = new Date(year, month, date).getDay()
    }
    
    return{dict}
}

function addOpts(start, end, optList, parent){
    for (let i = start; i < end; i++) {
        const opt = document.createElement('option')
        const element = optList[i];
        opt.value = i
        opt.innerText = element

        parent.appendChild(opt)
    }
}


// Константы для хранения элементов заголовка (Номер таблицы, Даты отчета,)
const tableNum = document.getElementById('table-num')
const tableDate = document.getElementById('table-date')
const tableYear = document.getElementById('table-year')

// Сегодняшняя дата и время
const date = new Date()

// Переменная для хранения года
let year = date.getFullYear()

// Переменная для слежки за последней датой февраля
let febLastDate = new Date(year, 2, 0).getDate()

// Константы для выбора месяца и его части
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

// Создание опций для выбора года
for (let i of [-1,0,1,2]) {
    const optYr = document.createElement('option')
    optYr.textContent = year + i
    optYr.value = year + i
    if (i == 0){
        optYr.selected = true
    }

    tableYear.appendChild(optYr)
}

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
    febLastDate = new Date(year, 2, 0).getDate()
    feb.textContent = `Февраль 16-${febLastDate}`
})

tableDate.addEventListener("change", () => {
    tableNum.textContent = "Табель №" + `${tableDate.value}`
})



// ------------- Логика таблицы ------------- //

// Начальная генерация таблицы
const tableDates = document.getElementById("Table-dates-title")
const tableRes = document.getElementById("Table-resultT-title")
const tableNums = document.getElementById("Table-nums-title")
const tableResNum = document.getElementById("Table-result-num")

let dayCells = 15

tableDates.style.gridTemplateColumns = 'repeat(15, 1fr)'
tableNums.style.gridTemplateColumns = `repeat(15, 1fr)`

for (let i = 1; i <= 15; i++){
    const tabelcell = document.createElement('div')
    tabelcell.textContent = i
    tabelcell.className = "tabel-day-title"
    tableDates.appendChild(tabelcell)

    const tabelnum = document.createElement('div')
    tabelnum.textContent = i + 5
    tabelnum.className = "tabel-num-title"
    tableNums.appendChild(tabelnum)
}

// Обноваление дат и цифр в таблице при изменении месяца и/или его части
tableDate.addEventListener('change', ()=>{
    tableDates.innerHTML = ''
    tableNums.innerHTML = ''
    tableResNum.innerHTML = ''


    if (tableDate.value % 2 != 0) {
        tableDates.style.gridTemplateColumns = `repeat(${15}, 1fr)`
        tableNums.style.gridTemplateColumns = `repeat(${15}, 1fr)`
        tableRes.textContent = "Итого дней(часов) явок(неявок) с 1 по 15"
        tableResNum.textContent = 21

        for (let i = 1; i <= 15; i++) {
            const tabelcell = document.createElement('div')
            const tabelnum = document.createElement('div')
            
            tabelcell.textContent = i
            tabelcell.className = "tabel-day-title"
            tableDates.appendChild(tabelcell)

            tabelnum.textContent = i + 5
            tabelnum.className = "tabel-num-title"
            tableNums.appendChild(tabelnum)
        }
    } else {
        dayCells = new Date(year, tableDate.value/2, 0).getDate() - 15
        
        tableDates.style.gridTemplateColumns = `repeat(${dayCells}, 1fr)`
        tableNums.style.gridTemplateColumns = `repeat(${dayCells}, 1fr)`
        tableRes.textContent = "Всего дней (часов) явок (неявок) за месяц"
        tableResNum.textContent = 22 + dayCells

        for (let i = 16; i <= dayCells + 15; i++) {
            const tabelcell = document.createElement('div')
            const tabelnum = document.createElement('div')
            
            tabelcell.textContent = i
            tabelcell.className = "tabel-day-title"
            tableDates.appendChild(tabelcell)

            tabelnum.textContent = i + 6
            tabelnum.className = "tabel-num-title"
            tableNums.appendChild(tabelnum)
        }
    }
})

const workerDatesCell = document.getElementById("Table-Worker-dates")
workerDatesCell.style.gridTemplateColumns = "repeat(15, 1fr)"

let defaultAppearence = createAppearence(year, date)

const listTypes = ["В","П","О","Б","Н","К","Х"]

for (let i = 1; i <=15; i++) {
    
    const datek = document.createElement('div')
    const datev = document.createElement('select')
    
    datek.textContent = i
    
    workerDatesCell.appendChild(datek)
    workerDatesCell.appendChild(datev)

    addOpts(0, listTypes.length, listTypes, datev)
}

tableDate.addEventListener('change', ()=>{
    workerDatesCell.innerHTML = ''
    if (tableDate.value % 2 != 0) {
        workerDatesCell.style.gridTemplateColumns = "repeat(15, 1fr)"
        for (let i = 1; i <=15; i++) {
            const datek = document.createElement('div')
            const datev = document.createElement('select')
            
            datek.textContent = i
            
            workerDatesCell.appendChild(datek)
            workerDatesCell.appendChild(datev)

            addOpts(0, listTypes.length, listTypes, datev)
        }
    } else {
        workerDatesCell.style.gridTemplateColumns = `repeat(${dayCells}, 1fr)`
        for (let i = 16; i <=dayCells + 15; i++) {
            const datek = document.createElement('div')
            const datev = document.createElement('select')
            
            datek.textContent = i
            
            workerDatesCell.appendChild(datek)
            workerDatesCell.appendChild(datev)

            addOpts(0, listTypes.length, listTypes, datev)
        }
    }
})