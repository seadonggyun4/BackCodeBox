function setTmmediate(){
    console.log('immediate')
}

Process.nextTick(() => {
    console.log('nextTick')
})

setTimeout(() => {
    console.log('timeout')
}, 0)

Promise.resolve().then(() => console.log('Promise'))