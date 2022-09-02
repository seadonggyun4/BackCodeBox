const {odd, even} = require('./vue.js')


function getdata(num){
    if(num%2){
        console.log(odd)
    } else {
        console.log(even)
    }
}


module.exports = getdata