const http = require('http')
const fs = require('fs')


const courses = [{name: 'HTML'}, {name: 'CSS'}, {name: 'JS'} ,{name: 'NODE'} ,{name: 'VUE'} ,{name: 'REACT'}]




const server = http.createServer((req,  res) => {
    const URL = req.url
    const method = req.method

    if(URL === '/courses'){
        switch (method) {
            case 'GET':
                res.writeHead(200, {'Content-Type':'application/json'});
                res.end(JSON.stringify(courses))
                break
            case 'POST':
                const body = []
                req.on('data', (chunk) => {
                    console.log(chunk)
                    body.push(chunk)
                })


                req.on('end', () => {
                    const bodyStr = Buffer.concat(body).toString()
                    const course = JSON.parse(bodyStr)
                    courses.push(course)
                    console.log(course)
                    res.writeHead(201)
                    res.end()
                })

                break
        }
    }
})




server.listen(8888, () => {
    console.log('Start Server')
})