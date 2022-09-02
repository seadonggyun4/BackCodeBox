const {odd, even} = require('./vue.js')
const getdata = require('./func.js')


function setdata(num){
    if(num%2 === 1){
        console.log(odd)
    } else {
        console.log(even)
    }
}

console.log(setdata(9))

console.log(getdata(10))


